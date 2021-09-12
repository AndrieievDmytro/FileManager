package main

import "sync"

// type Folder struct {
// 	mu          sync.Mutex
// 	fileList    map[string][]string
// 	dirCount    int64
// 	folderCount int64
// 	grNum       int64
// }

type FolderTree struct {
	Mute      sync.Mutex   `json:"-"`
	GoRtNum   int64        `json:"-"`
	FolderNum int64        `json:"-"`
	Name      string       `json:"Path"`
	File      string       `json:"File"`
	Folders   []FolderTree `json:",omitempty"`
}
