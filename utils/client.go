package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
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
func (c *Client) Get() (string, error) {
	resp, err := c.client.Get(c.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Post - finally performs the request
func (c *Client) Post(content string, contentType string) (string, error) {
	payload := []byte(content)
	resp, err := c.client.Post(c.url, contentType, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Debug - outputs some debug infos
func (c *Client) Debug() {
	log.Printf("%s", c)
}
