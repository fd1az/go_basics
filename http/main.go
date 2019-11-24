package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

//Custom Write implementation
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Total bytes", len(bs))
	return len(bs), nil
}
