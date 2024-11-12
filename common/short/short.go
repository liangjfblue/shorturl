package short

import (
	"common/utils"
	"strings"
)

type Short struct {
	id *utils.SnowFlake
}

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewShort(id *utils.SnowFlake) *Short {
	return &Short{id: id}
}

func (s Short) Create() string {
	id := s.id.NextId()

	if id == 0 {
		return string(base62Chars[0])
	}

	base := int64(len(base62Chars))
	var result []byte
	for id > 0 {
		remainder := id % base
		result = append(result, base62Chars[remainder])
		id /= base
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func (s Short) ToSnowFlakeID(str string) int64 {
	var result int64
	base := int64(len(base62Chars))
	for i := 0; i < len(str); i++ {
		result *= base
		result += int64(strings.IndexByte(base62Chars, str[i]))
	}
	return result
}
