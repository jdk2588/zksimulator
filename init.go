package zksimulator

import (
  "sync"
  "zksimulator/config"
)

var clientC = &ClientCaches{}
var Env config.Settings

type ClientCaches struct {
  sync.RWMutex
  clientMap  map[int]*Client
  clientInit []sync.Once
}

func Init() {
  clientC.clientMap = map[int]*Client{}
  clientC.clientInit = make([]sync.Once, Env.NoOfClients)
}
