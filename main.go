package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

type Options struct {
	urls string
	output string
}

var opts Options

func init() {
	flag.StringVar(&(opts.urls), "urls", "", "defines paths of the downloading")
	flag.StringVar(&(opts.output), "output", "output.zip", "defines path where stores result")
}


func ValidateUrls(str string) ([]string, error){
	urls := strings.Split(str, ",")
	validUrls := make([]string, 0)
	pattern := `^https?:\/\/`
    re := regexp.MustCompile(pattern)
	for i, url := range urls {
		matched := re.MatchString(url)
		if !matched {
			fmt.Printf("url %d is not valid: %s It must begin from http(s):// \n", i, url)
			continue
		}
		validUrls = append(validUrls, url)
	}
	if len(validUrls) == 0 {
		return validUrls, errors.New("no one urls are not valid")
	}
	return validUrls, nil
}

func ValidateOutput(output string) (string, error) {
	_, err := os.Stat(output)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(output, os.ModePerm); err != nil {
			return "", err
		}
		return output, nil
	}
	return "", err
}

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

func DownloadFiles(urls []string, output string) {
	for i := range urls {
		if err := DownloadFile(urls[i], output); err != nil {
			fmt.Printf("Url %d : %s does not download. Error: %s", i + 1, urls[i], err.Error())
		}
	}
}

func main() {
	flag.Parse()
	urls, err := ValidateUrls(opts.urls)
	if err != nil {
		panic(err.Error())
	}
	output, err := ValidateOutput(opts.output)
	if err != nil {
		panic(err.Error())
	}
	DownloadFiles(urls, output)
}