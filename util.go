package codacytool

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func parseJSONFile(fileLocation string, out interface{}) error {
	file, err := os.Open(fileLocation)
	if err != nil {
		return err
	}

	defer file.Close()

	fileContentByte, err := ioutil.ReadAll(file)
	err = json.Unmarshal(fileContentByte, &out)
	if err != nil {
		return err
	}
	return nil
}
