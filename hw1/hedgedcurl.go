package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"hw1/handlers"
	"os"
	"strconv"
	"time"
)

func mai() {
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
		writer := bufio.NewWriter(os.Stdout)
		defer writer.Flush()
		writer.WriteString("timeout is ")
		writer.WriteString((time.Duration(timeout) * time.Second).String())
		writer.WriteString("\n")
		writer.Flush()
		urls := flag.Args()
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		resultChan := make(chan handlers.Output)

		for _, url := range urls {
			go handlers.DoGetWithContext(ctx, url, resultChan)

		}

		total := len(urls)
		errorsCount := 0
		for {
			select {
			case res := <-resultChan:
				if res.Err == nil {
					cancel()
					writer.WriteString(fmt.Sprintf("From: %s\n", res.From))
					writer.WriteString(fmt.Sprintf("Status code: %s\n", strconv.Itoa(res.Status)))
					writer.WriteString(fmt.Sprintf("Headers: %s\n", res.Headers))
					writer.WriteString(fmt.Sprintf("Body: %s\n", res.Body))
					return
				} else {
					errorsCount++
					if errorsCount == total {
						os.Exit(228)
					}
				}
			case <-ctx.Done():
				os.Exit(228)
			}

		}

	}
}

func main() {
	os.Exit(228)
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
