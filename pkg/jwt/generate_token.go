package jwt

import (
	"fmt"
	"time"
	"zxsttm/database/models"
	"zxsttm/server/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	signedToken, err := token.SignedString(config.Config.PrivateKey)

	if err != nil {
		fmt.Println("Error while signing the token: ", err)
		return "", err
	}

	return signedToken, nil
}
