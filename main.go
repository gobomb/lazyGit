package main

import (
	"github.com/fsnotify/fsnotify"
	"lazyGit/cmd"
	"log"
	"os/exec"
	"fmt"
	"time"
	"path/filepath"
	"os"
	"strings"
	"sync"
	"regexp"
)

var PATH string

type needPush struct {
	i  int
	is bool
	sync.Mutex
}

// build to win exe:
// env GOOS=windows GOARCH=amd64 go build
func main() {
	cmd.Execute(ExampleNewWatcher)
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func ExampleNewWatcher(path string) {
	np := needPush{
		is: false,
		i:  0,
	}
	i := 0
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)

				if regexp.MustCompile(`.git`).MatchString(event.Name) == true {
					continue
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}

				cmd := exec.Command("git", "add", ".")
				cmd.Dir = path
				err = cmd.Run()
				if err != nil {
					log.Println("git add:", err)
				}
				log.Printf("git add %d ok!", i)

				np.Lock()
				np.i = i
				np.is = true
				np.Unlock()

				i++
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	gitPushFunc := func() {
		np.Lock()
		if np.is == false {
			np.Unlock()
			return
		} else {

			commit := fmt.Sprintf("-m commit %d time at %s", i, time.Now().Format(time.Stamp))
			cmd := exec.Command("git", "commit", commit)
			cmd.Dir = path
			err = cmd.Run()
			if err != nil {
				log.Println("git commit", err)
			}
			log.Printf("git commit %d ok!", np.i)

			cmd = exec.Command("git", "push")
			cmd.Dir = path
			err = cmd.Run()
			if err != nil {
				log.Println("git push", err)
			}
			log.Printf("git push %d ok!", np.i)
			np.is = false
		}
		np.Unlock()
	}
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			gitPushFunc()
		}
	}
	<-done
	log.Println("done")

}
