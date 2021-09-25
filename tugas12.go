package main

import "fmt"
import "regexp"

func main(){
  var a string = "ini adalah SESUATU yang biasa saja"
  var regex, err = regexp.Compile(`[A-Z]+`)

  if err!=nil{
    fmt.Println(err.Error())
  }

  hasil1:= regex.FindAllString(a,-1)
  hasil2:= regex.FindStringIndex(a)
  fmt.Println(hasil1,hasil2)
}
