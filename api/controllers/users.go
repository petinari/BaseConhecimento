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

// 	c.JSON(http.StatusOK, "users1")
// }

func CadastrarUsuarios(c *gin.Context) {
	corpoRequest, erro := ioutil.ReadAll(c.Request.Body)
	if erro != nil {
		c.JSON(http.StatusUnprocessableEntity, erro)
		return
	}
	var user models.User
	if erro := json.Unmarshal(corpoRequest, &user); erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	if erro := user.Preparar("cadastro"); erro != nil {
		c.JSON(http.StatusUnprocessableEntity, erro)
		return
	}
	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewUserRepository(db)
	user.ID, erro = repository.SaveUser(user)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	c.JSON(http.StatusCreated, user.ID)
}

func GetAllUsers(c *gin.Context) {
	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewUserRepository(db)
	users, err := repository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusFound, users)
}
func GetUserById(c *gin.Context) {
	params, _ := c.Params.Get("id")
	id, _ := strconv.Atoi(params)
	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NewUserRepository(db)
	users, err := repository.GetUserById(uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if users == nil {
		c.JSON(http.StatusNoContent, "user not found")
		return
	}
	c.JSON(http.StatusOK, users)
}
