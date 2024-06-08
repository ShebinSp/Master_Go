# Mutex

* There are more approaches to the data race issue such as mutexes and channels.

* Mutex comes from mutual exclusion. This is called explicit synchronization, where variables access is protected through synchronization primitives such as a mutex.

* Explicit synchronization means that the programmer recognises candidates for concurrent execution and the contexts in which they will be executed.

* `var m sync.Mutex`. There are two methods defined on mutex type, `lock` and `unlock`. Any code that is present between a call to `lock` and `unlock` will be executed by only one goroutine. This way race conditions are avoided.

* The `lock` method of the mutex variable blocks the access to the variable until the `unlock` method is called.

* If one goroutine already holds the lock and if a new goroutine is trying to acquire the lock, the new goroutine will be blocked until the mutex is unlocked.