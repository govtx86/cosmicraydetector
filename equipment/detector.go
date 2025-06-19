package equipment

import (
	"errors"
	"math/rand"
)

type Detector struct {
	seedText string
}

func NewDetector() Detector {
	return Detector{
		seedText: "Totally Uncorruptable text!!",
	}
}

func (detector *Detector) Detect() (bool, error) {
	var err error

	if rand.Intn(100) == 69 {
		err = errors.New("unknown error encounterd")
	} else {
		err = nil
	}

	if detector.seedText == "Totally Uncorruptable text!!" {
		return false, err
	} else {
		return true, err
	}
}
