package doctron_core

import (
	"context"
	"github.com/lampnick/doctron/converter"
	"time"
)

type Doctron struct {
	ctx            context.Context
	cc             converter.ConvertConfig
	buf            []byte
	convertElapsed time.Duration
}

type DoctronI interface {
	converter.Converter
}
