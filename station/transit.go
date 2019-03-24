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
func (scooter *Scooter) Transit(to State, role user.Role) error {
	if err := validateTransit(scooter.State, to, role); err != nil {
		return err
	}

	scooter.State = to
	scooter.commitTransit()
	return nil
}

func (scooter Scooter)commitTransit() {
	SetScooterState(scooter.ID, scooter.State)
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
