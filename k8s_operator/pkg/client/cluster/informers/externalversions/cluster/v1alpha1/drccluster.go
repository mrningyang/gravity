/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	cluster_v1alpha1 "github.com/moiot/gravity/k8s_operator/pkg/apis/cluster/v1alpha1"
	versioned "github.com/moiot/gravity/k8s_operator/pkg/client/cluster/clientset/versioned"
	internalinterfaces "github.com/moiot/gravity/k8s_operator/pkg/client/cluster/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/moiot/gravity/k8s_operator/pkg/client/cluster/listers/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DrcClusterInformer provides access to a shared informer and lister for
// DrcClusters.
type DrcClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DrcClusterLister
}

type drcClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDrcClusterInformer constructs a new informer for DrcCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDrcClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDrcClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDrcClusterInformer constructs a new informer for DrcCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDrcClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DrcV1alpha1().DrcClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DrcV1alpha1().DrcClusters(namespace).Watch(options)
			},
		},
		&cluster_v1alpha1.DrcCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *drcClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDrcClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *drcClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cluster_v1alpha1.DrcCluster{}, f.defaultInformer)
}

func (f *drcClusterInformer) Lister() v1alpha1.DrcClusterLister {
	return v1alpha1.NewDrcClusterLister(f.Informer().GetIndexer())
}