package downloader
import (
	"strings"
	"fmt"
	"regexp"
	"os"
	"errors"
)

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
			return output, err
		}
		return output, nil
	}
	if err == nil {
		return output, nil
	}
	return "", err
}
