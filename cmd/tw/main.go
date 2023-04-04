package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/boyter/gocodewalker"

	"github.com/esnunes/tag-wizard/pkg/g"
	"github.com/esnunes/tag-wizard/pkg/tagger"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	taggers := []tagger.Tagger{
		&tagger.FileExtensionTagger{},
	}

	queue := make(chan *gocodewalker.File, 1_000)
	walker := gocodewalker.NewFileWalker(wd, queue)
	walker.SetErrorHandler(func(e error) bool {
		log.Printf("failed to read file: %v", e.Error())
		return true
	})

	go func() {
		if err := walker.Start(); err != nil {
			log.Printf("failed to walk through files: %v", err)
		}
	}()

	totalTags := map[string]struct{}{}

	for f := range queue {
		for _, tagger := range taggers {
			file, err := os.Open(f.Location)
			if err != nil {
				log.Fatalf("failed to open file %v: %v", f.Location, err)
			}

			info, err := file.Stat()
			if err != nil {
				log.Fatalf("failed to load file stats for %v: %v", f.Location, err)
			}

			tags, err := tagger.Tag(f.Location, info)
			if err != nil {
				log.Fatalf("failed to tag file %v: %v", f.Location, err)
			}
			if tags == nil {
				continue
			}
			for _, v := range tags {
				totalTags[v] = struct{}{}
			}
		}
	}
	tags := g.Keys(totalTags)
	sort.Strings(tags)
	fmt.Println(strings.Join(tags, " "))
}
