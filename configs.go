package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var MConfigs *MConfig
var ShowItem []string
var ShowDetail []string

type MConfig struct {
	Prefix       string
	Sort         string
	DetailPrefix string
	DetailSort   string
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
func (jst *JsonStruct) Load(filename string, v interface{}) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
func padNumberWithZero(value uint32) string {
	return fmt.Sprintf("%09d", value)
}

func loadConfig() {
	JsonParse := NewJsonStruct()
	JsonParse.Load("config.json", &MConfigs)
	ShowItem = strings.Split(strings.ReplaceAll(readFile("./ShowItem.txt"), "\r\n", ""), ",")
	ShowDetail = strings.Split(strings.ReplaceAll(readFile("./ShowDetail.txt"), "\r\n", ""), ",")

}
func StrHtmlTop(x string) string {
	return strings.ReplaceAll("<!DOCTYPE html><html lang='zh-CN'><head><link rel='stylesheet' type='text/css' href='tb.css'></head><br>", "tb.css", x)

}

/*
	id, _ := uuid.NewUUID()
	t := id.Time()
	sec, nsec := t.UnixTime()
	timeStamp := time.Unix(sec, nsec)
	//print(padNumberWithZero(563))
	fmt.Printf("Your unique id is: %s \n", id)
	fmt.Printf("The id was generated at: %v \n", timeStamp)

now := time.Now()
beginningOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
passedTimeNs := now.UnixNano() - beginningOfDay.UnixNano()
log.Println("今日已过（纳秒）：", passedTimeNs)
log.Println("今日已过（微秒）：", passedTimeNs/1e3)
log.Println("今日已过（毫秒）：", passedTimeNs/1e6)
log.Println("今日已过（秒）：", passedTimeNs/1e9)
now = time.Now()
beginningOfDay = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
passedTimeNs = now.UnixNano() - beginningOfDay.UnixNano()
log.Println("今日已过（微秒）：", passedTimeNs/1e3)
*/
