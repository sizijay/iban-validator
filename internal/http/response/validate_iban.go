package response

type ValidateIBANResponse struct {
	Data struct {
		IsValidIBAN bool `json:"IsValidIBAN"`
	} `json:"Data"`
}
