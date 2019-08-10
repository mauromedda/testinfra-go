package akamai

import (
	"time"
)

// XCacheKey is the Response Header "X-Cache-Key" data type. It does consist
// of the following Cache Key Component
type XCacheKey struct {
	SecureDeliveryIndicator string
	TypeCode                string
	Serial                  int
	CPCode                  int
	TTL                     time.Duration
	FwdPath                 string
	QString                 string
}

// XCacheKeyUnmarshal map the X-Cache-Key Header into the XCacheKey data type.
func XCacheKeyUnmarshal(xck string) *XCacheKey {
	return &XCacheKey{
		SecureDeliveryIndicator: "S",
		TypeCode:                "L",
		Serial:                  1,
		CPCode:                  1,
		TTL:                     1 * time.Minute,
		FwdPath:                 "www.mockorig.com/it/donna",
		QString:                 "?test=1",
	}
}
