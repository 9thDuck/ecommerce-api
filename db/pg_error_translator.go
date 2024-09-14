package db

import (
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	UniqueViolationTemplate = "The given %s is already taken"
	ForeignKeyViolationTemplate = "Invalid %s: does not exist"
	NotNullViolationTemplate = "%s cannot be null"
	CheckViolationTemplate = "Invalid %s: fails check constraint"
	StringDataRightTruncationTemplate = "%s is too long"
	NumericValueOutOfRangeTemplate = "%s is out of range"
	DefaultErrorTemplate = "An error occurred with %s"
)

func TranslatePgErrors(pgError *pgconn.PgError) error {
	var fieldName string
	switch pgError.ConstraintName {
	// Users table errors
	case "uni_users_email":
		fieldName = "email"
	case "uni_users_username":
		fieldName = "username"
	case "fk_products_category":
		fieldName = "category"
	}

	var template string

	switch pgError.Code {
	case pgerrcode.UniqueViolation:
		template = UniqueViolationTemplate
	case pgerrcode.ForeignKeyViolation:
		template = "Invalid %s: does not exist"
	case pgerrcode.NotNullViolation:
		template = "%s cannot be null"
	case pgerrcode.CheckViolation:
		template = CheckViolationTemplate
	case pgerrcode.StringDataRightTruncationDataException:
		template = StringDataRightTruncationTemplate
	case pgerrcode.NumericValueOutOfRange:
		template = NumericValueOutOfRangeTemplate
	default:
		template = "An error occurred with %s"
	}

	return fmt.Errorf(template, fieldName)
}
