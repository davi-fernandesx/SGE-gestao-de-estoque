package interfaces

import "net/http"

type ControllerApiRest interface {
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
	Post()http.HandlerFunc
	Delete()http.HandlerFunc
	Update()http.HandlerFunc
}