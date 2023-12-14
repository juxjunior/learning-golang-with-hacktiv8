package util

import (
	"assignment3/model"
	"encoding/json"
	"io/ioutil"
)

func SaveToJSON(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromJSON(filename string) ([]model.WaterStatus, error) {
	var data []model.WaterStatus

	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
