package service

import (
	"assignment3/model"
	"assignment3/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateData(c echo.Context) error {
	log.Println("--> Execute update data water & wind")

	water := model.WaterStatus{
		Water: rand.Intn(15) + 1,
		Wind:  rand.Intn(15) + 1,
	}

	// Load existing data
	existingData, err := util.ReadFromJSON("data.json")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to load data"})
	}

	// Append new data
	existingData = append(existingData, water)

	// Save data to JSON file
	err = util.SaveToJSON("data.json", existingData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save data"})
	}

	// Unmarshal the JSON array
	var wotah []model.WaterStatus
	fileContent, err := ioutil.ReadFile("data.json")
	err = json.Unmarshal(fileContent, &wotah)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save data"})
	}

	util.GetStatus(wotah)

	return c.JSON(http.StatusOK, nil)
}
