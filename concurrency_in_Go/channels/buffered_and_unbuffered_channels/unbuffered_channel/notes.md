# Unbuffered Channel

* There are two types of channels that have different behaviour Buffered and Unbuffered channels.

* A channel created with a simple call to make() built-in function is called an unbuffered channel.

* `var unbufferedChannel = make(chan int)`

* The `make()` function accepts an optional second argument of type `int` called the channel capacity. If the capacity is non-zero, `make()` creates a buffered channel.
    `var bufferedChannel = make(chan int, 3)`

* An unbuffered channel is a channel with no capacity to store any values. It ensures that any send operation (ch <- value) is synchronized with a corresponding receive operation (value := <- ch). In other words, the send and receive operations on an unbuffered channel happen simultaneously, allowing for direct handoff of data between goroutines.

**Characteristics of Unbuffered Channels**
1. Blocking Behavior:

* Sending: When a goroutine attempts to send a value to an unbuffered channel, it blocks until another goroutine is ready to receive from that channel.
* Receiving: Similarly, when a goroutine attempts to receive a value from an unbuffered channel, it blocks until another goroutine is ready to send a value to that channel.

2. Synchronization:

* Unbuffered channels provide a way to synchronize execution between goroutines. They ensure that a send operation completes only when a corresponding receive operation occurs, and vice versa.