package builder

import (
	"fmt"
)

func MakeBuilder(t string) (Builder, error) {
	if t == "normal" {
		return &NormalBuilder{}, nil
	}

	if t == "igloo" {
		return &IglooBuilder{}, nil
	}

	return nil, fmt.Errorf("undefined builder type '%s'", t)
}
