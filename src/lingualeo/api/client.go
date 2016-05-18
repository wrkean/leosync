package api

import (
	"net/http"
	"strings"

	"github.com/franela/goreq"
)

// Client is a LinguaLeo API client
type Client struct {
	cookie http.CookieJar
}

func (c Client) get(url string, requestData interface{}, result interface{}) error {
	resp, err := goreq.Request{
		Uri:         url,
		QueryString: requestData,
		CookieJar:   c.cookie,
	}.Do()

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return newWrongResponseStatusError(resp.Status)
	}

	if err := resp.Body.FromJsonTo(&result); err != nil {
		return err
	}

	return err
}

func (c Client) authorize(email string, password string) error {
	var err error

	req := LoginRequest{
		Email:    email,
		Password: password,
	}

	var loginResp LoginResponse
	if err := c.get(loginURL, req, &loginResp); err != nil {
		return err
	}

	if strings.TrimSpace(loginResp.ErrorMsg) != "" {
		err = newResponseError(loginResp.ErrorMsg)
	}

	return err
}

func (c Client) validateCredentials(email string, password string) error {
	var err error

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		err = newEmptyCredentialsError()
	}
	return err
}

// GetTranslations returns translations of a word
func (c Client) GetTranslations(word string) (Word, error) {
	req := TranslationRequest{
		Word: word,
	}
	var err error

	translations := Word{}

	if err = c.get(translateURL, req, &translations); err != nil {
		return translations, err
	}

	if strings.TrimSpace(translations.ErrorMsg) != "" {
		err = newResponseError(translations.ErrorMsg)
	}

	return translations, err
}

// AddWord adds word with translation to LinguaLeo
func (c Client) AddWord(word, translation string) (Word, error) {
	req := AddWordRequest{
		Word:        word,
		Translation: translation,
	}

	var err error

	var result Word

	if err = c.get(addWordURL, req, &result); err == nil {
		return result, err
	}
	if strings.TrimSpace(result.ErrorMsg) != "" {
		err = newResponseError(result.ErrorMsg)
	}

	return result, err
}
