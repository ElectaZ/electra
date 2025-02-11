package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProducerbillsKeyPrefix is the prefix to retrieve all Producerbills
	ProducerbillsKeyPrefix = "Producerbills/value/"
)

// ProducerbillsKey returns the store key to retrieve a Producerbills from the index fields
func ProducerbillsKey(
	billCycleID uint64,
	producerdeviceID string,
) []byte {
	var key []byte

	billCycleIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(billCycleIDBytes, billCycleID)
	key = append(key, billCycleIDBytes...)
	key = append(key, []byte("/")...)

	producerdeviceIDBytes := []byte(producerdeviceID)
	key = append(key, producerdeviceIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
