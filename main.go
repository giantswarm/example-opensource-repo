// Package main does all the things
package main

import "fmt"
import "github.com/spf13/viper"

// This is where no magic happens
func main() {
	// setting our default output message
	viper.SetDefault("Message", "Hello world, you rock!")
	// printing that message
	fmt.Println(viper.Get("Message"))
}
