package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/vacmar01/simpleQ/worker"
	"flag"
	"os"
)

var (
	script string
	queue string
	maxJobs int
)

func main() {
	flag.StringVar(&script, "script", "", "The script to run")
	flag.StringVar(&queue, "queue", "", "The queue to watch")
	flag.IntVar(&maxJobs, "maxJobs", 5, "The maximum number of jobs to run")
	flag.Parse()

	// Check if the script exists
	if _, err := os.Stat(script); os.IsNotExist(err) {
		fmt.Println("Script does not exist")
		return
	}

	// Check if the queue exists
	if _, err := os.Stat(queue); os.IsNotExist(err) {
		fmt.Println("Queue file does not exist")
		return
	}

	t, err := tail.TailFile(queue, tail.Config{Follow: true})
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}	

	sem := make(chan struct{}, maxJobs)
	fmt.Println("Watching queue", queue, "for new tasks with a maximum of", maxJobs, "jobs running at once")
	
	for line := range t.Lines {
		job := worker.Job{Action: runScript, Payload: map[string]string{"script": script, "line": line.Text}}
		job.Do(sem)
	}
}