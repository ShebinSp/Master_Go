# Buffered Channel 
**a buffered channel is a type of channel that has a specified capacity to hold a certain number of values. This means it can store multiple values sent to it, allowing some flexibility in how goroutines communicate and synchronize.**

### Unbuffered Channel:

* No Capacity: It does not have any capacity to store values.
* No Storage: It cannot store any values. Each send operation must have a corresponding receive operation.
* Send Blocks: Sending a value into the channel will block (pause) until another goroutine receives that value.
* Receive Blocks: Receiving from the channel will block (pause) until another goroutine sends a value.