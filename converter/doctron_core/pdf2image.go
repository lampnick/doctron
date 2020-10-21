package doctron_core

import (
	"time"
)

type pdf2Image struct {
	Doctron
}

func (p pdf2Image) Convert() ([]byte, error) {
	panic("implement me")
}

func (p pdf2Image) GetConvertElapsed() time.Duration {
	panic("implement me")
}
