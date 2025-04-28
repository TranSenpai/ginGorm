package service

import (
	"context"
	"encoding/base64"
	"main/internal/entity"
	model "main/internal/models"
	repo "main/internal/repo"
	errorx "main/internal/utils/myerror"
)

type ContractBuilder struct {
	model      *model.Contract
	entity     *entity.Contract
	buildError error
	ctx        context.Context
}

func NewContractBuilder(ctx context.Context, m *model.Contract) *ContractBuilder {
	return &ContractBuilder{
		model:      m,
		entity:     &entity.Contract{},
		buildError: nil,
		ctx:        ctx,
	}
}

func (b *ContractBuilder) MapStudentInfo() *ContractBuilder {
	if b.buildError != nil {
		return b
	}

	b.entity.StudentCode = b.model.StudentCode
	b.entity.FullName = b.model.FullName
	b.entity.Gender = b.model.Gender
	b.entity.DOB = &b.model.DOB

	return b
}

func (b *ContractBuilder) MapContactInfo() *ContractBuilder {
	if b.buildError != nil {
		return b
	}

	b.entity.Email = b.model.Email
	b.entity.Phone = b.model.Phone
	b.entity.Address = &b.model.Address

	return b
}

func (b *ContractBuilder) MapRoomInfo() *ContractBuilder {
	if b.buildError != nil {
		return b
	}

	b.entity.RoomID = b.model.RoomID
	b.entity.NotificationChannels = b.model.NotificationChannels

	return b
}

func (b *ContractBuilder) MapAvatar() *ContractBuilder {
	if b.buildError != nil {
		return b
	}

	avatarData, err := base64.StdEncoding.DecodeString(b.model.Avatar)
	if err != nil {
		b.buildError = errorx.New(errorx.StatusUnprocessableEntity, "Invalid Avatar format", err)
		return b
	}
	b.entity.Avatar = avatarData

	return b
}

func (b *ContractBuilder) ValidateUniqueness(repo repo.IContractRepo) *ContractBuilder {
	if b.buildError != nil {
		return b
	}

	_, err := repo.Search(b.ctx, b.entity.StudentCode)
	if err == nil {
		b.buildError = errorx.New(errorx.StatusConflict, "Contract already exists", nil)
		return b
	}

	return b
}

func (b *ContractBuilder) GetContract() (*entity.Contract, error) {
	if b.buildError != nil {
		return nil, b.buildError
	}
	return b.entity, nil
}
