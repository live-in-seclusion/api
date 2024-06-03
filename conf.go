package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/olivere/elastic.v3"
)

//Config define config
var (
	Config config
	Log    *log.Logger
	ES     *elastic.Client
)

type config struct {
	Elastic string `json:"elastic"`
	Address string `json:"address"`
}

//Init utilsl
func Init() {
	initLog()
	initConfig()
	initElastic()
}

func initElastic() {
	client, err := elastic.NewClient(elastic.SetURL(Config.Elastic))
	exit(err)
	ES = client
	ES.CreateIndex("torrent").Do()
}

func initConfig() {
	f, err := os.Open("config/api.conf")
	exit(err)
	b, err := ioutil.ReadAll(f)
	exit(err)
	err = json.Unmarshal(b, &Config)
	exit(err)
}

func initLog() {
	Log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func exit(err error) {
	if err != nil {
		Log.Fatalln(err)
	}
}
