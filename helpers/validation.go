package helpers

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func ValidateCannotBeEmpty(key, value string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s cannot be empty", key)
	}

	return nil
}

func ValidateCannotBeEmptyAndExceedCharacter(key, value string, number int) error {
	if err := ValidateCannotBeEmpty(key, value); err != nil {
		return err
	}

	if len(value) > number {
		return fmt.Errorf("%s cannot exceed %d characters", key, number)
	}

	return nil
}

func ValidateDuplicatedKey(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil {
			return gorm.ErrDuplicatedKey
		}

		return err
	}

	return nil
}
