package station

import (
	"errors"
	"muskooters/user"
)

var (
	// role does not have permission for requested transition
	ErrNotAuthorized = errors.New("not authorized with such role")
	// transition is not valid in defined machine
	ErrInvalidTransition = errors.New("invalid transition destination")
)

// validates a scooter transit action
func (scooter *Scooter) Transit(to State, role user.Role) error {
	if err := validateTransit(scooter.State, to, role); err != nil {
		return errors.New("could not validate transit: " + err.Error())
	}

	scooter.State = to
	scooter.commitTransit()
	return nil
}

// updates state status in database
func (scooter Scooter) commitTransit() {
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
