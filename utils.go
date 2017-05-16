package zksimulator

import (
  "fmt"
  "github.com/fatih/color"
)

var red = color.New(color.FgRed)

var boldRed = red.Add(color.Bold)

var green = color.New(color.FgGreen)

var boldGreen = green.Add(color.Bold)

var cyan = color.New(color.FgCyan)

var boldCyan = cyan.Add(color.Bold)

func detailLog(text, senderId string, node1, node2 Instance) {
     if Env.Debug {
       fmt.Println(text+" "+node1.getIdent()+" "+node2.getIdent()+" "+senderId)
     }
}

func clientLog(text, senderId string, node1 Instance) {
     if Env.Debug {
       fmt.Println(text+" "+node1.getIdent()+" "+senderId)
     }
}

func simpleLog(text string){
    if Env.Debug {
      fmt.Println(text)
    }
}

func failLog(text, senderId string, node1, node2 Instance) {
    if Env.Debug {
     boldRed.Println(text+" "+node1.getIdent()+" "+node2.getIdent()+" "+senderId)
    }
}

func successLog(text, senderId string, node1 Instance) {
      boldGreen.Println(text+" "+node1.getIdent()+" "+senderId)
}

func noLockgiven(text, senderId string, node1 Instance) {
    if Env.Debug {
     boldCyan.Println(text+" "+node1.getIdent()+" "+senderId)
   }
}
