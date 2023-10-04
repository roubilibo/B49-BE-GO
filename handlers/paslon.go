package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Paslons struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Vision string `json:"vision"`
	Image  string `json:"image"`
}

var paslons = []Paslons{
	{
		Id:     "1",
		Name:   "Cupang Baronang",
		Vision: "saya berjanji",
		Image:  "image.jpg",
	},
	{
		Id:     "2",
		Name:   "Cupang Bali",
		Vision: "saya tidak berjanji",
		Image:  "image.jpg",
	},
	{
		Id:     "3",
		Name:   "Cupang Lombok",
		Vision: "saya akan melakukan",
		Image:  "image.jpg",
	},
}

func FindPaslons(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(paslons)
}

func GetPaslon(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	id := c.Param("id")

	var paslonData Paslons
	var isGetPaslon = false

	for _, paslon := range paslons {
		if id == paslon.Id {
			isGetPaslon = true
			paslonData = paslon
		}
	}

	if !isGetPaslon {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(paslonData)
}

func CreatePaslon(c echo.Context) error {
	var data Paslons

	json.NewDecoder(c.Request().Body).Decode(&data)

	paslons = append(paslons, data)

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(paslons)
}

func UpdatePaslon(c echo.Context) error {
	id := c.Param("id")
	var data Paslons
	var isGetPaslon = false

	json.NewDecoder(c.Request().Body).Decode(&data)

	for idx, paslon := range paslons {
		if id == paslon.Id {
			isGetPaslon = true
			paslons[idx] = data
		}
	}

	if !isGetPaslon {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(paslons)
}

func DeletePaslon(c echo.Context) error {
	id := c.Param("id")
	var isGetPaslon = false
	var index = 0

	for idx, paslon := range paslons {
		if id == paslon.Id {
			isGetPaslon = true
			index = idx
		}
	}

	if !isGetPaslon {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	paslons = append(paslons[:index], paslons[index+1:]...)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(paslons)
}
