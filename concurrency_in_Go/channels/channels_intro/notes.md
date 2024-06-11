# Channel

* Channel is a more advanced concept in Go.

* A channel in Go provides a connection between two goroutines, allowing them to communicate.

* Channels are used to communicate in between running goroutines.

* The data we are sending through or receiving from a channel must be always of the same type. The type specified when we've declared the channel.

* `var ch chan int` - The value of uninitialized channel is `nil`.

* A channel is like a pointer, and passing channels to functions has the same effect as passing pointers to functions.

* `ch1 := make(chan int) - Channel declared and initialized.

* A channel is a two way messaging object. It has two principal operations `send` and `receive`.

* A `send` statement transmits a value from one goroutine through the channel to another goroutine executing a corresponding receive expression. Both operations are written using the channel operator( <- ).

* `ch <- 10` - sending 10 to the channel `ch`.

* `num := <- ch` - waiting a value to be sent into the channel, once a value received to channel, then the value is assigned to `num`.

* Another operation supported by the channel type is `close`. It sets a flag indicating that no more values will ever be sent on this channel. Subsequent attempts to send will panic.

* Receive operations on closed channel yield values that have been sent until no more values are left. Any receive operations thereafter yield the zero value of the channels element type.

* To close a channel call the built-in `close()` function.

* `close(ch)` - It closes the channel.