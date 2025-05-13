package service

import (
	"errors"
	model "main/internal/models"
	errorx "main/internal/utils/myerror"
	"net/http"
	"strings"
	"time"
)

func checkStudentCode(studentCode string) error {
	if len(studentCode) != 10 {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}

	return nil
}

func checkFirstName(firstName string) error {
	arrayString := []byte(firstName)
	if arrayString[0] == 32 {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input first name", errors.New("missing required field"), time.Now())
	}
	return nil
}

func checkLastName(lastName string) error {
	arrayString := []byte(lastName)
	if arrayString[0] == 32 {
	}
	for _, v := range arrayString {
		if v < 65 || v > 90 && v < 97 {
			return errorx.NewMyError(http.StatusUnprocessableEntity, "Special character", errors.New("wrong data type"), time.Now())
		}
	}
	return nil
}

func checkEmail(email string) error {
	ok := strings.Contains(email, "@gmail.com")
	if !ok {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input email", errors.New("missing required field"), time.Now())

	}

	return nil
}

func checkPhone(phone string) error {
	if len(phone) != 10 {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input phone", errors.New("missing required field"), time.Now())
	}

	return nil
}

func CheckRequiredField(contractModel *model.Contract) error {
	if err := checkStudentCode(*contractModel.StudentCode); err != nil {
		return err
	}
	if err := checkFirstName(*contractModel.FirstName); err != nil {
		return err
	}
	if err := checkLastName(*contractModel.LastName); err != nil {
		return err
	}
	if err := checkEmail(*contractModel.Email); err != nil {
		return err
	}
	if err := checkPhone(*contractModel.Phone); err != nil {
		return err
	}

	return nil
}
