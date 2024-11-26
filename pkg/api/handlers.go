package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		w.Write([]byte("Hello, noname\n"))
		return
	}
	ans := fmt.Sprintf("%s, %s\n", "hello", name)
	w.Write([]byte(ans))

}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye\n"))
}
