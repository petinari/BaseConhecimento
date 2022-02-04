package controllers

import (
	"api/db"
	"api/models"
	"api/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func BuscarUsuarios(c *gin.Context) {

// 	c.JSON(http.StatusOK, "articless1")
// }

func SaveArticles(c *gin.Context) {
	corpoRequest, erro := ioutil.ReadAll(c.Request.Body)
	if erro != nil {
		c.JSON(http.StatusUnprocessableEntity, erro)
		return
	}
	var articles models.Articles
	if erro := json.Unmarshal(corpoRequest, &articles); erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewArticlesRepository(db)
	articles.ID, erro = repository.SaveArticles(articles)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	c.JSON(http.StatusCreated, articles.ID)
}

func GetAllArticles(c *gin.Context) {
	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewArticlesRepository(db)
	articles, err := repository.GetAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusFound, articles)
}

func GetArticlesById(c *gin.Context) {
	params, _ := c.Params.Get("id")
	id, _ := strconv.Atoi(params)
	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewArticlesRepository(db)
	articles, err := repository.GetArticlesById(uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if articles.ID == 0 {
		c.JSON(http.StatusNoContent, "article not found")
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticlesByCategory(c *gin.Context) {
	params, _ := c.Params.Get("cat")

	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewArticlesRepository(db)
	articles, err := repository.GetAllArticlesByCategory(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}
