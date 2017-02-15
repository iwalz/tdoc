package outputs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

type Output interface {
	HandleFile(string) error
}

type Executor struct {
	fs  afero.Fs // File system mock
	ext string   // File extension
}

func NewExecutor(fs afero.Fs, ext string) *Executor {
	return &Executor{fs: fs, ext: ext}
}

// Handle the command line args as a list of files and directories
func (e *Executor) Exec(o Output, args []string) error {
	for _, arg := range args {
		isDir, err := afero.IsDir(e.fs, arg)
		if err != nil {
			return err
		}
		if isDir {
			e.handleDir(o, arg)
			continue
		}

		exists, err := afero.Exists(e.fs, arg)
		if err != nil {
			return err
		}
		if exists {
			o.HandleFile(arg)
		}
	}

	return nil
}

func handleFile(e string, o Output) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip if file doesn't end with given extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), "."+e) {
			o.HandleFile(path)
		}

		return nil
	}
}

func (e *Executor) handleDir(o Output, d string) error {
	err := afero.Walk(e.fs, d, handleFile(e.ext, o))
	if err != nil {
		return err
	}

	return nil
}
