package controlcenter

import (
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, getPath()+"/layout/index.html")
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, getPath()+"/"+r.URL.Path[1:])
}
