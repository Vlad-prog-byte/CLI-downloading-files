package downloader

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"strings"
	"path"
)

func getFilePath(url, output string) string{
	partsUrl := strings.Split(url, "/")
	fileName := partsUrl[len(partsUrl) - 1]
	return path.Join(output, fileName)
}

func DownloadFile(url, output string) error{
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filepath := getFilePath(url, output)
	// fmt.Printf("filePath is %v | url is %v", filepath, output)
	fs, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fs.Close()
	_, err = io.Copy(fs, resp.Body)
	if err != nil {
		return err
	}
	return nil
}


func Run(urls, output string) {
	validUrls, err := ValidateUrls(urls)
	if err != nil {
		panic(err.Error())
	}

	validOutput, err := ValidateOutput(output)
	if err != nil {
		panic(err.Error())
	}

	for i := range validUrls {
		if err := DownloadFile(validUrls[i], validOutput); err != nil {
			fmt.Printf("Url %d : %v does not download. Error: %s", i + 1, urls[i], err.Error())
		}
	}
}