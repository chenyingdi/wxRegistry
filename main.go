package main

import (
	"crypto/sha1"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"log"
	"net/http"
	"sort"
	"unsafe"
)

// 微信注册校验
func WxRegistry(r *ghttp.Request) {
	var (
		lst       []string
		temp      string
		result    []byte
		resStr    *string
		sign      = r.GetString("signature")
		timeStamp = r.GetString("timestamp")
		nonce     = r.GetString("nonce")
		echoStr   = r.GetString("echostr")
		h         = sha1.New()
	)

	log.Println("sign: ", sign)
	log.Println("timestamp: ", timeStamp)
	log.Println("nonce: ", nonce)
	log.Println("echo: ", echoStr)

	lst = append(lst, sign, timeStamp, nonce)

	// 1. 字典序排序
	sort.Strings(lst)

	// 2. 拼接成字符串
	for _, s := range lst {
		temp += s
	}

	log.Println("temp: ", temp)

	// 3. sha1加密
	h.Write([]byte(temp))
	result = h.Sum(nil)

	resStr = (*string)(unsafe.Pointer(&result))

	log.Println("res: ", resStr)

	// 4. 与sign比对
	if sign != *resStr {
		log.Println("sign is not equal with resStr!!")
		r.Response.WriteStatusExit(http.StatusInternalServerError)
		return
	}

	// 5. 原样返回echo字符串内容
	r.Response.WriteStatusExit(http.StatusOK, echoStr)
}

func main() {
	s := g.Server()
	s.BindHandler("/", WxRegistry)
	s.Run()
}
