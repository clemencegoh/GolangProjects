package Parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type EncodedMessage struct {
	Header string
	Message string
}

// parses string only, split into header and text
func ParseMessageString(msg string) []EncodedMessage {
	toReturn := []EncodedMessage{}

	prevHeader := ""
	prevMessage := ""
	isNotBinary := regexp.MustCompile(`^[0-1].*`).MatchString
	for _, line := range strings.Split(msg, "\n"){
		if !isNotBinary(line){
			// Header may only be a single line
			if prevMessage != "" && prevHeader != ""{
				toReturn = append(toReturn, EncodedMessage{
					Header: prevHeader,
					Message: prevMessage,
				})
			}

			prevHeader = strings.TrimSpace(line)
			prevMessage = ""
		}else{
			prevMessage += strings.TrimSpace(line)
		}
	}

	// Get last
	toReturn = append(toReturn, EncodedMessage{
		Header: prevHeader,
		Message: prevMessage,
	})

	return toReturn
}

func ParseMessageTextFile(filepath string) []EncodedMessage {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	return ParseMessageString(string(data))
}
