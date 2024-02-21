package worker

type Job struct {
	id string
	Action func(map[string]string)
	Payload map[string]string
}

func (j Job) Do(sem chan struct{}) {
	sem <- struct{}{} // Acquire semaphore before starting the job
	go func() {
		defer func() { <-sem }() // Release semaphore after job is done
		j.Action(j.Payload)
	}()
}
