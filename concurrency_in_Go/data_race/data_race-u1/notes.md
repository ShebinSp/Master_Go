# Data Race

**Go language is known for how easy it is to build concurrent programs. Unfortunately, ensuring concurrent correctness requires the combination of many different techniques in order to minimise the chances of concurrency related errors.**

* What is Concurrency?

* Concurrency is when two goroutines are executing at the same time.

* Executing many goroutines at the same time without special handling can introduce a dreaded error called race condition or data race.

* Data races are among the most common and hardest to debug types of bugs in concurrent systems.

* A data race occurs when two goroutines are accessing memory at the same time, one of which is writing.

* Race conditions occur because of unsynchronized access to shared memory.