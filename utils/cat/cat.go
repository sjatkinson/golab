package main

import (
   "os"
   "io"
   "flag"
   "log"
   "fmt"
)

func main() {
   flag.Parse()
   file, err := os.Open(flag.Arg(0))
   if err != nil {
      log.Fatal(err)
   }
   copied, err := io.Copy(os.Stdout, file)
   if err != nil {
      log.Fatal(err)
   }
   err = os.Stdout.Sync()
   if err != nil {
      log.Fatal(err)
   }
   fmt.Printf("copied %d\n", copied)
}
