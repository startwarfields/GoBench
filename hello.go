package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)
var p1 = fmt.Println

func main() {
  // Example code from tutorial https://www.youtube.com/watch?v=YzLrWHZa-Kc
 Convert() 
}
func Convert() {
  cv3:= "5000000"
  cv4, err := strconv.Atoi(cv3)
  p1(cv4, err, reflect.TypeOf(cv4))
 
}
func Greet() {
  fmt.Println("Hello!")
  p1("What is your name?")
  reader := bufio.NewReader(os.Stdin)
  name, err := reader.ReadString('\n')

  if err == nil {
    p1("Hello", name)
  } else {
    log.Fatal(err)
  }
}
