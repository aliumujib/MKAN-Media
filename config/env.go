package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type env struct {
	ENV                       string
	PORT                      string
	MONGO_URL                 string
	REDIS_URL                 string
	SOUND_CLOUD_CLIENT_ID     string
	SOUND_CLOUD_CLIENT_SECRET string
}

func (e env) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.ENV, validation.Required),
		validation.Field(&e.PORT, validation.Required),
		validation.Field(&e.MONGO_URL, validation.Required),
		validation.Field(&e.REDIS_URL, validation.Required),
		validation.Field(&e.SOUND_CLOUD_CLIENT_ID, validation.Required),
		validation.Field(&e.SOUND_CLOUD_CLIENT_SECRET, validation.Required),
	)
}
