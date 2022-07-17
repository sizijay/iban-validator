package request

type ValidateIBANRequest struct {
	IBAN string `json:"IBAN" validate:"required"`
}
