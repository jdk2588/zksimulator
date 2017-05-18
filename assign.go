package zksimulator

type Instance interface {
    setRack(int)
    setRegion(int)
    getRack() int
    getRegion() int
    getIdent() string
}

func assignRackandRegion(node Instance, index int) {

  switch index {
  case 0,4,8:
          node.setRack(1)
          node.setRegion(1)
  case 1,5,9:
          node.setRack(1)
          node.setRegion(2)
  case 2,6,10:
          node.setRack(2)
          node.setRegion(1)
  case 3,7,11:
          node.setRack(2)
          node.setRegion(2)
  }
}
