package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

// 结构内可少于文件，但是包括[]必须有总结构，如 tomlConfig
type tomlConfig struct {
	Title   string // 可无，下同
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string // 可无，下同
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

func main() {
	var conf tomlConfig
	if _, err := toml.DecodeFile("./eg.toml", &conf); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf)
}
