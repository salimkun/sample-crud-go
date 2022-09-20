package util

import (
	"fmt"
	"net/mail"
	"net/url"
	"time"
)

func ValidateMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return address, false
	}
	return addr.Address, true
}

func ValidateUrl(uri string) bool {
	sd, err := url.ParseRequestURI(uri)

	fmt.Println("aerr ", err, " s ", sd)
	return err == nil
}

func ValidateDate(req string) bool {
	_, err := time.Parse(time.RFC3339, req)
	return err == nil
}
