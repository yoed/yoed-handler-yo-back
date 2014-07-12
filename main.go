package main

import (
	clientInterface "github.com/yoed/yoed-client-interface"
	"net/http"
	"net/url"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type YoBackYoedClient struct {
	Config *YoBackYoedClientConfig
}

type YoBackYoedClientConfig struct {
	clientInterface.YoedClientConfig
	ApiKey string `json:"apiKey"`
}

func (c *YoBackYoedClient) loadConfig(configPath string) (*YoBackYoedClientConfig, error) {

	configFile, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	configJson, err := ioutil.ReadAll(configFile)

	if err != nil {
		return nil, err
	}

	config := &YoBackYoedClientConfig{}

	if err := json.Unmarshal(configJson, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *YoBackYoedClient) Handle(username string) {
	resp, err := http.PostForm("http://api.justyo.co/yo/", url.Values{
		"api_token": {c.Config.ApiKey},
		"username":  {username},
	})

	log.Printf("Yo-ing back %s", username)

	if err != nil {
		log.Printf("yobackHandler: %s", err)
		return
	}

	defer resp.Body.Close()

	log.Printf("yobackHandler: %s", resp.Status)

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Printf("yobackHandler: %s", err)
	} else {
		log.Printf("yobackHandler: %s", string(body))
	}
}

func (c *YoBackYoedClient) GetConfig() (*clientInterface.YoedClientConfig) {
	return &clientInterface.YoedClientConfig{Listen:c.Config.Listen, ServerUrl: c.Config.ServerUrl}
}

func NewYoBackYoedClient() (*YoBackYoedClient, error) {
	c := &YoBackYoedClient{}
	config, err := c.loadConfig("./config.json")

	if err != nil {
		panic(fmt.Sprintf("failed loading config: %s", err))
	}

	c.Config = config

	return c, nil
}

func main() {
	c, _ := NewYoBackYoedClient()

	clientInterface.Run(c)
}