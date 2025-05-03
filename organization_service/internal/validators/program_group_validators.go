package validators

import "errors"

func ValidateCourseNumber(courseNumber uint8) error {
	if courseNumber > 8 {
		return errors.New("course number can not be higher than 8")
	}
	return nil
}
