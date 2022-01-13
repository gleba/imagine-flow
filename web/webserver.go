package web

import (
	"fmt"
	"imagine-flow/vars"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func StartWebs() {
	fs := http.FileServer(http.Dir(vars.PublicImagesFolder))
	if vars.URL_PREFIX != "" && vars.URL_PREFIX != "/" {
		fxURL := "/" + vars.URL_PREFIX + "/"
		http.Handle(fxURL, http.StripPrefix(fxURL, fs))
	} else {
		http.Handle("/", fs)
	}
	http.HandleFunc("/sync/", syncHandler)
	http.HandleFunc(vars.URL_PREFIX+"/re", resizeHandler)
	fmt.Println("imagine started at", vars.PORT, "port")
	err := http.ListenAndServe(":"+vars.PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func write(w http.ResponseWriter, code int, bytes []byte) {
	lengthInt := len(bytes)
	length := strconv.Itoa(lengthInt)
	w.Header().Add("Content-Length", length)
	w.WriteHeader(code)
	w.Write(bytes)
}

func syncHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	file := strings.Join(parts[3:], "/")
	imageCache.Drop(file)
	write(w, 200, []byte(file))
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.RawQuery
	imagePath := r.URL.Query().Get("f")
	size := r.URL.Query().Get("size")
	format := r.URL.Query().Get("as")
	if size == "" {
		write(w, 200, []byte("need assign size"))
		return
	}
	cache, have := imageCache.Get(key)
	if have {
		write(w, 200, cache)
		return
	}
	fine, code, bytes := resize(path.Join(vars.PublicImagesFolder, imagePath), size, format)
	write(w, code, bytes)
	if fine {
		imageCache.Set(key, bytes)
		imageCache.Alias(key, imagePath)
	}
}
