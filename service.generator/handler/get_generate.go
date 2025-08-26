package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vhall1/password-generator/service.generator/domain"
)

func GetGenerate(w http.ResponseWriter, r *http.Request) {
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
	password := domain.RandPassword(length, sets)
	fmt.Fprint(w, password)
}
