package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var taskRunning int32

func runTask(){
	atomic.StoreInt32(&taskRunning, 1)
	MatchAngMergePapers()
	atomic.StoreInt32(&taskRunning, 0)
}

func HelpHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Use /start to start match-merge process, use /check to check if process is running")
}

func StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	tr := atomic.LoadInt32(&taskRunning)
	if tr == 0 {
		go runTask()
		fmt.Fprintln(w, "Task run!")
		return
	}

	fmt.Fprintln(w, "Task is already running")
}

func IsTaskRunningHandler(w http.ResponseWriter, r *http.Request) {
	tr := atomic.LoadInt32(&taskRunning)

	if tr == 0 {
		fmt.Fprintln(w, "Task isn't running")
		return
	}

	fmt.Fprintln(w, "Task is running now")
}

func main() {
	http.HandleFunc("/start", StartTaskHandler)
	http.HandleFunc("/check", IsTaskRunningHandler)
	http.HandleFunc("/", HelpHandler)

	err := http.ListenAndServe(":9303", nil)
	if err != nil {
		panic(err)
	}
}