package realip

import (
	"net/http"
	"strings"
)

// checks if the given slice contains the given string
func sContains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

type realIPHandler struct {
	h              http.Handler
	trustedProxies []string
}

func New(h http.Handler, trustedProxies []string) *realIPHandler {
	realip := &realIPHandler{
		h:              h,
		trustedProxies: trustedProxies,
	}

	return realip
}

func (h realIPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		if sContains(h.trustedProxies, ip) {
			r.RemoteAddr = strings.TrimSpace(strings.Split(xff, ",")[0])
		}
	}
	h.h.ServeHTTP(w, r)
}
