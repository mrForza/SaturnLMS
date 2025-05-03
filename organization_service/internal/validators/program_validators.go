package validators

import "errors"

func ValidateType(programType string) error {
	if programType != "Bachelor" && programType != "Master" && programType != "Post-Degree" {
		return errors.New("inocrrect type of program")
	}

	return nil
}
