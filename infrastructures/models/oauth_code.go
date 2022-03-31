package models

import "encoding/json"

type OauthCode struct {
	ClientID    string `json:"client_id"`
	UserID      string `json:"user_id"`
	Scope       string `json:"scope"`
	RedirectURL string `json:"redirect_url"`
}

func (c *OauthCode) JSONString() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *OauthCode) BindFromJSON(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), c)
}
