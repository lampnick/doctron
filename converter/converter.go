package converter

import (
	"time"
)

type Converter interface {
	Convert() ([]byte, error)
	GetConvertElapsed() time.Duration
}
