package main

import (
  "log"
  "net"
  "bufio"
 // "io"
 "fmt"
 "bytes"
)

func echo(conn net.Conn) {
  defer conn.Close()

  // approach 0

  // if _, err := io.Copy(conn, conn); err != nil {
    // fmt.Println("IO COPY ERROR")
    // log.Fatalln("Unable to read/write data")
  // }
  // approach 1- this doesnt echo, but above does. Have that intermiediate buffer doesnt work for some reason?
  // for {
    // var b bytes.Buffer
    // if _, err := io.Copy(&b, conn); err != nil { // this blocks for some reason
      // log.Fatalln("Unable to read/write data from conn to buffer")
      // // continue
      // conn.Close()
      // return
    // }
    // log.Printf("Read %d bytes: %s\n", len(b.Bytes()), b.String())

    // if _, err := io.Copy(conn, &b); err != nil {
      // fmt.Println("IO COPY ERROR")
      // log.Fatalln("Unable to read/write data")
    // }
  // }

  //  approach 2
  // input := bufio.NewScanner(conn)
  // for input.Scan() {
    // var b bytes.Buffer
    // b.WriteString("From Server: " + input.Text())
    // b_string := b.String()
    // fmt.Println("About to print scanne text in server.")
    // fmt.Println(b_string)
    // fmt.Fprintln(conn, b_string)
  // }

  // approach 3
  reader := bufio.NewReader(conn)
  writer := bufio.NewWriter(conn)
  for {
    // client will eventually stop sending ?
    s, err := reader.ReadString('\n')
    if err != nil {
      log.Fatalln("Unable to read data")
    }
    log.Printf("Read %d bytes: %s", len(s), s)

    log.Println("Writing data")
    var b bytes.Buffer
    b.WriteString(s + ". Hi back from server!")
    //if _, err := writer.WriteString(s + ". Hi back from server!"); err != nil {
    if _, err := writer.WriteString(b.String()); err != nil {
      log.Fatalln("Unable to write data")
    }
    writer.Flush()
  }
}

func main() {
  // Bind to TCP port 20080 on all interfaces.
  listener, err := net.Listen("tcp", ":20080")
  fmt.Println("buf_echo listenint on port 20080")
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
    go echo(conn)
  }
}
