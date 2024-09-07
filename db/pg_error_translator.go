package db

import (
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	UniqueViolationTemplate = "given %s is already taken"
)

func TranslatePgErrors(pgError *pgconn.PgError) error {
	var fieldName string
	switch pgError.ConstraintName {
	// Users table errors
	case "uni_users_email":
		fieldName = "email"
	case "uni_users_username":
		fieldName = "username"
	}

	var template string

	switch pgError.Code {
	case pgerrcode.UniqueViolation:
		template = UniqueViolationTemplate
	}

	return fmt.Errorf(template, fieldName)
}
