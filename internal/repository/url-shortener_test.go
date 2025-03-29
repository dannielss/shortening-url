package repository

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRepo_StoreShortURL(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()
	repo := NewRepo(client)

	mockRedis.ExpectSet("short123", "http://longurl.com", 0).SetVal("OK")

	err := repo.StoreShortURL("short123", "http://longurl.com")

	assert.NoError(t, err)

	assert.NoError(t, mockRedis.ExpectationsWereMet())
}

func TestRepo_StoreShortURLError(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()
	repo := NewRepo(client)

	mockRedis.ExpectSet("short123", "http://longurl.com", 0).SetErr(errors.New("error"))

	err := repo.StoreShortURL("short123", "http://longurl.com")

	assert.Error(t, err)
}

func TestRepo_GetLongURL(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()

	repo := NewRepo(client)

	mockRedis.ExpectGet("short123").SetVal("http://longurl.com")

	result, err := repo.GetLongURL("short123")

	assert.NoError(t, err)
	assert.Equal(t, "http://longurl.com", result)

	assert.NoError(t, mockRedis.ExpectationsWereMet())
}

func TestRepo_GetLongURLError(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()

	repo := NewRepo(client)

	mockRedis.ExpectGet("short123").SetErr(errors.New("error"))

	_, err := repo.GetLongURL("short123")

	assert.Error(t, err)
}

func TestRepo_GetShortURL(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()

	repo := NewRepo(client)

	mockRedis.ExpectGet("http://longurl.com").SetVal("short123")

	result := repo.GetShortURL("http://longurl.com")

	assert.Equal(t, "short123", result)

	assert.NoError(t, mockRedis.ExpectationsWereMet())
}

func TestRepo_GetShortURLError(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()

	repo := NewRepo(client)

	mockRedis.ExpectGet("http://longurl.com").SetErr(errors.New("error"))

	result := repo.GetShortURL("http://longurl.com")

	assert.Equal(t, "", result)

}

func TestRepo_GetLongURLNotFound(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()

	repo := NewRepo(client)

	mockRedis.ExpectGet("short123").SetErr(redis.Nil)

	result, err := repo.GetLongURL("short123")

	assert.EqualError(t, err, "short URL not found")
	assert.Empty(t, result)

	assert.NoError(t, mockRedis.ExpectationsWereMet())
}
