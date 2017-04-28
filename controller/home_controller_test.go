package controller

import (
	"log"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

//const DefaultRemoteAddr = "localhost:60006"

// 测试主页
func TestIndex(t *testing.T) {
	// 构造请求
	req := httptest.NewRequest("GET", "/", nil)
	// 构造响应 response 对象
	w := httptest.NewRecorder()
	// 执行 Index 请求
	Index(w, req)

	// 获取响应结果
	resp := w.Result()
	// 读取响应的 Body
	body, _ := ioutil.ReadAll(resp.Body)

	// 打印各个值
	log.Println(resp.StatusCode)
	log.Println(resp.Header.Get("Content-Type"))
	log.Println(string(body))
}

// 测试 GetName
func TestGetName(t *testing.T) {
	// 构造请求
	req := httptest.NewRequest("GET", "/name", nil)
	// 构造响应 response 对象
	w := httptest.NewRecorder()
	// 执行 Index 请求
	GetName(w, req)

	// 获取响应结果
	resp := w.Result()
	// 读取响应的 Body
	body, _ := ioutil.ReadAll(resp.Body)

	// 打印各个值
	log.Println(resp.StatusCode)
	log.Println(resp.Header.Get("Content-Type"))
	log.Println(string(body))
}