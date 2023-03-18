package tagger

import "os"

type Tagger interface {
	Tag(path string, info os.FileInfo) ([]string, error)
}
