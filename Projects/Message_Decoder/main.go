package main

import (
	"GolangProjects/Projects/Message_Decoder/Parser"
	"fmt"
	"path/filepath"
)

func main(){
	// Driver code to show testfile works
	path, err := filepath.Abs("Projects/Message_Decoder/Parser/testfile.txt")
	if err!=nil{
		panic(err)
	}
	encodedMessages := Parser.ParseMessageTextFile(path)
	for _, msg := range encodedMessages{
		fmt.Println(Parser.DecodeMessageBlock(msg.Header, msg.Message))
	}
}