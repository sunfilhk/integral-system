package utils

import (
	"fmt"
	"github.com/ipipdotnet/datx-go"
)

type Location struct {
	Country       string
	Province      string
	City          string
	Organization  string
	ISP           string
	Latitude      string
	Longitude     string
	TimeZone      string
	TimeZone2     string
	CityCode      string
	PhonePrefix   string
	CountryCode   string
	ContinentCode string
	IDC           string // IDC | VPN
	BaseStation   string // WIFI | BS (Base Station)
	CountryCode3  string
	EuropeanUnion bool
	CurrencyCode  string
	CurrencyName  string
	Anycast       bool
}

func GetCityUsers(ip string) (string, error) {
	city, err := datx.NewCity("./resources/mydata4vipday2.datx") // For City Level IP Database
	if err != nil {
		return "", err
	}

	loc, err := city.FindLocation(ip)
	if err != nil {
		return "", err
	}
	ipaddr := fmt.Sprintf("%s %s %s %s %s", loc.Country, loc.Province, loc.City, loc.Organization, loc.ISP)
	return ipaddr, nil
}
