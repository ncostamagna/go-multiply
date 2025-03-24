Concurrency is about MANAGING multiple task at once, parallelism is about EXECUTING multiple tasks at once

# Concurrency

- Program that can handler multiple task.
- Performing many tasks in a single CPU Core, create the illusion that tasks are progressing simultaneously, but really not

# Parallelism

- Simultaniean executions.
- Multiple task can be processed simultained, using multiples CPU cores.

## Capacity
```go 
make(chan *string, 20) // capatiyy of 20
```

affter the channel is closed we cant send anyrhing to the channel but we can recive values


https://www.youtube.com/watch?v=RlM9AfWf1WU&ab_channel=ByteByteGo

# Fork-Join Model

The Fork-Join model is a parallel programming pattern used to break a big task into smaller subtasks that can be done concurrently, and then combined to get the final result.

It’s widely used in divide-and-conquer algorithms, multithreading, and parallel processing.

## Fork
You split (fork) the main task into smaller, independent subtasks, and run them in parallel (e.g., using threads, goroutines, or processes).

## Join
You wait (join) until all the subtasks are done, then combine their results to produce the final output.

## Pattern
It's a pattern commonly used in parallel processing, like in divide-and-conquer algorithms (e.g. parallel mergesort, image processing, etc.)
<br />
<img src="images/001.webp">

## 🔍 Fork-Join Model vs Fork-Join Pool

- Can be mimicked in Go using goroutines + channels + worker pools


| Concept      | Fork-Join **Model**             | Fork-Join **Pool**                      |
|--------------|----------------------------------|------------------------------------------|
| What it is   | A pattern or strategy            | A system to implement the pattern        |
| Level        | High-level idea                  | Practical tool/framework                 |
| Example      | "Split task, run subtasks, join" | Java's `ForkJoinPool` class              |
| Usage        | Theory or any language           | Java or other specific platforms         |

---


- **Fork-Join model** is the recipe  
- **Fork-Join pool** is the kitchen appliance that makes it easier to cook that recipe

# Pool

sync.Pool in Go is a structure for efficiently reusing objects and reducing the overhead of repeatedly creating and destroying them.

- It's ideal when you need to create many temporary objects (like buffers, resources, etc.).

- Internally, sync.Pool maintains a set of ready-to-use objects.

- When you call Get(), it tries to give you an existing object.

    - If the pool is empty, it calls the New() function to create a new one.

- When you're done with an object, you call Put() to return it to the pool.

## What is sync.Pool
To put it simply, sync.Pool in Go is a place where you can keep temporary objects for later reuse.
<br />
But here’s the thing, you don’t control how many objects stay in the pool, and anything you put in there can be removed at any time, without any warning and you’ll know why when reading last section.
<br />
The good point is, the pool is built to be thread-safe, so multiple goroutines can tap into it simultaneously. Not a big surprise, considering it’s part of the sync package.
<br />
https://victoriametrics.com/blog/go-sync-pool/