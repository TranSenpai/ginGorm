package repo

import (
	errox "main/internal/utils/myerror"
	"net/http"
	"time"

	"gorm.io/gorm"
)

var mapError = map[string]int{
	gorm.ErrRecordNotFound.Error():                http.StatusNotFound,
	gorm.ErrInvalidTransaction.Error():            http.StatusInternalServerError,
	gorm.ErrNotImplemented.Error():                http.StatusInternalServerError,
	gorm.ErrMissingWhereClause.Error():            http.StatusInternalServerError,
	gorm.ErrUnsupportedRelation.Error():           http.StatusInternalServerError,
	gorm.ErrPrimaryKeyRequired.Error():            http.StatusBadRequest,
	gorm.ErrModelValueRequired.Error():            http.StatusBadRequest,
	gorm.ErrModelAccessibleFieldsRequired.Error(): http.StatusBadRequest,
	gorm.ErrSubQueryRequired.Error():              http.StatusInternalServerError,
	gorm.ErrInvalidData.Error():                   http.StatusBadRequest,
	gorm.ErrUnsupportedDriver.Error():             http.StatusInternalServerError,
	gorm.ErrRegistered.Error():                    http.StatusInternalServerError,
	gorm.ErrInvalidField.Error():                  http.StatusBadRequest,
	gorm.ErrDryRunModeUnsupported.Error():         http.StatusInternalServerError,
	gorm.ErrInvalidDB.Error():                     http.StatusInternalServerError,
	gorm.ErrInvalidValue.Error():                  http.StatusBadRequest,
	gorm.ErrPreloadNotAllowed.Error():             http.StatusInternalServerError,
	gorm.ErrDuplicatedKey.Error():                 http.StatusBadRequest,
	gorm.ErrForeignKeyViolated.Error():            http.StatusInternalServerError,
}

func GetError(err error) error {
	value, exist := mapError[err.Error()]
	if !exist {
		value = http.StatusInternalServerError
	}
	return errox.NewMyError(value, err.Error(), err, time.Now())
}
