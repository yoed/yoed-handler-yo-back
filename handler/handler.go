package handler

import (
	"log"
	"net/http"
	"net/url"
	httpInterface "github.com/yoed/yoed-http-interface"
	"io/ioutil"
)

type Handler struct {
	Config *Config
}

type Config struct {
	httpInterface.Config
	ApiKey string `json:"api_key"`
}

func (c *Handler) Handle(username string) {
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

func New() *Handler {

	c := &Handler{}

	if err := httpInterface.LoadConfig("./config.json", &c.Config); err != nil {
		log.Fatalf("failed loading config: %s", err)
	}

	return c
}