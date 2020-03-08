package controllers

import (
	"github.com/noorelbahr/golearn/helpers"
	"net/http"
)

/**
 * Welcome page or Home Page
 */
func HomePage(w http.ResponseWriter, r *http.Request) {
	helpers.JsonSuccess(w, "This is my homepage", 200)
}
