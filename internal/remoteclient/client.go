package remoteclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type client struct {
}

func newClient() client {
	return client{}
}

func (c *client) doCreate(req *http.Request) (
	resp any, err error) {

	client := &http.Client{}
	azResp, err := client.Do(req)

	if err != nil {
		return resp, errors.Wrap(err, "doing request")
	}

	if azResp.StatusCode != http.StatusCreated {
		respBody, err := io.ReadAll(azResp.Body)
		if err != nil {
			respBody = []byte("cannot read response body content")
		}

		return resp, fmt.Errorf("response code: %d - "+
			"response body: %+v", azResp.StatusCode, string(respBody))

	}

	defer azResp.Body.Close()
	err = json.NewDecoder(azResp.Body).Decode(&resp)
	if err != nil {
		return resp, errors.Wrap(err, "decoding response body")
	}

	return
}
