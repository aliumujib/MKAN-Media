package models

import (
	"encoding/json"
	"fmt"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func (receiver TokenResponse) ToJson() ([]byte, *error) {
	data, err := json.Marshal(receiver)
	return data, &err
}

func (receiver TokenResponse) TokenResponseFromJson(tokenJson string) (TokenResponse, *error) {
	fmt.Println("Json:", tokenJson)
	err := json.Unmarshal([]byte(tokenJson), &receiver)
	fmt.Println("Marshalled", receiver)
	return receiver, &err
}
