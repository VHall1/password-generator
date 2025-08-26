package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

var sets = map[string][]string{
	"lowercase": {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	"uppercase": {"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
	"digits":    {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	"special":   {"@", "%", "!", "?", "*", "^", "&"},
}

func generate(length int, s []string) string {
	allowed := []string{}
	for _, name := range s {
		if chars, ok := sets[name]; ok {
			allowed = append(allowed, chars...)
		}
	}

	if len(allowed) == 0 || length <= 0 {
		return ""
	}

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		password[i] = allowed[rand.Intn(len(allowed))][0]
	}

	return string(password)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		// TODO: this shouldn't be deployed like this
		// only allowing * domains to facilitate development
		w.Header().Set("Access-Control-Allow-Origin", "*")

		setsParam := r.URL.Query().Get("sets")
		sets := []string{}

		if setsParam != "" {
			sets = strings.Split(setsParam, ",")
			for i := range sets {
				sets[i] = strings.TrimSpace(sets[i])
			}
		}

		length := 12
		password := generate(length, sets)
		fmt.Fprint(w, password)
	})

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", mux)
}
