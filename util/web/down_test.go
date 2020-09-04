package web

import (
	"fmt"
	"testing"
)

func TestHttpFileServer(t *testing.T) {
	HttpFileServer(8082, "./static/www.ethereumx.xyz")
	select {

	}
}

func TestDown(t *testing.T) {
	isok := DownloadUrl("http://www.ethereumx.xyz/", "./static")
	fmt.Println(isok)
}
func TestUrl(t *testing.T) {
	fmt.Println("get ", GetUrlFileName("http://www.ethereumx.xyz/MyCrypto-ForETX-Setup1.7.9.exe"))
}
func TestUrlHost(t *testing.T) {
	fmt.Println("get ", GetUrlHost("https://www.ethereumx.xyz/"))
}

func TestGetUrlContent(t *testing.T) {
	content := GetUrlContent("https://www.ethereumx.xyz/")

	if len(content) > 0 {

		linkMap := GetCssList(content)

		fmt.Println(linkMap)

	}
}
func TestGetUrlScriptContent(t *testing.T) {
	content := GetUrlContent("https://www.ethereumx.xyz/")

	if len(content) > 0 {

		linkMap := GetScriptsList(content)

		fmt.Println(linkMap)

	}
}

func TestDownloadWebsite(t *testing.T) {

	isok := DownloadWebsite("http://www.ethereumx.xyz/", "./static")
	fmt.Println(isok)

}
