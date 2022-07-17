package entities

type IBAN struct {
	Value       string
	CountryCode string
	CheckDigits string
	Length      int
	BBAN        string
}
