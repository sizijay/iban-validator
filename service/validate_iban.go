package service

import (
	"context"
	"fmt"
	"github.com/sizijay/iban-validator/config"
	"github.com/sizijay/iban-validator/domain"
	"github.com/sizijay/iban-validator/domain/entities"
	"github.com/sizijay/iban-validator/domain/errors"
	requestEntity "github.com/sizijay/iban-validator/internal/http/request"
	"runtime/debug"
	"strconv"
	"strings"
)

type ValidateIBANService struct{}

func (i *ValidateIBANService) Do(ctx context.Context, arg []interface{}) ([]interface{}, error) {
	request, ok := arg[0].(requestEntity.ValidateIBANRequest)
	if !ok {
		fmt.Println(fmt.Sprintf("IBAN Service: Error decoding the IBAN number."))

		return []interface{}{nil}, errors.NewApplicationError("Error Decoding the IBAN Value", errors.ErrDecodeIBAN, string(debug.Stack()))
	}

	var IBAN entities.IBAN
	var isValidIBAN bool
	var countryConfig config.CountryData

	// check if the IBAN is valid --------------------------------------------------------------------------------------

	// format IBAN
	iban := strings.ToUpper(strings.Replace(request.IBAN, " ", "", -1))

	// check if the IBAN length is within the range
	if len(iban) < domain.MIN_IBAN_LENGTH || len(iban) > domain.MAX_IBAN_LENGTH {
		fmt.Println(fmt.Sprintf("IBAN Service: Invalid IBAN length: %v characters", len(iban)))
		return []interface{}{isValidIBAN}, errors.NewDomainError("Invalid IBAN. Length should be between 5-34", errors.ErrValidateIBAN, string(debug.Stack()))
	}

	IBAN.Value = iban
	IBAN.Length = len(iban) //min 5 max 34
	IBAN.CountryCode = iban[0:2]
	IBAN.BBAN = iban[4:IBAN.Length]
	IBAN.CheckDigits = iban[2:4] //max 30

	// 01 - validate the IBAN length with respect to the country code
	countryConfig, ok = config.IBANCountryConfigMap[IBAN.CountryCode]
	if !ok {
		fmt.Println(fmt.Sprintf("IBAN Service: Country is not configured for the CountryID: %v", IBAN.CountryCode))
		return []interface{}{nil}, errors.NewDomainError("Unable to validate. Country is not yet configured", errors.ErrNoCountryConfig, string(debug.Stack()))
	}

	if countryConfig.IBANLength != IBAN.Length {
		fmt.Println(fmt.Sprintf("IBAN Service: IBAN Length is not matching with the Country-%v. CountryCode:%v | Length:%v | ExpectedLength:%v", countryConfig.Country, IBAN.CountryCode, IBAN.Length, countryConfig.IBANLength))
		return []interface{}{isValidIBAN}, nil
	} else {
		// 02,03,04 re-arranging the IBAN
		reArrangedIBAN := IBAN.BBAN + IBAN.CountryCode + "00" // check digits are replaced with "00"

		// 05 - converting the re-arranged IBAN to integer (replace alpha characters)
		intOnlyIBAN := replaceIBANCharacters(reArrangedIBAN)

		// 06 - compute remainder - modulo by 97
		modRemainder := computeRemainder(intOnlyIBAN)
		if modRemainder < 0 {
			fmt.Println(fmt.Sprintf("IBAN Service: Error while computing remainder for IBAN converted to int: %v", intOnlyIBAN))
			return []interface{}{isValidIBAN}, errors.NewApplicationError("Error while computing the mod value", errors.ErrComputeMod, string(debug.Stack()))
		}

		// 07 - validate the check digits after subtracting from 98
		strModRem := strconv.Itoa(98 - modRemainder)


		// add additional 0 if remainder is a single digit
		if len(strModRem) == 1 {
			strModRem = "0" + strModRem
		}

		if strModRem != IBAN.CheckDigits {
			fmt.Println(fmt.Sprintf("IBAN Service: Error: Check digits is not a match with the remainder. CheckDigits:%v | Remainder:%v", IBAN.CheckDigits, strModRem))
			return []interface{}{isValidIBAN}, nil
		}
	}

	// IBAN check passed
	return []interface{}{true}, nil
}

// replace char values with int values starting from A=10
func replaceIBANCharacters(iban string) string {
	runes := []rune(iban)

	intIBAN := ""

	for i := 0; i < len(runes); i++ {
		val := int(runes[i])
		if val >= domain.ASCII_VALUE_A && val <= domain.ASCII_VALUE_Z {
			intIBAN += strconv.Itoa(int(runes[i]) - 55)
			continue
		}
		intIBAN += string(runes[i])
	}

	return intIBAN
}

// compute remainder (mod 97)
func computeRemainder(iban string) int {
	var tempRemainder int

	tempIBAN := ""

	for true {
		if len(iban) <= 9 {
			break
		}

		tempIBAN = iban[0:9]

		// convert first 9 digits to a single number
		val, err := strconv.Atoi(tempIBAN)
		if err != nil {
			return -1
		}

		// get the remainder and process with the rest
		tempRemainder = val % domain.IBAN_MOD_NUMBER
		iban = strconv.Itoa(tempRemainder) + iban[9:]
	}

	// convert directly
	val, err := strconv.Atoi(iban)
	if err != nil {
		return -1
	}

	return val % domain.IBAN_MOD_NUMBER
}