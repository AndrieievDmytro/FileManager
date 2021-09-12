package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"sync"
// )

// var (
// 	subReaders     int
// 	waitSubReaders chan bool
// )

// func main() {
// 	subReaders = 0
// 	waitSubReaders = make(chan bool)
// 	fileList := readDir("c:/", true)
// 	dirsNum := 0
// 	filesNum := 0
// 	for k, v := range fileList {
// 		fmt.Println(k)
// 		fmt.Print("\t")
// 		fmt.Println(v)
// 		dirsNum++
// 		filesNum += len(v)
// 	}
// 	fmt.Println("Dirs:", dirsNum, "\t Files:", filesNum)
// }

// func errCheck(err error) bool {
// 	result := true
// 	if err != nil {
// 		// log.Fatal(err)
// 		fmt.Println(err)
// 		result = false
// 	}
// 	return result
// }

// func readDir(folderName string, firstLevel bool) map[string][]string {
// 	fileList := make(map[string][]string)
// 	files, err := ioutil.ReadDir(folderName)
// 	localSubReaders := 0
// 	waitLocalSubReaders := make(chan bool)
// 	var mutex = &sync.Mutex{}
// 	if errCheck(err) {
// 		for _, file := range files {
// 			if file.IsDir() {
// 				dir := file
// 				if subReaders > 128 {
// 					subFolder := readDir(folderName+"/"+dir.Name(), false)
// 					for k, v := range subFolder {
// 						mutex.Lock()
// 						fileList[k] = v
// 						mutex.Unlock()
// 					}
// 				} else {
// 					subReaders++
// 					localSubReaders++
// 					go func() {
// 						subFolder := readDir(folderName+"/"+dir.Name(), false)
// 						for k, v := range subFolder {
// 							mutex.Lock()
// 							fileList[k] = v
// 							mutex.Unlock()
// 						}
// 						subReaders--
// 						localSubReaders--
// 						if localSubReaders == 0 {
// 							waitLocalSubReaders <- true
// 						}
// 						if subReaders == 0 {
// 							waitSubReaders <- true
// 						}
// 					}()
// 				}
// 			} else {
// 				mutex.Lock()
// 				fileList[folderName] = append(fileList[folderName], file.Name())
// 				mutex.Unlock()
// 			}
// 		}
// 	}
// 	for localSubReaders > 0 {
// 		<-waitLocalSubReaders
// 	}
// 	for firstLevel && subReaders > 0 {
// 		<-waitSubReaders
// 	}
// 	// mutex.Lock()
// 	return fileList
// }
