package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type CountryData struct {
	CountryCode 	string 	`yaml:"code"`
	Country			string 	`yaml:"country"`
	IBANLength 		int 	`yaml:"length"`
}

var (
	IBANCountryConfigMap = make(map[string]CountryData)
	iBANCountryArray     []CountryData
)


func LoadConfigurations() {
	yamlData, err := ioutil.ReadFile("yaml/country_data.yaml")
	if err != nil {
		log.Fatalf(`Unable to load country_data from yaml. Error:%+v`, err)
	}

	err = yaml.Unmarshal(yamlData, &iBANCountryArray)
	if err != nil {
		log.Fatalf(`Unable to unmarshal country_data. Error:%+v`, err)
	}

	for _, conf := range iBANCountryArray {
		IBANCountryConfigMap[conf.CountryCode] = CountryData{
			CountryCode: conf.CountryCode,
			Country:     conf.Country,
			IBANLength:  conf.IBANLength,
		}
	}

	fmt.Println(fmt.Sprintf("%v country IBAN records were successfully fetched from the config", len(iBANCountryArray)))
}