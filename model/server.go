package model

import (
	"fmt"
)

type Server interface {
	Start() error
	Stop() error
	fmt.Stringer
}
