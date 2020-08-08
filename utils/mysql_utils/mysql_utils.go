package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

const (
	email_UNIQUE = "email_UNIQUE"
	ErrorNoRows  = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record found with given id")
		}
		fmt.Println(err.Error())
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}
	//fmt.Println(err.Error())
	return rest_errors.NewInternalServerError("error processing request", fmt.Errorf("database error"))
}
