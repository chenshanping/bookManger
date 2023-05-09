package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

var (
	allRandomStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ToolsUtil    = toolsUtil{}
)

type toolsUtil struct {
}

func (tu toolsUtil) RandomString(length int) string {
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = allRandomStr[rand.Intn(62)]
	}
	return string(byteList)
}

// MakeUuid 制作uuid
func (tu toolsUtil) MakeUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// MakeMd5 制作MD5
func (tu toolsUtil) MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

func (tu toolsUtil) Paging(pageSize int, pageNum int) int {
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	offSetval := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSetval = -1
	}
	return offSetval
}
