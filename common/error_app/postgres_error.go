package error_app

import (
	"errors"

	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	// PostgreSQL error codes
	errCodeUniqueViolation     = "23505"
	errCodeForeignKeyViolation = "23503"
	errCodeNotNullViolation    = "23502"
	// Add more codes as needed
)

// IsUniqueViolation returns true if the error is a PostgreSQL unique constraint violation.
func IsUniqueViolation(err error) bool {
	return isPgErrorCode(err, errCodeUniqueViolation)
}

// IsForeignKeyViolation returns true if the error is a foreign key violation.
func IsForeignKeyViolation(err error) bool {
	return isPgErrorCode(err, errCodeForeignKeyViolation)
}

// IsNotNullViolation returns true if the error is a not-null constraint violation.
func IsNotNullViolation(err error) bool {
	return isPgErrorCode(err, errCodeNotNullViolation)
}

// Helper: checks if error is pgconn.PgError and matches given code
func isPgErrorCode(err error, code string) bool {
	var pgErr pgdriver.Error
	if errors.As(err, &pgErr) {
		return pgErr.Field('C') == code
	}
	return false
}


