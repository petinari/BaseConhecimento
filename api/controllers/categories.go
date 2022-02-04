package controllers

import (
	"api/db"
	"api/models"
	"api/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveCategories(c *gin.Context) {
	corpoRequest, erro := ioutil.ReadAll(c.Request.Body)
	if erro != nil {
		c.JSON(http.StatusUnprocessableEntity, erro)
		return
	}
	var categories models.Categories
	if erro := json.Unmarshal(corpoRequest, &categories); erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewCategoriesRepository(db)
	categories.ID, erro = repository.SaveCategories(categories)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	c.JSON(http.StatusCreated, categories.ID)
}

func GetAllCategories(c *gin.Context) {
	db, err := db.Conn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	repository := repository.NewCategoriesRepository(db)
	categories, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, categories)

}
