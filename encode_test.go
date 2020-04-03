package main

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"testing"
)

// b235e858a4e792a7525972e2103f13d737b7b8b8
func TestEncode(t *testing.T) {
	timestamp := "1585878217"
	nonce := "505210980"
	token := "cyd19960424"

	lst := []string{timestamp, nonce, token}

	// 1. 字典序排序
	sort.Strings(lst)

	t.Log("lst: ", lst)

	// 2. 拼接字符串
	var str string
	for _, s := range lst {
		str += s
	}

	// 3. sha1加密
	h := sha1.New()
	h.Write([]byte(str))
	res := h.Sum(nil)

	t.Log("result: ", fmt.Sprintf("%x", res))
}
