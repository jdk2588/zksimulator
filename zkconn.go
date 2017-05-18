package zksimulator

import (
  "time"
	"github.com/samuel/go-zookeeper/zk"
)

type ZkConn struct {
    c *zk.Conn
}

func Connect(hosts []string, timeOut time.Duration) *ZkConn {
	conn, _, _ := zk.Connect(hosts, timeOut)
	zc := &ZkConn{}
  zc.c = conn
  return zc
}

func Disconnect(zc *ZkConn) {
    zc.c.Close()
}

type ZkLock struct {
    l *zk.Lock
}

func GetLock(zkconn *ZkConn, msg *Message) *ZkLock {
  acl := zk.WorldACL(zk.PermAll)
  l := zk.NewLock(zkconn.c, "/lock/"+msg.key, acl)
  zl := &ZkLock{}
  zl.l = l
  return zl
}

func (zl *ZkLock) Lock() error {
  e := zl.l.Lock()
  return e
}

func (zl *ZkLock) Unlock() error {
  e := zl.l.Unlock()
  return e
}
