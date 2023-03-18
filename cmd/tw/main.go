package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	totalTags := map[string]struct{}{}

	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, tagger := range taggers {
			tags, err := tagger.Tag(path, info)
			if err != nil {
				return err
			}
			if tags == nil {
				continue
			}
			for _, v := range tags {
				totalTags[v] = struct{}{}
			}
		}
		return nil
	}

	if err = filepath.Walk(wd, walk); err != nil {
		log.Fatalf("failed to walk through files in %v: %v", wd, err)
	}

	fmt.Println(strings.Join(g.Keys(totalTags), " "))
}
