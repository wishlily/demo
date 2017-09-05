package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	{
		// 从文件读取 Json 数据
		type DefaultClient struct {
			Id          string
			Secret      string
			RedirectUri string
			UserData    interface{}
		}
		clients := make(map[string]DefaultClient)
		rd, err := os.Open("./a.json")
		if err != nil {
			panic(err)
		}
		decoder := json.NewDecoder(rd)
		if err := decoder.Decode(&clients); err != nil {
			panic(err)
		}
		fmt.Println(clients)
		v, ok := clients["0001"]
		if !ok {
			panic(ok)
		}
		fmt.Println(v.UserData)
		t := reflect.TypeOf(v.UserData)
		fmt.Println(t.Kind())
		d, ok := v.UserData.(map[string]interface{})
		fmt.Println(ok, d)
		vv := d["name2"]
		fmt.Println(vv)
	}
	{
		// 结构体
		// json 冒号后不要有空格，名字要大写
		type MySt struct {
			A string  `json:"a"`
			B int     `json:"b"`
			C float32 `json:"c"`
		}
		b, err := json.Marshal(MySt{"MySt", 1, 0.145})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d,%s\n", len(b), string(b))
		var v MySt
		if err := json.Unmarshal(b, &v); err != nil {
			panic(err)
		}
		fmt.Println(v)
	}
}
