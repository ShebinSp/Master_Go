# Select Statement

* `select` is only used with channels.

* The select statement lets a goroutine wait on multiple communication operations. A select blocks until one of its cases can run then it executes that case. It chooses on at random if multiple are ready.

* This is a very powerful feature of Go because it allows us to wait on multiple channel operations.

* Send or receive operations on a nil channel block forever. This can be used to disable a channel in a select statement.

----------------------------------------------------------------------------------

* In Go, the select statement is used to handle multiple channel operations. It allows a goroutine to wait on multiple communication operations, making it possible to work with several channels simultaneously. The select statement chooses one of the ready channels to proceed with, enabling non-blocking communication and more complex synchronization patterns.

### Key Points about the select Statement:

1. Multiple Cases: You can specify multiple case blocks, each involving a channel operation (send or receive).

2. Default Case: An optional default case can be included to handle the situation when none of the channels are ready. This prevents the select statement from blocking.

3. Blocking Behavior: The select statement blocks until at least one of its cases can proceed. If multiple cases are ready, one is chosen at random.

4. Non-blocking Behavior: With a default case, the select statement does not block and executes the default case if no other case is ready.