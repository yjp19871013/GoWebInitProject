package handlers

import (
	"net/http"
)

var Router = map[string]http.HandlerFunc {
	"/": Index,
}