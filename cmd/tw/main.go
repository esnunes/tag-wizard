package main

import (
	"log"
	"os"
	"path/filepath"

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

	totalTags := []string{}

	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, tagger := range taggers {
			tags, err := tagger.Tag(path, info)
			if err != nil {
				return err
			}
			if tags != nil {
				totalTags = append(totalTags, tags...)
			}
		}
		return nil
	}

	if err = filepath.Walk(wd, walk); err != nil {
		log.Fatalf("failed to walk through files in %v: %v", wd, err)
	}

	log.Printf("tags: %v", totalTags)
}
