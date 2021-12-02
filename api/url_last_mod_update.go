package main

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"strings"
	"time"
)

type (
	ErrorResponseData struct {
		Error   string `json:"error" firestore:",omitempty"`
		Message string `json:"message" firestore:",omitempty"`
	}
)

func updateUrlLastModTime(pageUrl string) error {
	pathArray := strings.SplitN(pageUrl, "/", 2)
	domain, path := pathArray[0], pathArray[1]

	url := StagingBaseUrl
	if domain == CommentoProdDomain {
		url = ProdBaseUrl
	}

	jsonStr := fmt.Sprintf("{\"url_entries\": {\"url\": %s, \"last_mod\": %s)}}", path, time.Now().Format(time.RFC3339))

	errResponse := ErrorResponseData{}

	_, err := sling.New().Client(&http.Client{}).Base(url).Set("Content-Type", "application/json").New().Post(EndPoint).BodyJSON([]byte(jsonStr)).Receive(nil, &errResponse)

	if err != nil {
		return err
	}
	return nil
}
