package content

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func FileContent(FileName string) ([]string, error) {
	Content := []string{}
	_, err := os.Stat(FileName)
	if err != nil {
		return Content, err
	}
	inputFile, inputError := os.Open(FileName)
	if inputError != nil {
		return Content, inputError
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')

		if readerError == io.EOF {
			break
		}

		if readerError != nil {
			return Content, readerError
		}
		Content = append(Content, strings.TrimSpace(inputString))
	}
	return Content, nil
}

func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}
