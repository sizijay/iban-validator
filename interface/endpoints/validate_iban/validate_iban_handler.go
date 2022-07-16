package validate_iban

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/sizijay/iban-validator/domain"
	"github.com/sizijay/iban-validator/domain/errors"
	requestEntity "github.com/sizijay/iban-validator/internal/http/request"
	responseEntity "github.com/sizijay/iban-validator/internal/http/response"
	goValidator "gopkg.in/go-playground/validator.v9"
	"net/http"
	"runtime/debug"
)

func ValidateIBANEndpoint(validateIBANService domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, err := validateIBANService.Do(ctx, []interface{}{request})

		if len(r) == 0 {
			return response, errors.NewApplicationError("Application Error", errors.ErrEmptyResponse, string(debug.Stack()))
		}

		return r[0], err
	}
}

func DecodeValidateIBANRequest(ctx context.Context, r *http.Request) (req interface{}, err error) {
	var request requestEntity.ValidateIBANRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(fmt.Sprintf("Error decoding the IBAN number. Error: %+v", err.Error()))

		return req, errors.NewValidationError("Error Decoding the IBAN Value", errors.ErrDecodeIBAN, string(debug.Stack()))
	}

	validate := goValidator.New()
	err = validate.Struct(&request)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error validating the IBAN number. Error: %+v", err.Error()))

		return req, errors.NewValidationError("Validating Decoding the IBAN Value", errors.ErrValidateIBAN, string(debug.Stack()))
	}

	return request, nil
}

func EncodeValidateIBANResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	r := response.(bool)

	resp := responseEntity.ValidateIBANResponse{}
	resp.Data.IsValidIBAN = r

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(resp)
}

