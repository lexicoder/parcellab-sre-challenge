package handlers

import (
	"fmt"
	"net/http"
)

func SalutationHandler(salutation string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, salutation)
	}
}
