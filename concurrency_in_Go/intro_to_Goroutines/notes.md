# CONCURRENCY and PARALLELISM in GO

* **Go History:**

   * Go is the first programming language built after the first dual core CPU was released.

   * Intel released its first duel core CPU in 2006, Google started developing Go in 2007 and publicly announced the language in November 2009 on Google OpenSource Blog.

   * Version 1.0 of Go was released in March 2012.

* Go is the first major programming language that has concurrency built-in.

* Go makes concurrency and parallelism easy and efficient.

## CONCURRENCY

* Concurrency means loading more goroutines at a time. These goroutines are multiple threads of execution. If one goroutine blocks, another one is picked up and started. On single core CPU  you can run ONLY concurrent applications but ther are not run in parallel.

* Imagine you are downloading some files while listen to music on a web page that you are scrolling at the same time. If there is only one CPU or a single core CPU, the operation system divides CPU time among things that need to be processed and we get a sensation of multiple things happening at the same time. But in reality only one thing is happening at a time. This is concurrency.

## PARALLELISM

* Parallelism means multiple goroutines executed at the same time. It requires multiple CPU's.

***Concurrency means independently executing processes or dealing with multiple things at once, while parallelism is the simultaneos execution of process and require multiple core CPU's.***

# Goroutines

* A goroutine is a lightweight thread of execution; **goroutines are key ingredients to achieve concurrency in Go.**

1. When we compile and run a Go program, we create a process which contains inside a goroutine, the `main` go routine.

2. The goroutine executes the instructions of the program one by one.

* A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword `go` followed by a function invocation.

* Goroutines are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes a fixed size of 1-2 Mb. The stack is a memory region where threads and goroutines store data they are working with. In general, this consists of local variables of function calls that are in progress or temporarily suspended while other function is called.

* An OS Thread Stack is fixed size but a goroutine stack size shrinks and grow as needed.

* **Scheduling a goroutine is much cheaper than scheduling a thread.** It is practical to have thousands, even hundreds of thousands of goroutines

* OS threads are scheduled by the OS kernel, but goroutines are scheduled by its own Go Scheduler using a technique called `m:n scheduling`, because it multiplexes (or schedules) **m goroutines on n OS threads.**

* Goroutines have no identity. There is no notation of identity that is accessible to the programmer.