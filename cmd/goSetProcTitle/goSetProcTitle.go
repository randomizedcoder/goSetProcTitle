package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/erikdubbelboer/gspt"
)

const(
    freqCst = 10 * time.Second
	textCst = "hello world"
	baseCst = "grep-for-me "
)

var (
	// Passed by "go build -ldflags" for the show version
	commit string
	date   string
)


func main() {

    version := flag.Bool("version", false, "version")
    freq := flag.Duration("freq", freqCst,"frequency to change he process title")
	text := flag.String("text",textCst,"text for process line, defaults to the time")
	base := flag.String("base",baseCst,"base text for process line")

	flag.Parse()

    if *version {
		fmt.Println("commit:", commit, "\tdate(UTC):", date)
		os.Exit(0)
	}


	ticker := time.NewTicker(*freq)

	for range ticker.C{
		t := getProcTitle(*text,*base)
		log.Println("SetProcTitle:",t)
		gspt.SetProcTitle(t)
	}
}

func getProcTitle(text string, base string) (title string){

	if text != textCst{
		title = base + text
		return title
	}

	title = base + time.Now().UTC().Format(time.RFC3339)

	return title
}