package station

import (
	"testing"
	"muskooters/user"
)

func TestValidateTransit(t *testing.T) {
	for _, test := range testCases {
		err := validateTransit(test.from, test.to, test.role)
		if err == nil {
			t.Error("all of cases should be rejected")
		}
	}
}

type transtionCase struct {
	from State
	to State
	role user.Role
}

var testCases = []transtionCase{
	{Bounty, Collected, user.Client},
	{Riding, BatteryLow, user.Client},
	{BatteryLow, Bounty, user.Client},
	{Collected, Dropped, user.Client},
	{Ready, Unknown, user.Hunter},
	{Riding, BatteryLow, user.Hunter},
	{BatteryLow, Ready, user.Hunter},
	{Collected, Ready, user.Hunter},
}