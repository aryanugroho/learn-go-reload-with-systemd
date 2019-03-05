package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-systemd/activation"
)

func main() {
	http.HandleFunc("/", index)

	listeners, err := activation.Listeners()
	if err != nil {
		panic(err)
	}

	if len(listeners) != 1 {
		panic("Unexpected number of socket activation fds")
	}
	log.Fatal(http.Serve(listeners[0], nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
