package main

import (
    "fmt"
    "log"
		"math"
		"errors"
    "net/http"
)

func greeting(name string) (string, error) {
	if name == "" {
    return "", errors.New("empty name")
  }

	x := 0.0001

	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}

	message := fmt.Sprintf("<b>%v</b>", name)
	return message, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	message, err := greeting("Code.education Rocks!")

	if err != nil {
		log.Fatal(err)
	}

  fmt.Fprintf(w, message)
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8082", nil))
}