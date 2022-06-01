//Я понимаю, что я, кажется, усложнил условие, но мне правда было интересно разработать устойчивую программу

//Такая большая программа получилась из-за того, что я считываю данные из клавиатуры
//newReader не мультиплексируется и поэтому ReadString() блокирует всю программу
//(Можно было поискать библиотеки с мультиплексированием, но мне было интеерсно справится без неё)
// Также мы пытаемся обработать все данные, отправленные до истечения времени
// Поэтому мы закрываем канал Receiver-a и считываем все данные через range
// (Вспомним, что ридер блокирует всю программу) и мы не можем быть уверены, что он ничего не отправит в закрытый канал
// Поэтому я использовал третью функцию под названием middleware

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	senderCh := make(chan string)
	storageCh := make(chan string, 10)
	quit := make(chan struct{})
	quitted := make(chan struct{})
	numSecondsWait := flag.Int("seconds", 5, "num seconds")

	go sender(senderCh)
	go middleware(senderCh, storageCh, quit)
	go receiver(storageCh, quit, quitted)

	time.Sleep(time.Duration(*numSecondsWait) * time.Second)
	//Отправить сообщение вообще всем
	close(quit)
	<-quitted
}

func sender(senderCh chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		senderCh <- text
	}
}

func receiver(storageCh <-chan string, quit, quitted chan struct{}) {
	for {
		select {
		case mes := <-storageCh:
			fmt.Printf("%s", mes)
		case <-quit:
			for mes := range storageCh {
				fmt.Printf("%s", mes)
			}
			//заставить основную программу ждать завершения
			quitted <- struct{}{}
			return
		}

	}

}

func middleware(senderCh <-chan string, storageCh chan<- string, quit <-chan struct{}) {
	for {
		select {
		case mes := <-senderCh:
			storageCh <- mes
		case <-quit:
			close(storageCh)
			return
		}

	}
}
