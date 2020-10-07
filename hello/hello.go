package main

import (
  "fmt"

  "hulkook.com/greetings"
)

func main() {
  // Get a greeting message and print it.
  message := greetings.Hello("Eli")
  fmt.Println(message)
}
