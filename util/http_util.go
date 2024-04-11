package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func HttpResponseHandler(resp *http.Response) ([]byte, error) {
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return bodyByte, nil
	case http.StatusUnauthorized:
		return nil, errors.New(string(bodyByte))
	default:
		return nil, fmt.Errorf("%+v", string(bodyByte))
	}
}
