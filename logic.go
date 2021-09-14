package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func errCheck(err error) bool {
	result := true
	if err != nil {
		// fmt.Println(err)
		result = false
	}
	return result
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

//****** Check (by path) if the folder or file exists *******//

// func fileExists(path string) (bool, error) {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return true, nil
// 	}
// 	if os.IsNotExist(err) {
// 		return false, nil
// 	}
// 	return false, err
// }

// ********************** Tree implementation **********************//

func (trS *FolderTree) writeToFileJson(filePath string) {
	t, err1 := json.Marshal(trS.Folders)
	errCheck(err1)
	err2 := os.WriteFile(filePath, t, 0644)
	errCheck(err2)
}

func (trS *FolderTree) setVal(fldName string, file string) {
	trS.Mute.Lock()
	trS.Folders = append(trS.Folders,
		FolderTree{
			Name: fldName,
			File: file,
		})
	trS.Mute.Unlock()
}

func (trS *FolderTree) readDir(folderName string, file fs.FileInfo, wg *sync.WaitGroup) {
	atomic.AddInt64(&trS.FolderNum, 1)
	if atomic.LoadInt64(&trS.GoRtNum) < 128 {
		atomic.AddInt64(&trS.GoRtNum, 1)
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			defer atomic.AddInt64(&trS.GoRtNum, -1)
			trS.readDrive(folderName + "/" + dir)
		}(file.Name())
	} else {
		trS.readDrive(folderName + "/" + file.Name())
	}

}

func (trS *FolderTree) readDrive(folderName string) {
	wg := &sync.WaitGroup{}
	files, err := ioutil.ReadDir(folderName)
	if errCheck(err) {
		if len(Dir) < 1 {
			for _, file := range files {
				if file.IsDir() {
					trS.readDir(folderName, file, wg)
				} else {
					if len(File) < 1 {
						trS.setVal(folderName, file.Name())
					} else {
						contains := strings.Contains(file.Name(), File)
						if contains {
							trS.setVal(folderName, file.Name())
						}
					}
				}
			}
		} else {
			for _, file := range files {
				containsDir := strings.Contains(folderName, Dir)
				if file.IsDir() {
					trS.readDir(folderName, file, wg)
				} else {
					if len(File) < 1 && containsDir {
						trS.setVal(folderName, file.Name())
					} else {
						containsFile := strings.Contains(file.Name(), File)
						if containsFile && containsDir {
							trS.setVal(folderName, file.Name())
						}
					}
				}
			}
		}
		wg.Wait()
	}
}

//************************ Map implementation ************************//

// func (f *Folder) setVal(key, value string) {
// 	f.mu.Lock()
// 	f.fileList[key] = append(f.fileList[key], value)
// 	f.mu.Unlock()
// }

// func (f *Folder) setVals(key string, value []string) {
// 	f.mu.Lock()
// 	f.fileList[key] = value
// 	f.mu.Unlock()
// }

// func (f *Folder) readDir(folderName string, level int) {
// 	wg := &sync.WaitGroup{}
// 	files, err := ioutil.ReadDir(folderName)
// 	if errCheck(err) {
// 		for _, file := range files {
// 			if file.IsDir() {
// 				atomic.AddInt64(&f.dirCount, 1)
// 				if atomic.LoadInt64(&f.grNum) < 1024 {
// 					atomic.AddInt64(&f.grNum, 1)
// 					wg.Add(1)
// 					go func(dir string) {
// 						defer wg.Done()
// 						defer atomic.AddInt64(&f.grNum, -1)
// 						f.setVals(folderName+"/"+dir, []string{})
// 						f.readDir(folderName+"/"+dir, level+1)
// 					}(file.Name())
// 				} else {
// 					f.setVals(folderName+"/"+file.Name(), []string{})
// 					f.readDir(folderName+"/"+file.Name(), level+1)
// 				}
// 			} else {
// 				atomic.AddInt64(&f.folderCount, 1)
// 				f.setVal(folderName, file.Name())
// 			}
// 		}
// 	}
// 	if level == 0 {
// 		fmt.Println("Ready for finish")
// 	}
// 	wg.Wait()
// 	if level == 0 {
// 		fmt.Println("Dirs:", f.dirCount, "Files:", f.folderCount, "Level:", level, "GOroutines:", f.grNum, folderName)
// 	}
// }
