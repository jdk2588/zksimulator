package zksimulator

import (
  "sync"
  "strconv"
)

type Client struct {
  sync.RWMutex
  ident string
  deComission bool
  quorum int
  rack int
  region int
}

func (c *Client) deactivate() {
    c.Lock()
    defer c.Unlock()

    c.deComission = true
}

func (c *Client) isActive() bool {
    c.Lock()
    defer c.Unlock()
    return !c.deComission
}

func (c *Client) getRegion() int {
      return c.region
}

func (c *Client) setRegion(region int) {
      c.Lock()
      defer c.Unlock()
      c.region = region
}

func (c *Client) getRack() int {
      return c.rack
}

func (c *Client) setRack(rack int) {
      c.Lock()
      defer c.Unlock()
      c.rack = rack
}

func (c *Client) getIdent() string {
    return c.ident
}

func (c *Client) setIdent(identity int) {
    c.ident = "client" + strconv.Itoa(identity)
}

func oneClient(index int) *Client {

    clientC.Lock()
    defer clientC.Unlock()

    for i := 0;i < Env.NoOfClients; i++ {
      if i == index {
          clientC.clientInit[i].Do(func() {
            c := &Client{}
            c.setIdent(index)
            c.deComission = false
            assignRackandRegion(c, i)
            clientC.clientMap[index] = c
        })
      }
    }

	  return clientC.clientMap[index]
}

func GetClients() []*Client {

    clientList := []*Client{}
    for i:=0;i<Env.NoOfClients;i++ {
        c := oneClient(i)
        if !c.isActive() {
          continue
        }
        clientList = append(clientList, c)
    }
    return clientList
}
