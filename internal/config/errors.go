package config

import "fmt"

type ValidationError struct {
	field string
	issue string
}

func NewValidationError(field, issue string) *ValidationError {
	return &ValidationError{
		field: field,
		issue: issue,
	}
}

func (s *ValidationError) Error() string {
	return fmt.Sprintf("%[1]s: %[2]s", s.field, s.issue)
}
