package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rovnyart/diversity/pkg/icon"

	"github.com/getlantern/systray"
	"github.com/go-co-op/gocron"
	"github.com/reujab/wallpaper"
	"github.com/spf13/viper"
)

func main() {
	systray.Run(onReady, onExit)
}

const API_URL = "https://api.unsplash.com/photos/random/"

type Links struct {
	Download string `json:"download"`
}

type Photo struct {
	Id    string `json:"id"`
	Links Links  `json:"links"`
}

type Config struct {
	Apikey   string `yaml:"apikey"`
	Scope    string `yaml:"scope"`
	Schedule string `yaml:"schedule"`
}

var conf *Config

func onReady() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	conf = &Config{}

	err = viper.Unmarshal(&conf)

	if err != nil {
		log.Fatal(err)
	}

	systray.SetIcon(icon.Data)
	systray.SetTitle("Diversity")
	systray.SetTooltip("Diversity wallpaper changer")

	change := systray.AddMenuItem("Change current wallpaper", "Change current wallpaper")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Exit", "Exit app")

	s := gocron.NewScheduler(time.UTC)
	s.Every(conf.Schedule).Do(changeWallpaper)
	s.StartAsync()

	//
	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()

	go func() {
		for {
			select {
			case <-change.ClickedCh:
				changeWallpaper()
			}
		}
	}()

}

func onExit() {
	// cleanup
}

func changeWallpaper() {
	resp, err := http.Get(API_URL + "?client_id=" + conf.Apikey + "&query=" + conf.Scope)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Panicln("Request error")

		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var pic Photo

	json.Unmarshal(bodyBytes, &pic)

	err = wallpaper.SetFromURL(pic.Links.Download)

	if err != nil {
		log.Println(err)
	}

}
