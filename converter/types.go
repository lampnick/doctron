package converter

type ConvertConfig struct {
	//the source document's url to be converted.
	Url    string
	Params Params
}

// convert params
type Params interface {
}

type ConvertOutput struct {
}
