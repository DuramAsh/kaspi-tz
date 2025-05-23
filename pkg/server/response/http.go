package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Object struct {
	Success bool   `json:"success"`
	Message string `json:"errors,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func OK(w http.ResponseWriter, r *http.Request, data any) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, data)
}

func NoContent(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)

	v := Object{
		Success: true,
	}
	
	render.JSON(w, r, v)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusBadRequest)

	v := Object{
		Success: false,
		Message: err.Error(),
	}
	
	render.JSON(w, r, v)
}

func NotFound(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusNotFound)

	v := Object{
		Success: false,
		Message: err.Error(),
	}

	render.JSON(w, r, v)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusInternalServerError)

	v := Object{
		Success: false,
		Message: err.Error(),
	}

	render.JSON(w, r, v)
}
