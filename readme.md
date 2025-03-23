Concurrency is about MANAGING multiple task at once, parallelism is about EXECUTING multiple tasks at once

## Concurrency
Program that can handler multiple task.
Performing many tasks in a single CPU Core, create the illusion that tasks are progressing simultaneously, but really not

## Parallelism
Simultaniean executions.
Multiple task can be processed simultained, using multiples CPU cores.

## Capacity
```go 
make(chan *string, 20) // capatiyy of 20
```

affter the channel is closed we cant send anyrhing to the channel but we can recive values


https://www.youtube.com/watch?v=RlM9AfWf1WU&ab_channel=ByteByteGo