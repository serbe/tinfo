package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/serbe/adb"
)

var (
	logErrors bool
	db        *adb.ADB
	cfg       Config
)

// Config all vars
type Config struct {
	Base struct {
		LogErr   bool   `json:"logerr"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Host     string `json:"host"`
	} `json:"base"`
	Bot struct {
		Token string `json:"token"`
	} `json:"bot"`
}

func getConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("getConfig ReadFile", err)
	}
	if err = json.Unmarshal(file, &cfg); err != nil {
		log.Fatal("getConfig Unmarshal", err)
	}
	logErrors = cfg.Base.LogErr
	if cfg.Base.Name == "" {
		err = errors.New("empty database name in config")
		log.Fatal("getConfig", err)
	}
}

func errmsg(str string, err error) {
	if logErrors {
		log.Println("Error in", str, err)
	}
}

func errChkMsg(str string, err error) {
	if logErrors && err != nil {
		log.Println("Error in", str, err)
	}
}

func getArgInt(str string) int {
	var (
		result int
		err    error
	)
	text := strings.Trim(str, " ")
	text = strings.Replace(text, "  ", " ", -1)
	split := strings.Split(text, " ")
	if len(split) == 2 {
		result, err = strconv.Atoi(split[1])
		errChkMsg("getArgInt Atoi", err)
	}
	if result > 100 {
		result = 100
	} else if result < 1 {
		result = 1
	}
	return result
}

func getArgString(str string) string {
	var result string
	text := strings.Trim(str, " ")
	text = strings.Replace(text, "  ", " ", -1)
	split := strings.Split(text, " ")
	if len(split) == 2 {
		result = split[1]
	}
	return result
}
