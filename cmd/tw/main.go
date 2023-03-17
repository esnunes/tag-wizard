package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var DefaultFileExtensions = map[string]string{
	"java":        ".java",
	"c":           ".c",
	"python":      ".py",
	"c++":         ".cpp",
	"c#":          ".cs",
	"visualbasic": ".vb",
	"javascript":  ".js",
	"php":         ".php",
	"sql":         ".sql",
	"objective-c": ".m",
	"swift":       ".swift",
	"groovy":      ".groovy",
	"assembly":    ".asm",
	"perl":        ".pl",
	"r":           ".r",
	"ruby":        ".rb",
	"dart":        ".dart",
	"delphi":      ".pas",
	"scratch":     ".sb3",
	"powershell":  ".ps1",
	"lua":         ".lua",
	"typescript":  ".ts",
	"kotlin":      ".kt",
	"go":          ".go",
	"shell":       ".sh",
	"scala":       ".scala",
	"fortran":     ".f",
	"abap":        ".abap",
	"vb.net":      ".vb",
	"crystal":     ".cr",
	"scheme":      ".scm",
	"bash":        ".sh",
	"apex":        ".cls",
	"haskell":     ".hs",
}

type Tagger interface {
	Tag(path string, info os.FileInfo) ([]string, error)
}

type FileExtensionTagger struct {
	Extensions map[string]string
}

func (f *FileExtensionTagger) Tag(path string, info os.FileInfo) ([]string, error) {
	if f.Extensions == nil {
		f.Extensions = make(map[string]string)
		for k, v := range DefaultFileExtensions {
			f.Extensions[k] = v
		}
	}

	if info.IsDir() {
		return nil, nil
	}

	for name, ext := range f.Extensions {
		if strings.HasSuffix(path, ext) {
			delete(f.Extensions, name)
			return []string{name}, nil
		}
	}
	return nil, nil
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	taggers := []Tagger{
		&FileExtensionTagger{},
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
