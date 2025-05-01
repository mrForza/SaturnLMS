package validators

import (
	"errors"
)

func ValidateCourseNumber(groupNumber uint8) error {
	if groupNumber < 1 {
		return errors.New("your 'course number' should not be less than 1")
	}

	if groupNumber > 6 {
		return errors.New("your 'course number' should not be greater than 6")
	}

	return nil
}

func ValidateGroupNumber(CourseNumber uint8) error {
	return nil
}

func ValidateUniversityName(aboutMe string) error {
	if len(aboutMe) > 64 {
		return errors.New("the maximum size of 'university name' section is 65536")
	}

	return nil
}

func ValidateFacultatyName(interests string) error {
	if len(interests) > 64 {
		return errors.New("the maximum size of 'facultatu name' section is 65536")
	}
	return nil
}

func ValidateProgramName(interests string) error {
	if len(interests) > 64 {
		return errors.New("the maximum size of 'program name' section is 65536")
	}
	return nil
}
