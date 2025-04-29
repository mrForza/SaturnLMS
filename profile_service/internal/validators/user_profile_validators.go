package validators

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidateUserInitials(initial string, nameOfInitial string) error {
	if len(initial) < 2 {
		return fmt.Errorf("your %s should not be less than 2 symbols", nameOfInitial)
	}

	if len(initial) > 64 {
		return fmt.Errorf("your %s should not be greater than 64 symbols", nameOfInitial)
	}

	pattern := `^[a-zA-Zа-яА-Я\-]`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(initial) {
		return fmt.Errorf("your %s should consists of only latin or cirillic letters", nameOfInitial)
	}

	return nil
}

func ValidateAge(age uint8) error {
	if age < 16 {
		return errors.New("your age should not be less than 16")
	}

	if age > 100 {
		return errors.New("your age should not be greater than 100")
	}

	return nil
}

func ValidateGender(gender bool) error {
	// if gender != false && gender != true {
	// 	return errors.New("you are available to choose only two genders: Male or Female")
	// }
	return nil
}

func ValidateAboutMe(aboutMe string) error {
	if len(aboutMe) > 65536 {
		return errors.New("the maximum size of AboutMe section is 65536")
	}

	return nil
}

func ValidateInterests(interests string) error {
	if len(interests) > 65536 {
		return errors.New("the maximum size of Interests section is 65536")
	}
	return nil
}
