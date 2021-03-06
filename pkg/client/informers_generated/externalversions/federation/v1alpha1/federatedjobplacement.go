/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	federation_v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation/v1alpha1"
	clientset "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/kubernetes-sigs/federation-v2/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/client/listers_generated/federation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FederatedJobPlacementInformer provides access to a shared informer and lister for
// FederatedJobPlacements.
type FederatedJobPlacementInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FederatedJobPlacementLister
}

type federatedJobPlacementInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedJobPlacementInformer constructs a new informer for FederatedJobPlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedJobPlacementInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedJobPlacementInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedJobPlacementInformer constructs a new informer for FederatedJobPlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedJobPlacementInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1alpha1().FederatedJobPlacements(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1alpha1().FederatedJobPlacements(namespace).Watch(options)
			},
		},
		&federation_v1alpha1.FederatedJobPlacement{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedJobPlacementInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedJobPlacementInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedJobPlacementInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federation_v1alpha1.FederatedJobPlacement{}, f.defaultInformer)
}

func (f *federatedJobPlacementInformer) Lister() v1alpha1.FederatedJobPlacementLister {
	return v1alpha1.NewFederatedJobPlacementLister(f.Informer().GetIndexer())
}
