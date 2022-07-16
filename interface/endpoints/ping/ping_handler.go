package ping

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/sizijay/iban-validator/domain"
	"github.com/sizijay/iban-validator/domain/errors"
	"net/http"
	"runtime/debug"
)

func PingingEndpoint(pingService domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, err := pingService.Do(ctx, nil)

		if len(r) == 0 {
			return response, errors.NewApplicationError("Application Error", errors.ErrPing, string(debug.Stack()))
		}

		return r[0], err
	}
}

func DecodePingRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}

func EncodePingResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	r := response.(string)

	w.Header().Set("content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(r)
}
