package main

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ApiKeyRepository interface {
	Save(ctx context.Context, apikey *ApiKey) (err error)
	Get(ctx context.Context, id uuid.UUID) (*ApiKey, error)
	FindByUser(ctx context.Context, userId string) ([]*ApiKey, error)
}

type apiKeyRepository struct {
	db *pgxpool.Pool
}

const (
	SqlTable      = "apikeys"
	SqlSave       = `INSERT INTO ` + SqlTable + ` (id, user_id, exchange, api_key_preview, api_key, secret) VALUES ($1, $2, $3, $4, $5, $6)`
	SqlGet        = `SELECT id, user_id, exchange, api_key_preview, api_key, secret FROM ` + SqlTable + ` WHERE id = $1`
	SqlFindByUser = `SELECT id, user_id, exchange, api_key_preview, api_key, secret FROM ` + SqlTable + ` WHERE user_id = $1`
)

func (apiKeyRepository *apiKeyRepository) Save(ctx context.Context, apikey *ApiKey) (err error) {
	_, err = apiKeyRepository.db.Exec(
		ctx,
		SqlSave,
		apikey.Id, apikey.UserId, apikey.Exchange, apikey.ApiKeyPreview, apikey.ApiKey, apikey.Secret)
	if err != nil {
		return err
	}

	return nil
}

func (apiKeyRepository *apiKeyRepository) Get(ctx context.Context, id uuid.UUID) (*ApiKey, error) {
	key := ApiKey{}
	err := pgxscan.Get(ctx, apiKeyRepository.db, &key, SqlGet, id)

	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (apiKeyRepository *apiKeyRepository) FindByUser(ctx context.Context, userId string) ([]*ApiKey, error) {
	var apiKeys []*ApiKey
	err := pgxscan.Select(ctx, apiKeyRepository.db, &apiKeys, SqlFindByUser, userId)
	if err != nil {
		return nil, err
	}
	return apiKeys, nil
}
