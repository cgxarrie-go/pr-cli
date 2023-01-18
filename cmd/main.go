package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	pat := "`:azurepat"
	companyName := "companyName"
	projectID := "projectID"
	repoID := "repoID"
	status := "1"
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/%s/pullrequests?searchCriteria.status=%s&$top=1001&api-version=5.1", companyName, projectID, repoID, status)

	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	fmt.Println(string(body))

}
