package middleware

import (
	"fmt"
	"net/http"
	"os"
)

var dir string = "assets"
var fs http.Handler = http.FileServer(http.Dir(dir))

func ServeFiles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := dir + r.URL.Path
		_, err := os.Stat(path)
		if os.IsNotExist(err) || r.URL.Path == "/" {
			if r.URL.Path == "/" {
				next.ServeHTTP(w, r)
				return
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "not found")
				return
			}
		}
		fs.ServeHTTP(w, r)
	})
}
