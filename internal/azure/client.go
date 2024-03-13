package azure

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cgxarrie-go/prq/internal/ports"
)

type client struct {
	pat string
}

func NewClient(pat string) ports.RemoteClient {
	return &client{
		pat: pat,
	}
}

func (c *client) Create(req ports.RemoteClientCreateRequest) error {

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(c.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	pullRequest := map[string]interface{}{
		"sourceRefName": req.Source,      // Source branch
		"targetRefName": req.Description, // Target branch
		"title":         req.Title,       // Title of PR
		"isDraft":       req.IsDraft,     // Draft PR
		"description":   req.Description, // Description of PR
	}

	body, err := json.Marshal(pullRequest)
	if err != nil {
		return fmt.Errorf("marshalling request body: %w", err)
	}

	azReq, err := http.NewRequest("POST", req.URL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("creating http request: %w", err)
	}

	azReq.Header.Add("Authorization", bearer)
	azReq.Header.Add("Content-Type", "application/json")
	azReq.Header.Add("Host", "dev.azure.com")
	azReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/117.0")
	azReq.Header.Add("Accept", "application/json;api-version=5.0-preview.1;excludeUrls=true;enumsAsNumbers=true;msDateFormat=true;noArrayWrap=true")
	azReq.Header.Add("Accept-Encoding", "gzip,deflate,br")
	azReq.Header.Add("Referer", fmt.Sprintf("https://dev.azure.com/%s/%s/_git/"+
		"%s/pullrequestcreate?sourceRef=%s&targetRef=%s"+
		"&sourceRepositoryId=%s&targetRepositoryId=%s", o.Organization(),
		o.Project(), o.Repository(), src.Name(), dest.Name(), o.Repository(),
		o.Repository()))
	azReq.Header.Add("Origin", "https://dev.azure.com")
	azReq.Header.Add("Connection", "keep-alive")
	azReq.Header.Add("Sec-Fetch-Dest", "empty")
	azReq.Header.Add("Sec-Fetch-Mode", "cors")
	azReq.Header.Add("Sec-Fetch-Site", "same-origin")

	return nil
}
