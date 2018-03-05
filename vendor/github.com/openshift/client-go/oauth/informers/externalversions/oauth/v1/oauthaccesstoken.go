// This file was automatically generated by informer-gen

package v1

import (
	oauth_v1 "github.com/openshift/api/oauth/v1"
	versioned "github.com/openshift/client-go/oauth/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/oauth/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/oauth/listers/oauth/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// OAuthAccessTokenInformer provides access to a shared informer and lister for
// OAuthAccessTokens.
type OAuthAccessTokenInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.OAuthAccessTokenLister
}

type oAuthAccessTokenInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewOAuthAccessTokenInformer constructs a new informer for OAuthAccessToken type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOAuthAccessTokenInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.OauthV1().OAuthAccessTokens().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.OauthV1().OAuthAccessTokens().Watch(options)
			},
		},
		&oauth_v1.OAuthAccessToken{},
		resyncPeriod,
		indexers,
	)
}

func defaultOAuthAccessTokenInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewOAuthAccessTokenInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *oAuthAccessTokenInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&oauth_v1.OAuthAccessToken{}, defaultOAuthAccessTokenInformer)
}

func (f *oAuthAccessTokenInformer) Lister() v1.OAuthAccessTokenLister {
	return v1.NewOAuthAccessTokenLister(f.Informer().GetIndexer())
}
