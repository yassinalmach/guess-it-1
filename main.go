// Complete Guide to Go's bufio Package

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // This guide demonstrates the key features of bufio package
    
    // 1. READER OPERATIONS
    // ==================
    
    // Creating a new buffered reader
    str := "Hello\nWorld\nGo Programming"
    reader := bufio.NewReader(strings.NewReader(str))
    
    // ReadString - reads until delimiter
    line1, _ := reader.ReadString('\n')
    fmt.Printf("ReadString result: %q\n", line1) // Output: "Hello\n"
    
    // ReadLine - low-level line reading
    line2, isPrefix, _ := reader.ReadLine()
    fmt.Printf("ReadLine result: %q, prefix: %v\n", line2, isPrefix)
    
    // Scanner Operations
    // =================
    scanner := bufio.NewScanner(strings.NewReader(str))
    
    // Default scanner is line by line
    for scanner.Scan() {
        fmt.Printf("Scanned line: %s\n", scanner.Text())
    }
    
    // Custom split function
    scanner = bufio.NewScanner(strings.NewReader("Hello World Go"))
    scanner.Split(bufio.ScanWords) // Split by words
    
    for scanner.Scan() {
        fmt.Printf("Scanned word: %s\n", scanner.Text())
    }
    
    // 2. WRITER OPERATIONS
    // ==================
    
    // Creating a buffered writer
    writer := bufio.NewWriter(os.Stdout)
    
    // Writing string
    writer.WriteString("Buffered ")
    writer.WriteString("output\n")
    
    // Writing single byte
    writer.WriteByte('!')
    
    // Writing rune
    writer.WriteRune('世')
    
    // Flush to ensure all buffered operations are applied
    writer.Flush()
    
    // 3. READWRITER
    // ============
    
    // Creating a ReadWriter - combines reader and writer capabilities
    readWriter := bufio.NewReadWriter(
        bufio.NewReader(strings.NewReader("Input text")),
        bufio.NewWriter(os.Stdout),
    )
    
    // Using ReadWriter methods
    text, _ := readWriter.ReadString('\n')
    readWriter.WriteString("\nRead from ReadWriter: " + text)
    readWriter.Flush()
}

// Key Features and Constants:
// =========================
/*
1. Reader Methods:
   - ReadBytes(delimiter) - reads until delimiter
   - ReadString(delimiter) - reads until delimiter as string
   - ReadLine() - low-level line reading
   - ReadSlice(delimiter) - reads until delimiter, returns slice
   - ReadRune() - reads single UTF-8 encoded rune
   - UnreadRune() - unreads last rune
   - UnreadByte() - unreads last byte
   - Peek(n) - returns next n bytes without advancing reader
   
2. Writer Methods:
   - WriteString(s) - writes string
   - WriteByte(c) - writes single byte
   - WriteRune(r) - writes UTF-8 encoded rune
   - Flush() - writes buffered data to underlying io.Writer
   - Available() - returns how many bytes are available
   - Buffered() - returns number of bytes in buffer
   
3. Scanner Features:
   - Split functions:
     * ScanLines (default) - split by lines
     * ScanWords - split by words
     * ScanBytes - split by bytes
     * ScanRunes - split by runes
   - Custom split functions possible
   
4. Important Constants:
   - defaultBufSize = 4096 // Default buffer size
   - MaxScanTokenSize = 64 * 1024 // Maximum token size for scanner
*/

// Example: Custom Scanner Split Function
func ExampleCustomSplit() {
    // Creating custom split function for comma-separated values
    scanner := bufio.NewScanner(strings.NewReader("apple,banana,cherry"))
    scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        for i := 0; i < len(data); i++ {
            if data[i] == ',' {
                return i + 1, data[:i], nil
            }
        }
        if atEOF {
            return len(data), data, nil
        }
        return 0, nil, nil
    })
    
    for scanner.Scan() {
        fmt.Printf("CSV item: %s\n", scanner.Text())
    }
}

// Example: Efficient File Reading
func ExampleFileReading() {
    file, err := os.Open("example.txt")
    if err != nil {
        return
    }
    defer file.Close()
    
    reader := bufio.NewReader(file)
    buffer := make([]byte, 1024)
    
    for {
        n, err := reader.Read(buffer)
        if err != nil {
            break
        }
        // Process buffer[:n]
        _ = n // Used in real code
    }
}

// Example: Writing with Size Control
func ExampleSizedWriter() {
    // Create a writer with specific buffer size
    writer := bufio.NewWriterSize(os.Stdout, 8192)
    
    // Write data
    writer.WriteString("Large amount of data")
    
    // Check buffer status
    fmt.Printf("Bytes available: %d\n", writer.Available())
    fmt.Printf("Bytes buffered: %d\n", writer.Buffered())
    
    writer.Flush()
}