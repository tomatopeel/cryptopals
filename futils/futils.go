package futils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(local string, url string) {
	f, err := os.Create(local)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
