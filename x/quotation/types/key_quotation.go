package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// QuotationKeyPrefix is the prefix to retrieve all Quotation
	QuotationKeyPrefix = "Quotation/value/"
)

// QuotationKey returns the store key to retrieve a Quotation from the index fields
func QuotationKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
