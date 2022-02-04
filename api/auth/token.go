package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CriarToken(userId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("akjsdbkajhbdakjhdsbkjabhds66712653718270398ewhfudaosbhfbv"))
}

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r.Header)
	token, err := jwt.Parse(tokenString, retornaChaveVerificacao)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")

}

func extrairToken(header http.Header) string {
	token := header.Get("authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornaChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("assinatura token invalida")
	}
	return []byte("akjsdbkajhbdakjhdsbkjabhds66712653718270398ewhfudaosbhfbv"), nil
}
