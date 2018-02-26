package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/dustinevan/adstxt"
	"github.com/dustinevan/go-utils/async"
	"context"
)

var canonicalMapFile = flag.String("canonicalmap", "", "optional canonical-adsys.json to overwrite internal canonical maps")

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
	if canonicalMapFile != nil && *canonicalMapFile != "" {
		err := adstxt.SetCanonicalMaps(*canonicalMapFile)
		if err != nil {
			log.Fatal(err, usage)
		}
	}

	if len(os.Args) == 1 {
		log.Fatal("not enough arguments", usage)
	}
	domainargs := os.Args[1:]
	if strings.Contains(domainargs[0], "canonicalmap") {
		domainargs = domainargs[1:]
	}
	if len(domainargs) == 0 {
		log.Fatal("not enough arguments", usage)
	}

	sem := async.NewSemaphore(100, context.Background())
	var wg sync.WaitGroup
	for _, url := range domainargs {
		wg.Add(1)
		go func(url string) {
			sem.Acquire()
			defer sem.Release()
			defer wg.Done()
			log.Printf("crawling %s", url)
			f, err := adstxt.Crawl(url)
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("successfully crawled %s", url)
			fmt.Println(f)
		}(url)
		// everything was kicked off
	}
	wg.Wait()
}

var usage = "\nUSAGE: atcrawl [-canonicalmap=\"file.json\"] url1 url2 ...\n"
