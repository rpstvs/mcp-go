package agentdiscovery

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Agent struct {
	Name        string
	Description string
	Model       string
}

func loadAgents() {
	filepath.WalkDir("Root", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && filepath.Ext(d.Name()) == ".md" {
			f, err := os.Open(path)

			if err != nil {
				return err
			}
			f.
		}
	})
}
