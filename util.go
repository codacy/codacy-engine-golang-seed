package codacytool

import (
	"io/ioutil"
	"os"
)

func readFile(fileLocation string) ([]byte, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
