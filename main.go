package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	urls := []string{"https://trackerslist.com/best.txt",
		"https://ngosang.github.io/trackerslist/trackers_best.txt"}
	responseData := make(chan []string)
	var wg sync.WaitGroup

	wg.Add(len(urls))
	for _, url := range urls {
		go returnURLs(url, &wg, responseData)
	}
	set := make(map[string]bool)

	go func() {
		for response := range responseData {
			for _, v := range response {
				if len(v) != 0 && set[v] == false {
					set[v] = true
				}
			}
		}
	}()
	wg.Wait()
	// writeTheFile()
	output, err := exec.Command("C:\\Program Files\\qBittorrent\\qbittorrent.exe").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}
func returnURLs(url string, wg *sync.WaitGroup, responseData chan []string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
		return
	}
	str := string(body)
	responseData <- strings.Split(str, "\n")
}

// func writeTheFile() {
// 	file, err := os.OpenFile("C:\\Users\\kusha\\Desktop\\data.txt", os.O_APPEND, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}os.
// 	if _, err := file.Write([]byte("appended some data\n")); err != nil {
// 		file.Close() // ignore error; Write error takes precedence
// 		log.Fatal(err)
// 	}
// 	if err := file.Close(); err != nil {
// 		log.Fatal(err)
// 	}
// }
