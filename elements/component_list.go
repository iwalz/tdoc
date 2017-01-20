package elements

import (
	"strings"

	"github.com/spf13/afero"
)

type ComponentsList struct {
	dir        string            // Theme directory
	components map[string]string // Internal component store
	fs         afero.Fs          // File system mock
}

func NewComponentsList(s string) *ComponentsList {
	c := make(map[string]string)
	c["node"] = ""
	c["actor"] = ""
	c["cloud"] = ""
	f := afero.NewOsFs()
	cl := &ComponentsList{dir: s, components: c, fs: f}

	return cl
}

func (cl *ComponentsList) Parse() error {
	return cl.readDir()
}

func (cl *ComponentsList) Exists(s string) bool {
	if _, ok := cl.components[s]; ok {
		return true
	}

	return false
}

func (cl *ComponentsList) readDir() error {
	if cl.dir != "" {
		files, err := afero.ReadDir(cl.fs, cl.dir)
		if err != nil {
			return err
		}

		for _, v := range files {
			if v.IsDir() {
				f, _ := afero.ReadDir(cl.fs, cl.dir+"/"+v.Name())

				for _, file := range f {
					name := strings.Replace(file.Name(), ".svg", "", 1)
					cl.components[v.Name()+"_"+name] = cl.dir + "/" + file.Name()
				}
			}
			if strings.HasSuffix(v.Name(), ".svg") {
				name := strings.Replace(v.Name(), ".svg", "", 1)
				cl.components[name] = cl.dir + "/" + v.Name()
			}
		}
	}

	return nil
}
