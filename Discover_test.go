package standard

import (
	"fmt"
	"net/http"
	"testing"
)

func Test1(t *testing.T){
	request, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	request.RemoteAddr = "127.0.0.1"
	headers := GetDiscoverHeader(request)
	fmt.Println(headers)

	request, _ = http.NewRequest("GET", "http://www.baidu.com", nil)
	request.RemoteAddr = "127.0.0.1"
	headers = GetDiscoverHeader(request)
	fmt.Println(headers)
}
