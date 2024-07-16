package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, job)
        time.Sleep(time.Second) // Simulate time-consuming work
        result := job * job     // Process the job
        fmt.Printf("Worker %d finished job %d\n", id, job)
        results <- result       // Send the result to the results channel
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    var wg sync.WaitGroup

    // Start workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Send jobs to the jobs channel
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Wait for all workers to finish
    wg.Wait()
    close(results)

    // Collect and print results
    for result := range results {
        fmt.Println("Result:", result)
    }
}
