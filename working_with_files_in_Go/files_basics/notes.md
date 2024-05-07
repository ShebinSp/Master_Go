# Working With Files in Go

* The common way to work with files in Go is using `package os` . It provides a platform independent interface to operating system functionality.

* The design is Unix like, although the error handling is go like.

* The OS interface is intended to be uniform across all operating systems, means the program we create works the same on Windows, Linux or Mac.

* Other Go packages used to work with files which provides more functionality are,

   1. io

   2. bufio


* Package `os` is used to basic file operations

* `var newFile *os.File` - created a new file

* os.Create() creates a file if it doesn't already exist. If it exists, the file is truncated. This function returns a file descriptor, which is in fact a pointer to os.File and an error value. Other methods can be used on this file descriptor for input output.

* File Descriptor : **In simple words, when you open a file, the operating system creates an entry to represent that file and store the information about that opened file. So if there are 100 files opened in your OS then there will be 100 entries in OS (somewhere in kernel). These entries are represented by integers like (...100, 101, 102....). This entry number is the file descriptor. So it is just an integer number that uniquely represents an opened file for the process. If your process opens 10 files then your Process table will have 10 entries for file descriptors.**

* `os.Truncate(name string size int)` is used to truncate a file or empty a file.

The first arg is the name of the file and in Second arg 0 will completely empty the file, If 50 is the arg the file will be truncate to 50 bytes. If the file is over 50 bytes, everything past 50 bytes will be lost.

* `os.Open(name string)` is used to open a file.

* `os.OpenFile(name string, flag int, perm fs.FileMode)` is used to open a file with more operations.

* `os.Stat(name string)` is used to get file information and to check if a file exist.

* `os.Rename(name string)` is used to rename or move a file.

* `os.Remove(name string)` is used to remove a file.