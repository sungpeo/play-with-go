package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
  // Return a greeting that embes the name in a message.
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message
}
