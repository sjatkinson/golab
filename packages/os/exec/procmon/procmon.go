/*
   Simple program starting and monitoring processes
*/
package main

import (
   "os/exec"
   "log"
)

func start(proc, args string) *exec.Cmd {
   cmd := exec.Command(proc, args)
   log.Printf("Starting %s %s", proc, args)
   err := cmd.Start()
   if err != nil {
      log.Fatal(err)
   }
   return cmd
}

func wait(proc *exec.Cmd, ch chan *exec.Cmd) {
   err := proc.Wait()
   if err != nil {
      log.Fatal(err)
   }
   ch <- proc
}

func report(c *exec.Cmd) {
   log.Printf("process done %s %s", c.Path, c.Args)

   log.Printf("ProcessState %s", c.ProcessState)
   log.Printf("SystemTime %s", c.ProcessState.SystemTime())
   log.Printf("UserTime %s", c.ProcessState.UserTime())
   log.Printf("Sys %d", c.ProcessState.Sys())
   log.Printf("Pid %d", c.ProcessState.Pid())
}


func main() {
   cmd1 := start("sleep", "5")
   cmd2 := start("sleep", "3")
   ch := make(chan *exec.Cmd)
   go wait(cmd1, ch)
   go wait(cmd2, ch)
   c := <- ch
   report(c)
   c = <- ch
   report(c)
}
