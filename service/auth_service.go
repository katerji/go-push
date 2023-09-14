package service

import (
	"errors"
	"github.com/katerji/gopush/db"
	"github.com/katerji/gopush/db/query"
	"github.com/katerji/gopush/db/queryrow"
	"github.com/katerji/gopush/input"
	gopush "github.com/katerji/gopush/proto"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (service AuthService) Register(input input.AuthInput) (int, error) {
	password, err := hashPassword(input.Password)
	if err != nil {
		return 0, err
	}
	input.Password = password
	return db.GetDbInstance().Insert(query.InsertUserQuery, input.Email, input.Password)
}

func (service AuthService) Login(input input.AuthInput) (*gopush.User, error) {
	result := queryrow.UserQueryRow{}
	client := db.GetDbInstance()
	row := client.QueryRow(query.GetUserByEmailQuery, input.Email)
	err := row.Scan(&result.ID, &result.Email, &result.Password)
	if err != nil {
		return nil, errors.New("email does not exist")
	}

	if !validPassword(result.Password, input.Password) {
		return nil, errors.New("incorrect password")
	}

	return &gopush.User{
		Id:    int64(result.ID),
		Email: result.Email,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
