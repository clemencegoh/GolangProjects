package Parser

import (
	"strconv"
	"strings"
)

/*
Class for parsing message once it has been separated into a block.
Intended to be used only with:
- 1 Header
- 1 Message
*/

var (
	customSymbols []string
	keyStringMapping map[string]int
)


// function to generate all binary strings of length n
func GenerateAllBinaryStrings(n int, arr []string, i int) {
	if i == n{
		nextBinString := strings.Join(arr, "")
		if strings.Contains(nextBinString, "0"){
			customSymbols = append(customSymbols, nextBinString)
		}
		//fmt.Println(customSymbols)
		//fmt.Println(arr)
		return
	}

	arr[i] = "0"
	GenerateAllBinaryStrings(n, arr, i + 1)

	arr[i] = "1"
	GenerateAllBinaryStrings(n, arr, i + 1)

}


func buildKeyStrings(length int) {
	// Builds keystrings from binary up to length specified
	for i := 1;i<length + 1;i++{
		// build
		GenerateAllBinaryStrings(i, make([]string, i), 0)
	}
}

func init(){
	// build key strings
	buildKeyStrings(7)
	keyStringMapping = make(map[string]int)

	// map key strings to their integer
	for pos, i := range customSymbols{
		keyStringMapping[i] = pos
	}
}


func DecodeMessageBlock(header string, encoded_string string) string {
	// Decodes single block, 1 header and 1 message only
	// mapRunes
	mappedRunes := mapItems(header)

	// iteratively convert
	newMessage := encoded_string
	size := 0
	var overallMessage []rune
	for newMessage != "000"{
		size, newMessage = getSizeMode(newMessage)
		newMessage, overallMessage = getMessageMode(newMessage, size, mappedRunes, overallMessage)
	}

	return string(overallMessage)
}

func mapItems(header string) map[int]rune{
	mappedRunes := map[int]rune{}

	// map numbers to corresponding single character string
	for pos, char := range header{
		mappedRunes[pos] = char
	}

	return mappedRunes
}

func getSizeMode(message string) (size int, restOfMessage string){
	// State: Determine the message size
	size = 1
	restOfMessage = message[3:]
	tempSize, err := strconv.ParseInt(message[:3], 2, 64)
	if err!=nil{
		panic(err)
	}
	size = int(tempSize)

	// Strange behaviour
	if size <= 0{
		size = 1
	}
	return
}

func getMessageMode(message string, size int, runeMap map[int]rune, currentRunes []rune) (string, []rune) {
	// State: Determine message using fixed size
	nextChar := message[:size]
	if !strings.Contains(nextChar, "0"){
		// stop here
		return message[size:], currentRunes
	}

	// int representation of binary key string
	pos := keyStringMapping[nextChar]

	// Call itself
	return getMessageMode(message[size:], size, runeMap, append(currentRunes, runeMap[pos]))
}

