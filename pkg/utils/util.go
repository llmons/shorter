package utils

import (
	"log"
	"sync"

	"github.com/speps/go-hashids"
)

// 单例 hashid 实例，避免每次重新创建
var (
	hashID    *hashids.HashID
	once      sync.Once
	salt      = "llmons" // 可以从配置读取
	minLength = 6        // 设置最小长度
)

// 初始化 HashID 实例（只初始化一次）
func initHashID() {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength

	var err error
	hashID, err = hashids.NewWithData(hd)
	if err != nil {
		log.Fatalf("Failed to initialize hashids: %v", err)
	}
}

// HashEncode 将 int64 编码为 hashid 字符串
func HashEncode(num int64) string {
	once.Do(initHashID)
	hash, err := hashID.EncodeInt64([]int64{num})
	if err != nil {
		log.Printf("HashEncode failed: %v", err)
		return ""
	}
	return hash
}

// HashDecode 将 hashid 字符串还原为原始 int64
func HashDecode(str string) int64 {
	once.Do(initHashID)
	decoded, err := hashID.DecodeInt64WithError(str)
	if err != nil || len(decoded) == 0 {
		log.Printf("HashDecode failed: %v", err)
		return -1
	}
	return decoded[0]
}
