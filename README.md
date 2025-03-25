# 🔄 Go Load Balancer with Worker Pool

This project demonstrates a **simple load balancer** pattern in Go using goroutines, channels, and the `select` statement. It simulates distributing jobs to a pool of workers, each of which processes jobs concurrently.

## 🚀 Features

- Fixed number of workers
- Central dispatcher sending jobs
- Workers process jobs with simulated delays
- Uses `select` and channels for concurrency
- Graceful shutdown using `sync.WaitGroup`

## 🧠 How It Works

1. A set number of **workers** are started as goroutines, each listening to a shared `jobs` channel.
2. A **dispatcher goroutine** sends jobs into the channel at a fixed interval.
3. Each worker reads from the channel, processes the job, and logs its output.
4. After all jobs are dispatched, the channel is closed, and workers exit gracefully.

## 📦 Running the Example

### Prerequisites

- Go 1.18+

### Run It

```bash
go run main.go
```

### Expected Output
```bash
Dispatcher: Sent job 1
Worker 1: Processed job 1 in 422ms
Dispatcher: Sent job 2
Worker 2: Processed job 2 in 383ms
...
All jobs processed.
```


