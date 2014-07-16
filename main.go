package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	clientInterface "github.com/yoed/yoed-client-interface"
)

type Handler struct {
	config *Config
}

type Config struct {
	clientInterface.Config
	ApiKey string `json:"api_key"`
}

func (c *Handler) Handle(username string) {
	resp, err := http.PostForm("http://api.justyo.co/yo/", url.Values{
		"api_token": {c.config.ApiKey},
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

func New() *Handler {

	c := &Handler{}

	if err := clientInterface.LoadConfig("./config.json", &c.config); err != nil {
		panic(fmt.Sprintf("failed loading config: %s", err))
	}

	return c
}

func main() {
	handler := New()
	client := clientInterface.New(handler, &handler.config.Config)
	client.Run()
}
