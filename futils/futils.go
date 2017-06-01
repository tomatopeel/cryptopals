package futils

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadLinesHexFile(localFile, remoteFile string) (result [][]byte, err error) {
	file, err := os.Open(localFile)
	if err != nil {
		err = downloadFile(localFile, remoteFile)
		if err != nil {
			return nil, err
		}
		file, err = os.Open(localFile)
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, hex.DecodedLen(len(bytes)))
		_, err = hex.Decode(line, bytes)
		if err != nil {
			return nil, err
		}
		result = append(result, line)
	}
	return
}

func ReadAllHexFile(localFile, remoteFile string) ([]byte, error) {
	file, err := os.Open(localFile)
	if err != nil {
		err = downloadFile(localFile, remoteFile)
		if err != nil {
			return nil, err
		}
		file, err = os.Open(localFile)
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := info.Size()

	bytes := make([]byte, size)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	_, err = hex.Decode(bytes, data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ReadAllBase64File(localFile, remoteFile string) ([]byte, error) {
	file, err := os.Open(localFile)
	if err != nil {
		err = downloadFile(localFile, remoteFile)
		if err != nil {
			return nil, err
		}
		file, err = os.Open(localFile)
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()

	reader := base64.NewDecoder(base64.StdEncoding, file)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func downloadFile(local string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(local)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
