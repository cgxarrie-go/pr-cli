package remoteclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/pkg/errors"
)

type client struct {
	remote ports.Remote
}

func newClient(r ports.Remote) client {
	return client{
		remote: r,
	}
}

func (c *client) doCreate(req *http.Request) (
	resp any, err error) {

	client := &http.Client{}
	clResp, err := client.Do(req)

	if err != nil {
		return resp, errors.Wrap(err, "doing request")
	}

	if clResp.StatusCode != http.StatusCreated {
		respBody, err := io.ReadAll(clResp.Body)
		if err != nil {
			respBody = []byte("cannot read response body content")
		}

		return resp, fmt.Errorf("response code: %d - "+
			"response body: %+v", clResp.StatusCode, string(respBody))

	}

	defer clResp.Body.Close()
	err = json.NewDecoder(clResp.Body).Decode(&resp)
	if err != nil {
		return resp, errors.Wrap(err, "decoding response body")
	}

	return
}

func (c *client) doGet(req *http.Request) (resp any, err error) {

	client := &http.Client{}
	clResp, err := client.Do(req)
	if err != nil {
		return resp, errors.Wrap(err, "doing request")
	}

	if clResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("%d - %s", clResp.StatusCode, clResp.Status)
	}

	defer clResp.Body.Close()
	err = json.NewDecoder(clResp.Body).Decode(resp)
	if err != nil {
		return resp, errors.Wrap(err, "decoding response body")
	}

	return
}
