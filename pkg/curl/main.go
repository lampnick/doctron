package curl

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func GetBytesFromUrl(u string) (body []byte, err error) {
	var res []byte
	parse, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	var file io.ReadCloser
	if parse.Scheme == "file" {
		file, err = os.Open(parse.Host + parse.Path)
		if err != nil {
			return nil, err
		}
	} else {
		resp, err := http.Get(u)
		if err != nil {
			return res, err
		}
		file = resp.Body
	}

	defer file.Close()

	return io.ReadAll(file)
}
