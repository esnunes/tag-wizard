package tagger

import (
	"os"
	"strings"
)

var DefaultFileExtensions = map[string]string{
	"java":          ".java",
	"c":             ".c",
	"python":        ".py",
	"c++":           ".cpp",
	"c#":            ".cs",
	"visualbasic":   ".vb",
	"javascript":    ".js",
	"php":           ".php",
	"sql":           ".sql",
	"objective-c":   ".m",
	"swift":         ".swift",
	"groovy":        ".groovy",
	"assembly":      ".asm",
	"perl":          ".pl",
	"r":             ".r",
	"ruby":          ".rb",
	"dart":          ".dart",
	"delphi":        ".pas",
	"scratch":       ".sb3",
	"powershell":    ".ps1",
	"lua":           ".lua",
	"typescript":    ".ts",
	"kotlin":        ".kt",
	"go":            ".go",
	"shell":         ".sh",
	"scala":         ".scala",
	"fortran":       ".f",
	"abap":          ".abap",
	"vb.net":        ".vb",
	"crystal":       ".cr",
	"scheme":        ".scm",
	"apex":          ".cls",
	"haskell":       ".hs",
	"rust":          ".rs",
	"emacs-lisp":    ".el",
	"zig":           ".zig",
	"clojure":       ".clj",
	"clojurescript": ".cljs",
	"odin":          ".odin",
	"html":          "html",
	"css":           ".css",
	"terraform":     ".tf",
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
