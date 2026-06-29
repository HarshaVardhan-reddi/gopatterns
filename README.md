# Go Concurrency Patterns

Practical examples demonstrating Go concurrency patterns using goroutines, channels, and sync primitives.

## Patterns

### 1. Three-Stage Pipeline (`3stagepipeline/`)

A multi-stage pipeline where data flows through three stages connected by channels:

- **Generator** — produces a stream of integers (0 to 10,000)
- **Modifier** — squares each number
- **Filter** — passes through only even numbers

Uses `context.Context` for cancellation and buffered channels between stages.

```
Generator → Modifier (n²) → Filter (even only) → Print
```

### 2. N-Worker Pool (`nworkers/`)

A fixed pool of 10 worker goroutines processing 1,000,000 jobs from a shared channel. Demonstrates the fan-out pattern where multiple goroutines consume from a single channel.

### 3. Bounded Concurrent Requests (`concurrentrequests/`)

Fires off 50 HTTP requests concurrently with a bounded concurrency of 10 using a semaphore channel. Each request has a 30-second timeout via `context.WithTimeout`.

## Running

Each pattern is a standalone Go module:

```bash
cd 3stagepipeline && go run .
cd nworkers && go run .
cd concurrentrequests && go run .
```
