package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"math/rand"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// DB is the main Database
var DB *gorm.DB

var testAlarm *alarm

// If true the music should be stopped
var musicHelper = false

const (
	musicRoot string = "/home/maxibanki/music/syndicate/"
)

func main() {
	alarmTime, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if err != nil {
		log.Fatal(err)
	}
	testAlarm = &alarm{
		Name:      "Main Alarm",
		AlarmTime: alarmTime,
		Active:    true,
		Repeat:    86400,
	}
	go func() {
		startHTTPServer()
	}()
	defer DB.Close()
	// a channel to tell it to stop
	stopchan := make(chan struct{})
	// a channel to signal that it's stopped
	stoppedchan := make(chan struct{})
	go func() {
		// close the stoppedchan when this func
		// exits
		defer close(stoppedchan)
		defer func() {
			fmt.Println("Teardown!")
		}()
		for {
			select {
			default:
				if time.Now().Sub(testAlarm.AlarmTime).Seconds() > 0 {
					power(true)
					_, err := exec.Command("mpg123", "play", musicRoot+getRandomMusic()).Output()
					if err != nil {
						fmt.Println(err)
					}
				}
			case <-stopchan:
				fmt.Println("stopping!")
				return
			}
		}
	}()
	go func() {
		for {
			if musicHelper {
				musicHelper = false
				close(stopchan)
				<-stoppedchan
			}
			time.Sleep(1 * time.Second)
		}

	}()
	for {
		time.Sleep(1 * time.Second)
	}
}

func power(state bool) error {
	stateAsString := "0"
	if state {
		stateAsString = "1"
	}
	_, err := exec.Command("sh", "-c", "/opt/raspberry-remote/send 10101 4 "+stateAsString).Output()
	if err != nil {
		return err
	}
	return nil
}

func numRand(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func getRandomMusic() string {
	files, _ := ioutil.ReadDir(musicRoot)
	fileKey := numRand(0, len(files))
	var i int
	var name string
	filepath.Walk(musicRoot, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if i == fileKey {
				name = info.Name()
			}
			i++
		}
		return nil
	})
	return name
}

// func createAlarm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	var a alarm
// 	if r.Body == nil {
// 		http.Error(w, "Please send a request body", 400)
// 		return
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&a)
// 	if err != nil {
// 		http.Error(w, err.Error(), 400)
// 		return
// 	}
// 	fmt.Println(a.Name)
// 	DB.NewRecord(&a)
// }

func testingFnc(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	musicHelper = true
}

func startHTTPServer() {

	router := httprouter.New()
	// router.GET("/api/v1/power/:state", power)
	// router.POST("/api/v1/createAlarm", createAlarm)
	router.GET("/api/v1/test", testingFnc)
	log.Println("Successfully started the server on Port 80")
	log.Fatal(http.ListenAndServe(":80", router))

}

func openDatabase() {
	var err error
	DB, err = gorm.Open("sqlite3", "main.db")
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&alarm{})
}

// StdOut is the basic structure for printing
type StdOut struct {
	Success bool
	Data    string
}

type alarm struct {
	gorm.Model
	Name      string
	AlarmTime time.Time
	Active    bool
	// In seconds
	Repeat time.Duration
}
