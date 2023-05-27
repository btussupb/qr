package transport

import (
	"net/http"

	wayToNomad "github.com/btussupb/qr/internal/nomad"
)

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/colon1", wayToNomad.Colon1)
	mux.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("front"))))
	return mux
}
