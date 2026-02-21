package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"hw1/handlers"
	"os"
	"time"
)

func main() {
	var helpBool bool
	var timeout int
	flag.BoolVar(&helpBool, "h", false, "show help")
	flag.BoolVar(&helpBool, "help", false, "show help")
	flag.IntVar(&timeout, "t", 15, "timeout for each request")
	flag.IntVar(&timeout, "timeout", 15, "timeout for each request")
	flag.Parse()

	if helpBool {
		help()
		return
	} else {
		urls := flag.Args()
		time.AfterFunc(time.Duration(timeout)*time.Second, func() {
			fmt.Println("timeout exceeded - ни один запрос не завершился вовремя")
			os.Exit(228)
		})
		if err := doHedgedShit(urls, time.Duration(timeout)*time.Second); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func help() {
	helpStr :=
		`Доступные команды:
hedgecurl [-t / --timeout] [urls...]

-t/--timeout значение таймаута (если не передавать значение по дефолту 15 сек)
urls массив url к которым будут переданы запросы
`
	fmt.Println(helpStr)
}

func doHedgedShit(urls []string, timeout time.Duration) error {
	if len(urls) == 0 {
		return errors.New("empty urls input")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resultChan := make(chan *handlers.Output, 1)

	for _, url := range urls {
		go worker(ctx, timeout, url, resultChan)
	}

	select {
	case result := <-resultChan:
		fmt.Println("Статус:", result.Status)
		fmt.Println("Заголовки:", result.Headers)
		fmt.Println("Тело:", result.Body)
		return nil

	case <-ctx.Done():
		fmt.Println("timeout exceeded - ни один запрос не завершился вовремя")
		os.Exit(228)
		return nil
	}
}

func worker(ctx context.Context, timeout time.Duration, url string, resultChan chan *handlers.Output) {
	res, err := handlers.DoGet(ctx, timeout, url)
	if err != nil {
		return
	}
	select {
	case resultChan <- res:
	default:
	}
}
