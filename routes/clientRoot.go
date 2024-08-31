package routes

import (
	pages "assignment/client/pages/home"
	"assignment/utils"
	"net/http"
)

func ClientRoot(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			utils.ParseHtml(w, pages.MainHtml)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
