package actions

import (
	"net/http"

	"github.com/tiger5226/skype-forwarder/skype"

	"github.com/lbryio/lbry.go/extras/api"
	"github.com/tiger5226/skype-forwarder/orderedmap"
)

type Routes struct {
	m *orderedmap.Map
}

func (r *Routes) Set(key string, h api.Handler) {
	if r.m == nil {
		r.m = orderedmap.New()
	}
	r.m.Set(key, h)
}

func GetRoutes() *Routes {
	routes := Routes{}

	routes.Set("/", Root)
	routes.Set("/test", Test)
	routes.Set("/skype", skype.SendMessage)

	return &routes
}

func (r *Routes) Each(f func(string, http.Handler)) {
	if r.m == nil {
		return
	}
	for _, k := range r.m.Keys() {
		a, _ := r.m.Get(k)
		f(k, a.(http.Handler))
	}
}

func (r *Routes) Walk(f func(string, http.Handler) http.Handler) {
	if r.m == nil {
		return
	}
	for _, k := range r.m.Keys() {
		a, _ := r.m.Get(k)
		r.m.Set(k, f(k, a.(http.Handler)))
	}
}
