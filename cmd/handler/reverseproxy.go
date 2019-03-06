package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// ReverseProxy forwards a request to the appropriate service and sends the
// response back to the caller.
func ReverseProxy(routes map[string]string) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		segments := strings.Split(r.URL.Path, "/")
		if len(segments) < 1 {
			return errors.New("invalid path")
		}

		target := segments[1]
		route, ok := routes[target]
		if !ok {
			return fmt.Errorf("unknown path \"%s\"", target)
		}

		url, err := url.Parse(route)
		if err != nil {
			return err
		}

		proxy := httputil.NewSingleHostReverseProxy(url)

		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-For", r.Header.Get("Host"))
		r.Host = url.Host

		proxy.ServeHTTP(w, r)

		return nil
	}
}
