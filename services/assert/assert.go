package assert

import (
	"github.com/sirupsen/logrus"
)

// Nil panic if the test is not nil
func Nil(test interface{}) {
	if test != nil {
		if e, ok := test.(error); ok {
			logrus.Panicln(e.Error())
			return
		}
	}
}

// True check if the value is true and panic if its not
func True(test bool) {
	if !test {
		logrus.Panicln("must be true but is not")
	}
}

// returns true if all are not empty
func String(runes ...string) bool {
	for i := range runes {
		if runes[i] == "" {
			return false
		}
	}

	return true
}