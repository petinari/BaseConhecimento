package controllers

import (
	"api/auth"
	"api/db"
	"api/models"
	"api/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	var userRequest models.User
	erro := json.Unmarshal(requestBody, &userRequest)
	if erro != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	db, erro := db.Conn()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetUserByEmail(userRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	erro = auth.VerificarSenha(user.Password, userRequest.Password)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, erro.Error())
		return
	}
	token, err := auth.CriarToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)

}
