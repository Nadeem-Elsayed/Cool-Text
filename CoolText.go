package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

//-------------------------------------------------------------------------
//Actual Code
var charsArray = [51]string{
	"!","@","#","$","%","^","&","*","(",")","{","}","[","]",":",";","'","","|","\\",",",".","/","?",
	"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"," ",
}
func main() {
	fmt.Println("This program prints text in a cool way, enter text: ")

	//reading multi-word input
	inputreader := bufio.NewReader(os.Stdin) //reads from stdin
	str, _ := inputreader.ReadString('\n')//stops at new-line
	var empty string = ""

	/*
	* Double for loop for a reason
	* first for loop cycles through each letter in the user inputted string
	* Second for loop prints values from charsArray until it matches the letter from user inputted string
	* time.Sleep() added to make it a little more fashionable
	*/
	for i := 0; i < len(str)-2; i++ {
		for j := 0; j <= len(charsArray); j++ {
			if charsArray[j] == strings.ToLower(string(str[i])){
				empty = empty + string(str[i])
				fmt.Println(empty)
				break
			} else if j < len(charsArray){
				fmt.Println(empty+charsArray[j])
			} else {//if the character is not within the charsArray
				empty = empty + string(str[i])
				fmt.Println(empty)
			}
			time.Sleep(5000000)//this is in nanoseconds
		}
		
	}
}