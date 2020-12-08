package main

import "github.com/gofrs/uuid"

type ApiKey struct {
	Id            uuid.UUID
	Exchange      string
	UserId        string
	ApiKeyPreview string
	ApiKey        []byte
	Secret        []byte
}

func NewApiKey(userId string, exchange string, apiKeyPreview string, apiKey []byte, secret []byte) ApiKey {
	id, err := uuid.NewV4()
	if err != nil {
		panic("Failed to generate UUID")
	}
	return ApiKey{
		Id:            id,
		Exchange:      exchange,
		UserId:        userId,
		ApiKeyPreview: apiKeyPreview,
		ApiKey:        apiKey,
		Secret:        secret,
	}
}
