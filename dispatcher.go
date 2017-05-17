package zksimulator

import (
  "sync"
)

func Dispatcher(ch chan string, zkconn *ZkConn, wg *sync.WaitGroup, clients []*Client) {
  defer wg.Done()

  message := <-ch

  simulate(message, clients, zkconn)

}

func simulate(message string, clients []*Client, zkconn *ZkConn) {

  var wg1 sync.WaitGroup
  for _ , c := range(clients) {
    wg1.Add(1)
    go c.process(zkconn, message, &wg1)
  }

  wg1.Wait()
}
