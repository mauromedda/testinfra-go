package akamai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXCacheKeyUnmarshal(t *testing.T) {
	want := &XCacheKey{
		SecureDeliveryIndicator: "S",
		TypeCode:                "L",
		Serial:                  1,
		CPCode:                  1,
		TTL:                     "1m",
		FwdPath:                 "www.mockorig.com/it/donna",
		QString:                 "?test=1",
	}

	t.Run("Check X-Cache-Key Vaidation", func(t *testing.T) {
		_, err := XCacheKeyUnmarshal("S/L/1/1/1/www.mockorig.com?test=1?test=1")
		assert.Equal(t, ErrNoValideXCacheKey, err)
	})
	t.Run("Map X-Cache-Key value to XCacheKey struct", func(t *testing.T) {
		got, err := XCacheKeyUnmarshal("S/L/1/1/1m/www.mockorig.com/it/donna?test=1")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, want, got)
	})
}

func TestGetCP(t *testing.T) {
	expectedCP := 1
	got := (&XCacheKey{
		SecureDeliveryIndicator: "S",
		TypeCode:                "L",
		Serial:                  1,
		CPCode:                  1,
		TTL:                     "1m",
		FwdPath:                 "www.mockorig.com/it/donna",
		QString:                 "?test=1",
	}).GetCP()
	assert.Equal(t, expectedCP, got)
}

func TestGetOrigin(t *testing.T) {
	expectedFwdPath := "www.mockorig.com/it/donna"
	got := (&XCacheKey{
		SecureDeliveryIndicator: "S",
		TypeCode:                "L",
		Serial:                  1,
		CPCode:                  1,
		TTL:                     "1m",
		FwdPath:                 "www.mockorig.com/it/donna",
		QString:                 "?test=1",
	}).GetOrigin()
	assert.Equal(t, expectedFwdPath, got)
}

func TestGetSerial(t *testing.T) {
	expectedSerial := 1
	got := (&XCacheKey{
		SecureDeliveryIndicator: "S",
		TypeCode:                "L",
		Serial:                  1,
		CPCode:                  1,
		TTL:                     "1m",
		FwdPath:                 "www.mockorig.com/it/donna",
		QString:                 "?test=1",
	}).GetSerial()
	assert.Equal(t, expectedSerial, got)
}
