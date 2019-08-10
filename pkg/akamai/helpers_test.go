package akamai

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestXCacheKeyUnmarshal(t *testing.T) {
	want := &XCacheKey{
		SecureDeliveryIndicator: "S",
		TypeCode:                "L",
		Serial:                  1,
		CPCode:                  1,
		TTL:                     1 * time.Minute,
		FwdPath:                 "www.mockorig.com/it/donna",
		QString:                 "?test=1",
	}
	got := XCacheKeyUnmarshal("S/L/1/1/1m/www.mockorig.com/it/donna?test=1")
	assert.Equal(t, want, got)
}
