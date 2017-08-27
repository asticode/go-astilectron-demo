package main

import (
	"io/ioutil"
	"os"
	"os/user"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// ListItem represents a list item
type ListItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// handleMessages handles messages
func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "get.list":
		// Get user
		var u *user.User
		if u, err = user.Current(); err != nil {
			return
		}

		// Read dir
		var files []os.FileInfo
		if files, err = ioutil.ReadDir(u.HomeDir); err != nil {
			return
		}

		// Build list items
		var items []ListItem
		for _, f := range files {
			var item = ListItem{Name: f.Name()}
			if f.IsDir() {
				item.Type = "dir"
			} else {
				item.Type = "file"
			}
			items = append(items, item)
		}
		payload = items
	}
	return
}
