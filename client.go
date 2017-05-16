package zksimulator

import (
 "sync"
)

func (c *Client) process(zkconn *ZkConn, msg string, wg *sync.WaitGroup) {

  defer wg.Done()

  m := &Message{client: c.getIdent()}
  m.setKey(msg)
  m.setVal(msg)

	zl := GetLock(zkconn, m)
  e := zl.Lock()
  if e != nil {
    noLockgiven("No lock given", msg, c)
  } else {
    successLog("Lock acquired by", msg, c)
  }
}
