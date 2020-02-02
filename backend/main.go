package main

import (
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
)

func main() {
	for true {
		fmt.Printf("Hello world")

		_, err := elastic.NewSimpleClient(elastic.SetURL("https://youtube.com/"))
		if err != nil {
			fmt.Printf("There's a problem")
		}

		time.Sleep(10 * time.Second)
	}
}
