package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/oauth1"
)

// Client - basic structure of a client request
type Client struct {
	url    string
	client *http.Client
}

// NewClient - this returns new client, based on the request
func NewClient(request map[string]interface{}) (*Client, error) {
	var mainError error = nil
	var err error = nil
	httpClient := &http.Client{Timeout: 5 * time.Second}
	host := request["host"].(string)

	if oauth1Config, ok := request["oauth1"]; ok {
		config := oauth1Config.(map[string]interface{})
		httpClient, err = returnOauth1Client(config)
		if err != nil {
			mainError = err
		}
	}

	return &Client{
		url:    host,
		client: httpClient,
	}, mainError
}

func returnOauth1Client(oauth1Config map[string]interface{}) (*http.Client, error) {
	consumerKey := oauth1Config["consumerKey"].(string)
	consumerSecret := oauth1Config["consumerSecret"].(string)
	accessToken := oauth1Config["accessToken"].(string)
	accessSecret := oauth1Config["accessSecret"].(string)
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return nil, errors.New("Missing required environment variable")
	}
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient, nil
}

// Get - finally performs the request
func (c *Client) Get() (interface{}, error) {
	resp, err := c.client.Get(c.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return mapContent(body, resp.Header.Get("Content-Type"))
}

// Post - finally performs the request
func (c *Client) Post(content string, contentType string) (interface{}, error) {
	payload := []byte(content)
	resp, err := c.client.Post(c.url, contentType, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return mapContent(body, resp.Header.Get("Content-Type"))
}

// mapContent - this maps the reveiced bytes to a fitting type, matching the content-type
func mapContent(body []byte, contentType string) (interface{}, error) {
	if strings.Contains(contentType, "/json") {
		var jsonBody interface{}
		err := json.Unmarshal(body, &jsonBody)
		if err != nil {
			return nil, err
		}
		return jsonBody, nil
	}
	if strings.Contains(contentType, "/text") {
		return string(body), nil
	}
	return body, nil
}

// Debug - outputs some debug infos
func (c *Client) Debug() {
	log.Printf("%v", c)
}
