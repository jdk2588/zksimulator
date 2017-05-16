package zksimulator

import ("sync")
type Message struct {
  sync.RWMutex
  client string
  key string
  value string
}

func (cm *Message) setClient(id string) {
    cm.Lock()
    defer cm.Unlock()
    cm.client = id
}

func (cm *Message) setKey(msg string) {
    cm.key = msg
}

func (cm *Message) setVal(msg string) {
    cm.value = msg + cm.client
}

func (cm *Message) getClient() string {
      return cm.client
}

func (cm *Message) getKey() string {
      return cm.key
}

func (cm *Message) getVal() string {
      return cm.value
}
