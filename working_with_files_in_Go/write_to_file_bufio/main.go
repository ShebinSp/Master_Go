package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// WRITING TO FILES USING A BUFFERED WRITER(`bufio` PACKAGE)

	// Another way to write to a file is to use `bufio` package.It lets you to create a-
	// buffered writer so you can write with a buffer in memory before writing it to disk.

	// the 1st step is to open the file for writing using `os.OpenFile()`
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// the next step is to create a buffered writer from the file variable using NewWriter() function from buffer.
	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}

	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n",bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n",bytesAvailable)
	
	// string contains ASCII characters, each character or rune is 1 byte.
	bytesWritten, err = bufferedWriter.WriteString("\nThis string also added to the bufferWriter later")
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n",unflushedBufferSize)

	bytesAvailable = bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n",bytesAvailable) // 4096 - 52 = 4044 bytes

	// The bytes have been written to buffer, not yet to the file. To do so , we use the-
	// flush method of the bufferedWriter object.
	// Flushing the buffer means ensuring that all the data in the buffer is written out to the underlying file
	bufferedWriter.Flush()

	// If you want to revert any changes done to buffer that have not been written yet to file with Flush(),
	// the `reset` method can be used. In this case buffer has been flushed to file, so ther is no reason to reset it.
}
