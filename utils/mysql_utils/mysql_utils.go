package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/shakilbd009/go-users-api/utils/errors"
)

const (
	email_UNIQUE = "email_UNIQUE"
	errorNoRows  = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record found with given id")
		}
		fmt.Println(err.Error())
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	//fmt.Println(err.Error())
	return errors.NewInternalServerError("error processing request")
}
