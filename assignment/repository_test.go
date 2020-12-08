package main

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCtx = context.Background()

func TestApiKeyRepositorySave(t *testing.T) {
	defer CleanUpDb()
	apikeyRepository := &apiKeyRepository{db: TestDbPool}
	expectedApikey := NewApiKey("u", "e", "previ", []byte("k"), []byte("s"))
	err := apikeyRepository.Save(testCtx, &expectedApikey)
	assert.Nil(t, err)

	apiKey := ApiKey{}
	err = pgxscan.Get(testCtx, TestDbPool, &apiKey, SqlGet, expectedApikey.Id)
	assert.Nil(t, err)
	assert.True(t, cmp.Equal(expectedApikey, apiKey))
}

func TestApiKeyRepositoryGet(t *testing.T) {
	defer CleanUpDb()
	apikeyRepository := &apiKeyRepository{db: TestDbPool}

	expectedApikey := NewApiKey("u", "e", "previ", []byte("k"), []byte("s"))
	_, err := TestDbPool.Exec(
		testCtx,
		SqlSave,
		expectedApikey.Id, expectedApikey.UserId, expectedApikey.Exchange, expectedApikey.ApiKeyPreview, expectedApikey.ApiKey, expectedApikey.Secret)
	assert.Nil(t, err)
	apiKey, err := apikeyRepository.Get(testCtx, expectedApikey.Id)
	assert.Nil(t, err)
	assert.True(t, cmp.Equal(&expectedApikey, apiKey))
}

func TestFindByUserGet(t *testing.T) {
	defer CleanUpDb()
	apikeyRepository := &apiKeyRepository{db: TestDbPool}
	const u1 = "u1"
	const u2 = "u2"
	expectedApikey := NewApiKey(u1, "e", "previ", []byte("k"), []byte("s"))
	_ = apikeyRepository.Save(testCtx, &expectedApikey)
	expectedApikey = NewApiKey(u2, "e1", "previ", []byte("k"), []byte("s"))
	_ = apikeyRepository.Save(testCtx, &expectedApikey)
	expectedApikey = NewApiKey(u2, "e2", "previ", []byte("k"), []byte("s"))
	_ = apikeyRepository.Save(testCtx, &expectedApikey)
	keys, err := apikeyRepository.FindByUser(testCtx, u1)
	assert.Nil(t, err)
	assert.Len(t, keys, 1)
	assert.Equal(t, keys[0].UserId, u1)
	assert.Equal(t, keys[0].Exchange, "e")
	assert.Equal(t, keys[0].ApiKeyPreview, "previ")
	assert.Equal(t, keys[0].ApiKey, []byte("k"))
	assert.Equal(t, keys[0].Secret, []byte("s"))

	keys, err = apikeyRepository.FindByUser(testCtx, u2)
	assert.Len(t, keys, 2)
	assert.Equal(t, keys[0].UserId, u2)
	assert.Equal(t, keys[1].UserId, u2)

}
