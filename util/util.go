package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func WriteFile(filePath, filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	configFile := path.Join(filePath, filename)
	return ioutil.WriteFile(configFile, file, 0755)
}

func Scan(prompt string) string {
	prompt = strings.TrimRight(prompt, " ")
	fmt.Printf(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CSVToArray(csv string) []string {
	if strings.Trim(csv, " ") == "" {
		return []string{}
	}
	arr := strings.Split(csv, ",")
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.Trim(arr[i], " ")
	}
	return arr
}
