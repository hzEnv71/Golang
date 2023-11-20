package _test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello baidu"))
}

func TestConn(t *testing.T) {
	req := httptest.NewRequest("GET", "https://www.baidu.com/", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	bytes, _ := io.ReadAll(w.Result().Body)
	if string(bytes) != "hello baidu" {
		t.Fatal("expected hello baidu,but got", string(bytes))
	}
}
