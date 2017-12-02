package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"sort"

	"github.com/asticode/go-astichartjs"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// PayloadExplore represents the payload of the explore event
type PayloadExplore struct {
	Dirs    []PayloadDir       `json:"dirs"`
	Files   *astichartjs.Chart `json:"files,omitempty"`
	Path    string             `json:"path"`
	PathDir PayloadDir         `json:"path_dir"`
}

// PayloadDir represents a dir payload
type PayloadDir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "explore":
		// If no path, path is home dir
		var path string
		if len(m.Payload) == 0 {
			var u *user.User
			if u, err = user.Current(); err != nil {
				payload = err.Error()
				return
			}
			path = u.HomeDir
		} else {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &path); err != nil {
				payload = err.Error()
				return
			}
		}

		// Read dir
		var files []os.FileInfo
		if files, err = ioutil.ReadDir(path); err != nil {
			payload = err.Error()
			return
		}

		// Init payload
		var p = PayloadExplore{
			Dirs: []PayloadDir{},
			Path: path,
			PathDir: PayloadDir{
				Name: "..",
				Path: filepath.Dir(path),
			},
		}

		// Loop through files
		var sizes []int
		var sizesMap = make(map[int]string)
		for _, f := range files {
			if f.IsDir() {
				p.Dirs = append(p.Dirs, PayloadDir{
					Name: f.Name(),
					Path: filepath.Join(path, f.Name()),
				})
			} else {
				var s = int(f.Size())
				sizes = append(sizes, s)
				sizesMap[s] = f.Name()
			}
		}

		// Prepare files chart
		sort.Ints(sizes)
		if len(sizes) > 0 {
			p.Files = &astichartjs.Chart{
				Data: &astichartjs.Data{Datasets: []astichartjs.Dataset{{
					BackgroundColor: []string{
						astichartjs.ChartBackgroundColorBlue,
						astichartjs.ChartBackgroundColorGreen,
						astichartjs.ChartBackgroundColorOrange,
						astichartjs.ChartBackgroundColorRed,
						astichartjs.ChartBackgroundColorYellow,
					},
					BorderColor: []string{
						astichartjs.ChartBorderColorBlue,
						astichartjs.ChartBorderColorGreen,
						astichartjs.ChartBorderColorOrange,
						astichartjs.ChartBorderColorRed,
						astichartjs.ChartBorderColorYellow,
					},
				}}},
				Type: astichartjs.ChartTypePie,
			}
			for i := len(sizes) - 1; i > len(sizes)-6 && i >= 0; i-- {
				p.Files.Data.Datasets[0].Data = append(p.Files.Data.Datasets[0].Data, sizes[i])
				p.Files.Data.Labels = append(p.Files.Data.Labels, sizesMap[sizes[i]])
			}
		}
		payload = p
	}
	return
}
