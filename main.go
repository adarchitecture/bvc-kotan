package bvc_kotan

import (
	"fmt"
	"github.com/gogama/httpx"
)
func call_hhtp() {
	fmt.Println("Hola mundo")
	client := &httpx.Client{} // Use default retry and timeout policies
	client.Get("http://example.com")
}
