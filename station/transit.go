package station

import (
	"errors"
	"muskooters/user"
)

var (
	ErrNotAuthorized     = errors.New("not authorized with such role")
	ErrInvalidTransition = errors.New("invalid transition destination")
)

// validates a scooter transit action
func (scooter Scooter) Transit(to State, role user.Role) error {
	return validateTransit(scooter.State, to, role)
}

func validateTransit(from, to State, transporterRole user.Role) error {
	// zeus can change every state
	if transporterRole == user.Zeus {
		return nil
	}

	if _, ok := hashMap[from]; !ok {
		return ErrInvalidTransition
	}

	validRoles, ok := hashMap[from][to]
	if !ok {
		return ErrInvalidTransition
	}

	if !containsRole(validRoles, transporterRole) {
		return ErrNotAuthorized
	}

	return nil
}

func containsRole(roles []user.Role, role user.Role) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}

	return false
}
