## Difference between Buffered and Unbuffered Channels

***The primary difference between buffered and unbuffered channels in Go lies in how they handle the storage of values and the synchronization of goroutines. Here’s a detailed comparison:***

### Buffered Channels
1. Capacity:

* Has a Capacity: Buffered channels have a specified capacity, which determines how many values they can hold.
* Example: ch := make(chan int, 3) creates a buffered channel with a capacity of 3.

2. Blocking Behavior:

* Send Without Blocking: Sending a value to a buffered channel does not block the sender if the buffer is not full. It only blocks if the buffer is full.
* Receive Without Blocking: Receiving a value from a buffered channel does not block the receiver if the buffer is not empty. It only blocks if the buffer is empty.

3. Use Case:

* Useful when you want to allow some degree of decoupling between the sender and receiver, providing a way to handle bursts of data without immediate synchronization.


### Unbuffered Channels
1. Capacity:

* No Capacity: Unbuffered channels have no capacity to store values. They do not hold any values and can only transfer a single value directly from sender to receiver.
* Example: ch := make(chan int) creates an unbuffered channel.

2. Blocking Behavior:

* Send Blocks: Sending a value to an unbuffered channel blocks the sender until another goroutine receives the value.

* Receive Blocks: Receiving a value from an unbuffered channel blocks the receiver until another goroutine sends a value.

3. Use Case:

* Useful for strict synchronization between goroutines, ensuring that the sender and receiver are synchronized at the same point in execution.


## Comparison of Buffered and Unbuffered Channels

| Feature              | Buffered Channel                                | Unbuffered Channel                                 |
|----------------------|-------------------------------------------------|----------------------------------------------------|
| **Capacity**         | Has a specified capacity                        | No capacity (size of 0)                            |
| **Send Behavior**    | Sends without blocking until the buffer is full | Sends block until a receiver is ready              |
| **Receive Behavior** | Receives without blocking until the buffer is empty | Receives block until a sender is ready             |
| **Synchronization**  | Provides partial synchronization                | Provides full synchronization                      |
| **Use Case**         | When you need a queue-like structure to handle bursts of data | When you need strict synchronization between sender and receiver |

* **BUFFERED CHANNELS OR UNBUFFERED CHANNELS?**
* The choice between unbuffered and buffered channels, and the choice of a buffered channel capacity may both effect
ho a concurrently program works.
Unbuffered channels give stronger synchronization guarantees because every send operation is synchronized with its
corresponding receive.

* With buffered channels, these operations are decoupled. Also when we know an upper bound on the number of values
that will send on a channel, it's usual to create a buffered channel of that size and perform all the send operations
before the first value is received.

* Channel buffering may also affect program performance. There is a metaphor and i find it very suggestive. ` Imagine
three cooks in a cake shop, one backing, one icing and one inscribing each cake before passing it on the next cook in
the assembly line in a kitchen with little space. Each cook that has finished a cake must wait for the next cook to 
become ready to accept it, this is analog to communication over an unbuffered channel. If there is space for on cake
between each cook, a cook may place a finished cake there and immediately start work on next, this is an analog to a
buffered channel with a capacity of 1. So as long as the cooks work at about the same rate on average, most of these
handover proceed quickly and smoothly.
    On the other hand, if an earlier stage of the assembly line is consistently faster than the following stage, the
buffer between them will spend most of it's time full. A buffer provides no benefits in this case.