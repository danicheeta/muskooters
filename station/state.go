package station

import (
	"muskooters/user"
	"muskooters/services/assert"
)

const (
	Ready State = iota
	BatteryLow
	Bounty
	Riding
	Collected
	Dropped
	Unknown
)

var stringToState = map[string]State{
	"Ready":      Ready,
	"BatteryLow": BatteryLow,
	"Bounty":     Bounty,
	"Riding":     Riding,
	"Collected":  Collected,
	"Dropped":    Dropped,
	"Unknown":    Unknown,
}

type State int

type Transition struct {
	From  State
	To    State
	Roles []user.Role
}

var (
	csRoles = []user.Role{user.Client, user.Scooter}
	sRoles  = []user.Role{user.Scooter}
	hRoles  = []user.Role{user.Hunter}

	hashMap = map[State]map[State][]user.Role{}
)

// Default transitions
var graph = []Transition{
	{From: Ready, To: Riding, Roles: csRoles},
	{From: Ready, To: Bounty, Roles: sRoles},
	{From: Riding, To: Ready, Roles: csRoles},
	{From: Riding, To: BatteryLow, Roles: sRoles},
	{From: BatteryLow, To: Bounty, Roles: sRoles},
	{From: Bounty, To: Collected, Roles: hRoles},
	{From: Collected, To: Dropped, Roles: hRoles},
	{From: Dropped, To: Ready, Roles: hRoles},
	{From: Ready, To: Unknown, Roles: sRoles},
}

// wrap to query easier
func init() {
	for _, t := range graph {
		from, to := t.From, t.To

		if f, ok := hashMap[from]; ok {
			// assert cause that's invalid type of graph
			_, ok := hashMap[from][to]
			assert.True(!ok)

			f[to] = t.Roles
		} else {
			hashMap[from] = map[State][]user.Role{t.To: t.Roles}
		}
	}
}
