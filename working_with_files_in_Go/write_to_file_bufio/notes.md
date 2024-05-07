# Writing to Files using a Buffered Writer(`bufio` Package)

* Another way to write to a file is to use `bufio` package. It lets you to create a buffered writer so you can write with a buffer in memory before writing it to disk.

* This is especially useful if you need to do a lot of manipulation on the data before writing to disk to save time from disk i/o.

* It is also useful if you only write one bite at a time and want to store a large number of bytes in memory before dumping it to a file at once. Otherwise you would be performing disk i/o for every byte. This is not efficient because the disk is slow in general.

**Notes:**

* Buffer : A buffer is temporary storage area in computer memory used to hold data while it's being transferred between two locations, typically between a fast source and a slower destination.

* Buffered Writer(`bufio.Writer`) : It's an abstraction provided by Go's `bufio` package. Instead of writing directly to the file every time you call `Write` , which can be slow because it involves interacting with the operating system frequently, the buffered writer gathers the data you want to write into a buffer, a temporary space in memory.

* Internal buffer : This is the temporary space, of buffer, that the buffered writer uses to hold the data before actually writing to the file. The internal buffer is a chunk of memory allocated by Go to temporarily store the data you write using the buffered writer.

*So, when you create a buffered writer using `bufio.NewWriter`, you are essentially creating an object that manages a buffer in memory. When you call `Write` on the buffered  writer, the data you pass to it gets stored in this buffer. When the buffer becomes full or when you call `Flush` on the buffered writer, the data in the buffer is written to the underlying file.*