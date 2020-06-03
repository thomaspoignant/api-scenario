package util

import "github.com/sirupsen/logrus"

// ExitIfErr exit if there is an error
func ExitIfErr(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
