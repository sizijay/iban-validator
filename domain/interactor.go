package domain

import "context"

type Service interface {
	Do(ctx context.Context, arg []interface{}) ([]interface{}, error)
}

const (
	MAX_IBAN_LENGTH = 34
	MIN_IBAN_LENGTH = 15
	IBAN_MOD_NUMBER = 97
	ASCII_VALUE_A 	= 65
	ASCII_VALUE_Z 	= 90
)