package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HTTPGet(url string, timeout time.Duration) (content []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	ctx, cancel_func := context.WithTimeout(context.Background(), timeout)
	request = request.WithContext(ctx)
	defer cancel_func()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("INVALID RESPONSE; status: %s", response.Status)
	}

	return ioutil.ReadAll(response.Body)
}

func main() {
	var url, output_path string
	var timeout time.Duration
	flag.StringVar(&url, "url", "https://yandex.ru", "url")
	flag.DurationVar(&timeout, "timeout", 7*time.Second, "duration of program")
	flag.StringVar(&output_path, "output", "./test.html", "file in which save")

	flag.Parse()

	content, err := HTTPGet(url, timeout)
	if err != nil {
		log.Fatalln("HTTPGET: ", err)
	}
	err = ioutil.WriteFile(output_path, content, 0666)
	if err != nil {
		log.Fatalln("WriteFile: ", err)
	}

}
