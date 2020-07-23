/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by injection-gen. DO NOT EDIT.

package containersource

import (
	fmt "fmt"

	types "k8s.io/apimachinery/pkg/types"
	cache "k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/eventing/pkg/apis/sources/v1beta1"
)

// state is used to track the state of a reconciler in a single run.
type state struct {
	// Key is the original reconciliation key from the queue.
	key string
	// Namespace is the namespace split from the reconciliation key.
	namespace string
	// Namespace is the name split from the reconciliation key.
	name string
	// reconciler is the reconciler.
	reconciler Interface
	// rof is the read only interface cast of the reconciler.
	roi ReadOnlyInterface
	// IsROI (Read Only Interface) the reconciler only observes reconciliation.
	isROI bool
	// rof is the read only finalizer cast of the reconciler.
	rof ReadOnlyFinalizer
	// IsROF (Read Only Finalizer) the reconciler only observes finalize.
	isROF bool
	// IsLeader the instance of the reconciler is the elected leader.
	isLeader bool
}

func newState(key string, r *reconcilerImpl) (*state, error) {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return nil, fmt.Errorf("invalid resource key: %s", key)
	}

	roi, isROI := r.reconciler.(ReadOnlyInterface)
	rof, isROF := r.reconciler.(ReadOnlyFinalizer)

	isLeader := r.IsLeaderFor(types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	})

	return &state{
		key:        key,
		namespace:  namespace,
		name:       name,
		reconciler: r.reconciler,
		roi:        roi,
		isROI:      isROI,
		rof:        rof,
		isROF:      isROF,
		isLeader:   isLeader,
	}, nil
}

// isNotLeaderNorObserver checks to see if this reconciler with the current
// state is enabled to do any work or not.
// isNotLeaderNorObserver returns true when there is no work possible for the
// reconciler.
func (s *state) isNotLeaderNorObserver() bool {
	if !s.isLeader && !s.isROI && !s.isROF {
		// If we are not the leader, and we don't implement either ReadOnly
		// interface, then take a fast-path out.
		return true
	}
	return false
}

func (s *state) reconcileMethodFor(o *v1beta1.ContainerSource) (string, doReconcile) {
	if o.GetDeletionTimestamp().IsZero() {
		if s.isLeader {
			return doReconcileKind, s.reconciler.ReconcileKind
		} else if s.isROI {
			return doObserveKind, s.roi.ObserveKind
		}
	} else if fin, ok := s.reconciler.(Finalizer); s.isLeader && ok {
		return doFinalizeKind, fin.FinalizeKind
	} else if !s.isLeader && s.isROF {
		return doObserveFinalizeKind, s.rof.ObserveFinalizeKind
	}
	return "unknown", nil
}
