package main

import (
  "fmt"
  "net"
  //"sync"
  "sort"
)

//func worker(ports chan int, wg *sync.WaitGroup) {
//  for p := range ports {
//    fmt.Println(p)
//    wg.Done()
//  }
//}

//func main() {
//  ports := make(chan int, 100)
//  var wg sync.WaitGroup
//  for i := 0; i < cap(ports); i++ {
//    go worker(ports, &wg)
//  }
//  for i := 1; i <= 1024; i++ {
//    wg.Add(1)
//    ports <- i
//  }
//  wg.Wait()
//  close(ports)
//}

func worker(ports, results chan int) { // both same type so can omit first type declaration. also material reminded me of receive only or send only channels. Review topics!
  for p := range ports {
    address := fmt.Sprintf("scanme.nmap.org:%d", p)
    conn, err := net.Dial("tcp", address)
    if err != nil {
      results <- 0
      continue
    }
    conn.Close()
    results <- p
  }
}

func main() {
  ports := make(chan int, 100)
  results := make(chan int)
  var openports []int

  for i := 0; i < cap(ports); i++ {
    go worker(ports, results)
  }

  go func() { // must do loop below in separate go routine because results channel will block after more than 100 scans because the results channel isnt buffered and receiving of results channel is done in loop below
    //for i := 1; i <= 1024; i++ {
    for i := 1; i <= 65535; i++ {
      ports <- i
    }
  }()

  for i := 0; i < 65535; i++ {
    port := <-results
    if port != 0 {
      openports = append(openports, port)
    }
  }

  close(ports)
  close(results)
  sort.Ints(openports)
  for _, port := range openports {
    fmt.Printf("%d open\n", port)
  }
}
