package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	textCh := make(chan string)
	quitCrtlC := make(chan os.Signal, 1)
	quitGoroutines := make(chan interface{}, 1)
	numWorkers := flag.Int("num-workers", 1, "num workers")
	flag.Parse()

	//Используем wg, чтобы быть уверенными в том, что все данные запроцессятся
	var wg sync.WaitGroup

	wg.Add(*numWorkers)
	for i := 0; i < *numWorkers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case mes := <-textCh:
					fmt.Fprint(os.Stdout, mes)

				//flag to know when we have to stop
				case <-quitGoroutines:
					for mes := range textCh {
						fmt.Fprint(os.Stdout, mes)
					}
					return
				}

			}

		}()
	}

	//Func for reading
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			textCh <- text
		}
	}()

	//Getting ctrl + c
	signal.Notify(quitCrtlC, syscall.SIGTERM, syscall.SIGINT)
	<-quitCrtlC

	close(textCh)

	//Выражение ниже позволяет завершить ВСЕ горутины, которые ждали сообщений
	close(quitGoroutines)
	wg.Wait()

}
