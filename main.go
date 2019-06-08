package main

import (
	"fmt"
	"net/http"
	"strings"

	utility "github.com/ankitm27/go_library/utility"
)

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
	s := utility.GetSum(2, 3)
	fmt.Println(s)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}
