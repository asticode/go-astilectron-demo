package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/asticode/go-astichartjs"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// PayloadExplore represents the payload of the explore event
type PayloadExplore struct {
	Dirs       []PayloadDir       `json:"dirs"`
	Files      *astichartjs.Chart `json:"files,omitempty"`
	FilesCount int                `json:"files_count"`
	FilesSize  string             `json:"files_size"`
	Path       string             `json:"path"`
	PathDir    PayloadDir         `json:"path_dir"`
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
		}

		// Previous dir
		if filepath.Dir(path) != path {
			p.Dirs = append(p.Dirs, PayloadDir{
				Name: "..",
				Path: filepath.Dir(path),
			})
		}

		// Loop through files
		var sizes []int
		var sizesMap = make(map[int][]string)
		var filesSize int64
		for _, f := range files {
			if f.IsDir() {
				p.Dirs = append(p.Dirs, PayloadDir{
					Name: f.Name(),
					Path: filepath.Join(path, f.Name()),
				})
			} else {
				var s = int(f.Size())
				sizes = append(sizes, s)
				sizesMap[s] = append(sizesMap[s], f.Name())
				p.FilesCount++
				filesSize += f.Size()
			}
		}

		// Prepare files size
		if filesSize < 1e3 {
			p.FilesSize = strconv.Itoa(int(filesSize)) + "b"
		} else if filesSize < 1e6 {
			p.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024), 'f', 0, 64) + "kb"
		} else if filesSize < 1e9 {
			p.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024), 'f', 0, 64) + "Mb"
		} else {
			p.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024*1024), 'f', 0, 64) + "Gb"
		}

		// Prepare files chart
		sort.Ints(sizes)
		if len(sizes) > 0 {
			p.Files = &astichartjs.Chart{
				Data: &astichartjs.Data{Datasets: []astichartjs.Dataset{{
					BackgroundColor: []string{
						astichartjs.ChartBackgroundColorYellow,
						astichartjs.ChartBackgroundColorGreen,
						astichartjs.ChartBackgroundColorRed,
						astichartjs.ChartBackgroundColorBlue,
						astichartjs.ChartBackgroundColorPurple,
					},
					BorderColor: []string{
						astichartjs.ChartBorderColorYellow,
						astichartjs.ChartBorderColorGreen,
						astichartjs.ChartBorderColorRed,
						astichartjs.ChartBorderColorBlue,
						astichartjs.ChartBorderColorPurple,
					},
				}}},
				Type: astichartjs.ChartTypePie,
			}
			var sizeOther int
			for i := len(sizes) - 1; i >= 0; i-- {
				for _, l := range sizesMap[sizes[i]] {
					if len(p.Files.Data.Labels) < 4 {
						p.Files.Data.Datasets[0].Data = append(p.Files.Data.Datasets[0].Data, sizes[i])
						p.Files.Data.Labels = append(p.Files.Data.Labels, l)
					} else {
						sizeOther += sizes[i]
					}
				}
			}
			if sizeOther > 0 {
				p.Files.Data.Datasets[0].Data = append(p.Files.Data.Datasets[0].Data, sizeOther)
				p.Files.Data.Labels = append(p.Files.Data.Labels, "other")
			}
		}
		payload = p
	}
	return
}
