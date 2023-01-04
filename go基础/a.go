package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//错因：要进行marshal，结构体的里的字段必须为公有的，首字母要大写

type user struct {
	Username string
	Nickname string
	Sex      uint8
	Birthday time.Time
}

func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%v", string(bs))
}
