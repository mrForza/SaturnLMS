package validators

import (
	"fmt"
)

func ValidateStringField(field string, maxLen int, nameOfField string) error {
	if len(field) > maxLen {
		return fmt.Errorf("your '%s' should not be greater than %d", nameOfField, maxLen)
	}
	return nil
}
