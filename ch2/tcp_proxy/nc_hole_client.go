package main

import (
  "net"
  "log"
  "io"
  "os"
  "os/exec"
  "time"
  "fmt"
  //"bufio"
  //"bytes"
)

func handle(conn net.Conn) {
  rp, wp := io.Pipe()
  for {
    cmd := exec.Command("echo", "Hello World")
    cmd.Stdin = conn // so is this the server?
    cmd.Stdout = wp // commands stdout is not os.Stdout?

    go io.Copy(conn, rp)

    fmt.Println("ABOUT TO RUN the command")
    // if err := cmd.Run(); err != nil {
      // log.Fatalln("COMMAND ERROR: " + err.Error())
      // break
    // }
    log.Printf("Running command and waiting for it to finish...")
    err := cmd.Run()
    log.Printf("Command finished with error: %v", err)
    time.Sleep(1*time.Second)
  }
}

func main() {
  conn, err := net.Dial("tcp", "localhost:20080")
  if err != nil {
    log.Fatalln("CONNECTION ERROR CLIENT: " + err.Error())
  }
  defer conn.Close()

  go func(c net.Conn) {
    _, err := io.Copy(os.Stdout, c) // I think this blocks cuz connection doesnt close which is supposed to be server?
    // not sure but this blocks from anymore sends to server
    if err != nil {
     log.Fatal(err)
    }

     //input := bufio.NewScanner(c)
     //for input.Scan() {
     //fmt.Println("TEXT: " + input.Text())
   //}
  }(conn)

  handle(conn)
}
