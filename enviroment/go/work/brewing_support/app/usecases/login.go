package usecases

import (
	"errors"
	"time"
	"brewing_support/app/domain"
	"brewing_support/app/utils/config"

	"github.com/dgrijalva/jwt-go"
)

// persistance API用のインターフェース
type loginRepo interface {
	SelectByAccountID(string) (*domain.Staff, error)
}

// Login 成功の場合accountIDに対応するstaffの構造体の参照と認証tokenを返す。
func Login(staffRepo loginRepo, id string, pass string) (*domain.Staff, string, error) {
	s, err := staffRepo.SelectByAccountID(id)

	if err != nil {
		return nil, "", err
	}

	if s.Password == pass {
		token, generateError := generateToken(s.ID)
		if generateError != nil {
			return nil, "", generateError
		}

		s.Password = ""

		return s, token, nil
	}
	return nil, "", errors.New("login incorrect")
}

func generateToken(id string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["staffID"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.LoginExpTime())).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SecretKey()))
	if err != nil {
		return "", err
	}
	return t, nil
}
