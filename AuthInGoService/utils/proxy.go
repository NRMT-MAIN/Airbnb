package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ReverseProxy(targetUrl string, oldPrefix string ,  pathPrefix string) http.HandlerFunc {
	target , _ := url.Parse(targetUrl)

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)
		r.URL.Scheme = target.Scheme
		r.URL.Host = target.Host

		r.URL.Path = strings.Replace(r.URL.Path , oldPrefix , pathPrefix , 1)

		fmt.Println("Forwarding request to :" , r.URL.String())

		r.Host = target.Host

		if userId, ok := r.Context().Value("userID").(string); ok {
			r.Header.Set("X-User-ID", userId)
		}
	}
	return proxy.ServeHTTP
}