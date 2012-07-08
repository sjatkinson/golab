package main

import "os/exec"
import "fmt"

func main() {
   cmd := exec.Command("sleep", "5")
   var e error
   select {
      case e = cmd.Wait():
         fmt.Printf("Cmd is done\n")
      default:
         fmt.Printf("default\n")
         }
         }

