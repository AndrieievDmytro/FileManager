package main

import (
	"flag"
	"fmt"
	"time"
)

//Global variables
var (
	Root string = "." //Root path
	File string       // File command line argument
	Dir  string       // Folder comman line argument
)

func init() {
	flag.StringVar(&Root, "r", Root, "Root. Default path: `.`")
	flag.StringVar(&File, "f", File, "Search a file(s) with the specified name")
	flag.StringVar(&Dir, "d", Dir, "Search a directory(s) with the specified name")
}

func main() {
	start := time.Now()
	flag.Parse() //Parse comandline arguments

	//************* Map implementation *************//
	// foldersFilesList := new(Folder)
	// foldersFilesList.fileList = make(map[string][]string)
	// foldersFilesList.readDir("C:/", 0)
	// fmt.Println(len(foldersFilesList.fileList))
	// fmt.Println(foldersFilesList.fileList)
	// fmt.Println("Done.")
	//**********************************************//

	//********************* Tree implementation ****************//
	folderList := new(FolderTree)
	folderList.readDirT(Root)
	folderList.writeToFileJson("files.json")
	fmt.Println("Folders: ", folderList.FolderNum, "Files: ", len(folderList.Folders))
	// *********************************************************//

	timeTrack(start, "File reading")
}
