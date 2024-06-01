package controller

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"planet.com/dto"
	"planet.com/model"
	"planet.com/pkg/helper"
	"planet.com/pkg/lib"
)

type Exoplanet interface {
	AddExoplanet()
	ListAllExoplanet()
	ListExoplanetById()
	UpdateExoplanet()
	DeleteExoplanetById()
	FuelEstimation()
}
type ExoplanetController struct{}

func (con ExoplanetController) AddExoplanet(c *gin.Context) {

	var InputDTO dto.AddExpplanet

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)

		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	helper.Trimmer(&InputDTO)
	var Expplanet model.Expplanet
	Expplanet.Name = InputDTO.Name
	Expplanet.Description = InputDTO.Description
	Expplanet.DistanceFromEarth = InputDTO.DistanceFromEarth
	Expplanet.Radius = InputDTO.Radius
	Expplanet.Mass = InputDTO.Mass
	Expplanet.Type = InputDTO.Type
	if result := model.DB.Create(&Expplanet); result.Error != nil {
		response := lib.Error("sql error", result.Error.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := lib.Success(true, "ok", "Exoplanet created successfully")
	c.JSON(http.StatusOK, response)
}

func (con ExoplanetController) ListAllExoplanet(c *gin.Context) {
	var InputDTO dto.SortAndFilter
	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)

		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	result, err := model.FinAll(InputDTO)
	if err != nil {
		response := lib.Error("sql", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := lib.Success(true, "ok", result)
	c.JSON(http.StatusOK, response)
}
func (con ExoplanetController) ListExoplanetById(c *gin.Context) {
	var InputDTO dto.GetId
	if errDTO := c.ShouldBindUri(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	result, err := model.FindById(InputDTO.Id)
	if err != nil {
		response := lib.Error("sql", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := lib.Success(true, "ok", result)
	c.JSON(http.StatusOK, response)
}

func (con ExoplanetController) UpdateExoplanet(c *gin.Context) {

	var InputDTO dto.UpdateExpplanet

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)

		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	helper.Trimmer(&InputDTO)
	var Expplanet model.Expplanet
	Expplanet.Name = InputDTO.Name
	Expplanet.Description = InputDTO.Description
	Expplanet.DistanceFromEarth = InputDTO.DistanceFromEarth
	Expplanet.Radius = InputDTO.Radius
	Expplanet.Mass = InputDTO.Mass
	Expplanet.Type = InputDTO.Type

	_, err := model.UpdateById(Expplanet, InputDTO.Id)
	if err != nil {
		response := lib.Error("sql", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := lib.Success(true, "ok", "Record updated successfully")
	c.JSON(http.StatusOK, response)
}
func (con ExoplanetController) DeleteExoplanetById(c *gin.Context) {
	var InputDTO dto.GetId
	if errDTO := c.ShouldBindUri(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	_, err := model.DeteteByID(InputDTO.Id)
	if err != nil {
		response := lib.Error("sql", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := lib.Success(true, "ok", "record deleted successfully")
	c.JSON(http.StatusOK, response)
}

func (con ExoplanetController) FuelEstimation(c *gin.Context) {

	var InputDTO dto.FuleCalculation

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)

		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	helper.Trimmer(&InputDTO)
	result, err := model.FindById(InputDTO.ExoPlanetId)
	if err != nil {
		response := lib.Error("sql", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if result == nil {
		response := lib.Error("Warning", "Inavlid exoplanel", lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var g float64
	if result.Type == "GasGiant" {
		//g = (0.5/r^2)
		g = 0.5 / math.Pow(result.Radius, 2)
	} else if result.Type == "Terrestrial" {
		//g = (m/r^2)
		g = result.Mass / math.Pow(result.Radius, 2)
	}
	// f = d / (g^2) * c units
	f := float64(result.DistanceFromEarth) / math.Pow(g, 2) * float64(InputDTO.CrewCapacity)
	units := fmt.Sprintf("%f units", f)
	response := lib.Success(true, "ok", units)
	c.JSON(http.StatusOK, response)
}
