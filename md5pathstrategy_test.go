package md5_path_strategy

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

func TestMd5Path(t *testing.T) {
	m := &Md5PathStrategy{}
	filePaths := make(map[string]int)
	resultRoot := make(map[string]int)
	resultLevel2 := make(map[string]int)

	for i := 0; i < 10000000; i++ {
		data := RandStringBytesMaskImprSrc(1024)
		fullPath, _ := m.GeneratePath(data, "jpg")
		filePaths[fullPath]++
	}
	for key, _ := range filePaths {
		resultRoot[key[:2]]++
		resultLevel2[key[:5]]++
	}
	log.Println(resultRoot)
}
