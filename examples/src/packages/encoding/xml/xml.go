package main

import (
   "fmt"
   "os"
   "flag"
   "encoding/xml"
)

var inputFile = flag.String("inFile", "test.xml", "Input file path")

func main() {
   flag.Parse()
   xmlFile, err := os.Open(*inputFile)
   if err != nil {
      fmt.Println("Error opening file")
      return
   }
   defer xmlFile.Close()
   decoder := xml.NewDecoder(xmlFile)
   for {
      t,_ := decoder.Token()
      if t == nil {
         break
      }
      switch se := t.(type) {
      case xml.StartElement:
         fmt.Println("Start ", se.Name.Local)
      case xml.EndElement:
         fmt.Println("end ", se.Name.Local)
      case xml.CharData:
         fmt.Println("char data")
      case xml.Comment:
         fmt.Println("comment data")
      case xml.Directive:
         fmt.Println("directive")
      case xml.ProcInst:
         fmt.Println("ProcInst ", se.Target, se.Inst)
      default:
         fmt.Println("something else!!!")
      }
   }
}
