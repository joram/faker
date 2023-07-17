package faker

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
)

var globalSeed int64 = 0

func SetSeed(seed int64) {
	globalSeed = seed
}

func calculateSeed(s string, seed int64) int64 {
	var stringSeed int64
	hashBytes := sha1.Sum([]byte(s))
	buf := bytes.NewBuffer(hashBytes[:])
	binary.Read(buf, binary.LittleEndian, &stringSeed)
	return seed + globalSeed + stringSeed
}
