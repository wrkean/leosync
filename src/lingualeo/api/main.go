package api

import (
	"net/http/cookiejar"
)

func NewClient(email string, password string) (Client, error) {
	cookieJar, _ := cookiejar.New(nil)
	client := Client{
		cookie: cookieJar,
	}
	errs := client.authorize(email, password)

	return client, errs
}
