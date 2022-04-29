package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	return errors.New("Incorrect Details")
}
