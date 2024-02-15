// Package handlers implements handler functions used to handle
// requests in the service.
package handlers

import (
	"fmt"
	"net/http"
)

// A handler fucntion for returning the configured salutation.
func SalutationHandler(salutation string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, salutation)
	}
}
