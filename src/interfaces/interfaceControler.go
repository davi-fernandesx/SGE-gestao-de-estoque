package interfaces

import "net/http"

type ControllerApiRest[t any] interface {
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
	Post()http.HandlerFunc
	Delete()http.HandlerFunc
	Update()http.HandlerFunc
}