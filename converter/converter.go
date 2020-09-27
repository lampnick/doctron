package converter

import "context"

type Converter interface {
	Convert(ctx context.Context, cc ConvertConfig) ([]byte, error)
	Upload(ctx context.Context, co ConvertOutput) (string, error)
}
