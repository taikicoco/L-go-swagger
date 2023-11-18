package server

import (
	"github.com/go-openapi/runtime/middleware"
	"l-go-swagger/server/gen/restapi/factory"
)

func GetGreeting(p factory.GetGreetingParams) middleware.Responder {
	payload := *p.Name
	return factory.NewGetGreetingOK().WithPayload(payload)
}