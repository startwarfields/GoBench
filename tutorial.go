package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)
var pl = fmt.Println

func main() {
  // Example code from tutorial https://www.youtube.com/watch?v=YzLrWHZa-Kc
  looparr()
}

func looparr() {

  aStr1 := "abcde"
  rArr :=[]rune(aStr1)
  for _, v := range rArr {
    fmt.Printf("Rune Array : %d \n", v)
  }
  byteArr := []byte{'a', 'b', 'c'}
  bStr := string (byteArr[:])
  pl ("I'm a string", bStr)
}

func mathme() {

  pl("5 + 4 =", 5+4)
}

func timeme() {

  now := time.Now()
  pl(now.Year(), now.Month(), now.Day())
  pl(now.Hour(), now.Minute(), now.Second())
}

func loopme() {
  rStr := "abcdefg"
  pl("Rune Count :", utf8.RuneCountInString(rStr))
  for i, runeVal := range rStr {
    fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
  }
}
func Convert() {
  cv3:= "5000000"
  cv4, err := strconv.Atoi(cv3)
  pl(cv4, err, reflect.TypeOf(cv4))
 
}
func Greet() {
  fmt.Println("Hello!")
  pl("What is your name?")
  reader := bufio.NewReader(os.Stdin)
  name, err := reader.ReadString('\n')

  if err == nil {
    pl("Hello", name)
  } else {
    log.Fatal(err)
  }
}
func AgeBirthday() {
  iAge := 8
  if (iAge >=1) && (iAge <18) {
    pl("Important Birthday")
  } else if (iAge == 21) || (iAge == 50) {
    pl("Important Birthday")
  }  else {
    pl("Not an important birthday")
  }
}
