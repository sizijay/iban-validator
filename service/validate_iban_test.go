package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sizijay/iban-validator/config"
	"github.com/sizijay/iban-validator/internal/http/request"
	"testing"
)

var (
	validIBANTestArray = []struct {
		iban string
	}{
		{"GB82 WEST 1234 5698 7654 32"},
	}
	inValidIBANTestArray = []struct {
		iban string
	}{
		{"GB82 WEST 1234 5698 7654 33"},
	}
)

func TestValidIBAN(t *testing.T) {

	tCtx := context.WithValue(context.Background(), "uuid", uuid.New())
	ctx, cancel := context.WithCancel(tCtx)
	defer cancel()

	validateIBANService := ValidateIBANService{}

	config.IBANCountryConfigMap["GB"] = config.CountryData{
		CountryCode: "GB",
		Country:     "United Kingdom",
		IBANLength:  22,
	}

	for _, value := range validIBANTestArray {
		res, err := validateIBANService.Do(ctx, []interface{}{request.ValidateIBANRequest{IBAN: value.iban}})
		if len(res) == 0  {
			t.Error("Error: TestValidIBAN test case failed. Response length should be > 0")
		}

		isValid := res[0].(bool)

		if !isValid  {
			t.Error("Error: TestValidIBAN test case failed. IBAN should be identified as valid")
		}

		if err != nil {
			t.Error("Error: TestValidIBAN test case failed. Error should be nil")
			t.Log(err)
		}
	}
}


func TestInvalidIBAN(t *testing.T) {

	tCtx := context.WithValue(context.Background(), "uuid", uuid.New())
	ctx, cancel := context.WithCancel(tCtx)
	defer cancel()

	validateIBANService := ValidateIBANService{}

	config.IBANCountryConfigMap["GB"] = config.CountryData{
		CountryCode: "GB",
		Country:     "United Kingdom",
		IBANLength:  22,
	}

	for _, value := range inValidIBANTestArray {
		res, _ := validateIBANService.Do(ctx, []interface{}{request.ValidateIBANRequest{IBAN: value.iban}})
		if len(res) == 0  {
			t.Error("Error: TestValidIBAN test case failed. Response length should be > 0")
		}

		isValid := res[0].(bool)

		if isValid  {
			t.Error("Error: TestValidIBAN test case failed. IBAN should be identified as invalid")
		}
	}
}