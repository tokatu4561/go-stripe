package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// _, err := session(w, r)
	// if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	// } else {
	// 	http.Redirect(w, r, "/todos", 302)
	// }
}