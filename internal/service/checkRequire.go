package service

import (
	"context"
	"errors"
	model "main/internal/models"
	errorx "main/internal/utils/myerror"
	"net/http"
	"strings"
	"time"
)

func (c *contractService) checkStudentCode(studentCode *string) error {
	if studentCode == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}
	if len(*studentCode) != 10 {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}

	return nil
}

func (c *contractService) checkFirstName(firstName *string) error {
	if firstName == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}
	arrayString := []byte(*firstName)
	if arrayString[0] == 32 {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input first name", errors.New("missing required field"), time.Now())
	}
	return nil
}

func (c *contractService) checkLastName(lastName *string) error {
	if lastName == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}
	arrayString := []byte(*lastName)
	if arrayString[0] == 32 {
	}
	for _, v := range arrayString {
		if v < 65 || v > 90 && v < 97 {
			return errorx.NewMyError(http.StatusUnprocessableEntity, "Special character", errors.New("wrong data type"), time.Now())
		}
	}
	return nil
}

func (c *contractService) checkEmail(email *string) error {
	if email == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}
	ok := strings.Contains(*email, "@gmail.com")
	if !ok {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input email", errors.New("missing required field"), time.Now())

	}

	return nil
}

func (c *contractService) checkPhone(phone *string) error {
	if phone == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}
	if len(*phone) != 10 {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input phone", errors.New("missing required field"), time.Now())
	}

	return nil
}

func (c *contractService) checkRoom(ctx context.Context, roomID *string) error {
	if roomID == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input student code", errors.New("missing required field"), time.Now())
	}
	if *roomID == "" {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Must input room", errors.New("missing required field"), time.Now())
	}

	totalContract, err := c.contractRepo.GetTotalContractRoom(ctx, *roomID)
	if err != nil {
		return err
	}

	if totalContract.Total > 4 {
		return errorx.NewMyError(http.StatusBadRequest, "Room full", errors.New("room full"), time.Now())

	}
	return nil
}

func (c *contractService) CheckRequiredField(ctx context.Context, contractModel *model.Contract) error {
	if err := c.checkStudentCode(contractModel.StudentCode); err != nil {
		return err
	}
	if err := c.checkFirstName(contractModel.FirstName); err != nil {
		return err
	}
	if err := c.checkLastName(contractModel.LastName); err != nil {
		return err
	}
	if err := c.checkEmail(contractModel.Email); err != nil {
		return err
	}
	if err := c.checkPhone(contractModel.Phone); err != nil {
		return err
	}
	if err := c.checkRoom(ctx, contractModel.RoomID); err != nil {
		return err
	}

	return nil
}
