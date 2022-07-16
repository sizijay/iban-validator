package server

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/sizijay/iban-validator/domain/errors"
	"github.com/sizijay/iban-validator/interface/endpoints/ping"
	"github.com/sizijay/iban-validator/interface/endpoints/validate_iban"
	"github.com/sizijay/iban-validator/service"
	"net/http"
)

//options for handler
var opts = []httpTransport.ServerOption{
	httpTransport.ServerErrorEncoder(errors.EncodeGeneralErrorResponse),
}

//ping handler
func Ping() (handler http.Handler) {
	return httpTransport.NewServer(
		ping.PingingEndpoint(
			&service.PingService{},
		),
		ping.DecodePingRequest,
		ping.EncodePingResponse,
		opts...,
	)
}

//validate iban handler
func ValidateIBAN() (handler http.Handler) {
	return httpTransport.NewServer(
		validate_iban.ValidateIBANEndpoint(
			&service.ValidateIBANService{},
		),
		validate_iban.DecodeValidateIBANRequest,
		validate_iban.EncodeValidateIBANResponse,
		opts...,
	)
}