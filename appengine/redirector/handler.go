package redirector

import (
	"math/rand"
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	url := calculateRedirectUrl()
	if url == "" {
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func calculateRedirectUrl() string {
	// TODO matijav 17.12.2016 read from datastore
	idx := rand.Intn(len(ENDPOINTS)) + 1
	if idx == len(ENDPOINTS) {
		return ""
	}
	return ENDPOINTS[idx]
}