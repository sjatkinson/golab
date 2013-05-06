//
// Simple ls example using go
//
package main

import (
       "io";
       "fmt";
       "os";
       "flag";
)


// TODO: command line handling
// -- -l
// -- -a
// -- -r
func main() {
  flag.Parse();
  var dir string = ".";
  if flag.NArg() > 0 {
     dir = flag.Args()[0];
  }
  dirs, err := io.ReadDir(dir);
  if len(dirs) == 0 {
     fmt.Printf("Can not open directory: err=%s\n", err.String());
     os.Exit(1);
  }
  for i := 0; i < len(dirs); i++ {
     fmt.Printf("%s\n",dirs[i].Name);
  }
}