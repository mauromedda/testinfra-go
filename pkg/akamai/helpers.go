package akamai

import (
	"errors"
	"regexp"
	"strconv"
)

// XCacheKey is the Response Header "X-Cache-Key" data type. It does consist
// of the following Cache Key Component
type XCacheKey struct {
	SecureDeliveryIndicator string
	TypeCode                string
	Serial                  int
	CPCode                  int
	TTL                     string
	FwdPath                 string
	QString                 string
}

var (
	// ErrNoValideXCacheKey it's the error associated to not valid X-Cache-Key Value
	ErrNoValideXCacheKey = errors.New("X-Cache-Key value not valid")
	reHeader             = regexp.MustCompile(`^(S)?/(L)/(\d+)/(\d+)/(\d+[msh]+)/([^?]+)([?]?[^?]+)*$`)
)

// XCacheKeyUnmarshal map the X-Cache-Key Header into the XCacheKey data type.
func XCacheKeyUnmarshal(xck string) (*XCacheKey, error) {
	matched := reHeader.FindStringSubmatch(xck)
	if len(matched) == 0 {
		return &XCacheKey{}, ErrNoValideXCacheKey

	}
	matched = matched[1:len(matched)]
	serial, _ := strconv.Atoi(matched[2])
	cpcode, _ := strconv.Atoi(matched[3])
	return &XCacheKey{
		SecureDeliveryIndicator: matched[0],
		TypeCode:                matched[1],
		Serial:                  serial,
		CPCode:                  cpcode,
		TTL:                     matched[4],
		FwdPath:                 matched[5],
		QString:                 matched[6],
	}, nil

}

// GetCP returns the CP code ID of the Akamai Cache configuration
func (xk *XCacheKey) GetCP() int {
	return xk.CPCode
}

// GetOrigin returns the Forward Path of the Akamai Cache configuration
func (xk *XCacheKey) GetOrigin() string {
	return xk.FwdPath
}

// GetSerial returns the Serial of the Akamai Cache configuration
func (xk *XCacheKey) GetSerial() int {
	return xk.Serial
}
