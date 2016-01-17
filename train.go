package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
)

//電車運行情報の情報源
const TRAIN_URL = "http://transit.yahoo.co.jp/traininfo/area/4/"
const SLACK_URL = "https://slack.com/api/chat.postMessage"
const CONFIG_FILENAME = "./config.json"

type Config struct {
	Trains []string
	Slack  map[string]string
}

//関東の電車の運行情報全て取得
func Status() map[string]string {
	train_status := make(map[string]string)
	doc, _ := goquery.NewDocument(TRAIN_URL)
	doc.Find(".elmTblLstLine table tr").Each(func(_ int, tr *goquery.Selection) {
		var train, status string
		tr.Find("td").Each(func(i int, td *goquery.Selection) {
			switch i {
			case 0: //電車名
				train = td.Text()
			case 1: //電車の状況
				status = td.Text()
			}
		})
		train_status[train] = status
	})
	return train_status
}

//設定ファイルの読み込み
func Parse(filename string) (Config, error) {
	var c Config
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("error:%v", err)
		return c, err
	}
	err = json.Unmarshal(jsonString, &c)
	if err != nil {
		fmt.Printf("error:%v", err)
		return c, err
	}
	return c, nil
}

func main() {
	config, err := Parse(CONFIG_FILENAME)
	fmt.Printf("%v", config)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	status := Status()
	status_str := ""
	for _, train := range config.Trains {
		status_str += fmt.Sprintf("%s : %s\n", train, status[train])
	}

	//Slackへ通知
	values := url.Values{}
	values.Add("token", config.Slack["Token"])
	values.Add("channel", config.Slack["Channel"])
	values.Add("text", status_str)
	http.PostForm(SLACK_URL, values)
}
