package curl

import (
	"io/ioutil"
	"net/http"
)

func GetBytesFromUrl(url string) (body []byte, err error) {
	res := []byte{}
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
