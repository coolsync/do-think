package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

var (
	cli *elastic.Client
	ch  chan *LogData
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

func Init(addr string, chanMaxSize int, nums int) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}

	// init connect, get client obj
	cli, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		return
	}
	ch = make(chan *LogData, chanMaxSize)
	// 开启多个 goroutine
	for i := 0; i < nums; i++ {
		go sendToEs() // 不断从 SendToEsChan 内部chan 读取 struct携带的data(kafka data)
	}
	return nil
}

func SendToEsChan(logData *LogData) {
	ch <- logData
}

// 将 kafka data 发送到 es
func sendToEs() {
	for {
		select {
		case msg := <-ch: // struct chan
			// chain operate, insert obj to es
			put1, err := cli.Index().Index(msg.Topic).Type("xxx").BodyJson(msg).Do(context.Background())
			if err != nil {
				// Handle error
				// panic(err)
				fmt.Println(err)
				continue
			}
			fmt.Printf("Indexed users %q to index %q, type %q\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
