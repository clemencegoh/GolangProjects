package Parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Simple_Message(t *testing.T){
	msg := DecodeMessageBlock("TNM AEIOU", "0010101100011101000100111011001111000")
	fmt.Println(msg)
	assert.Equal(t, msg, "TAN ME", "")
}

func Test_Parse_Message_Text(t *testing.T){
	text := "TNM AEIOU\n0010101100011101000100111011001111000"
	newText := ParseMessageString(text)
	assert.Equal(t, newText[0].Header, "TNM AEIOU")
	assert.Equal(t, newText[0].Message, "0010101100011101000100111011001111000")
}

func Test_Parse_Message_Text_Longer(t *testing.T){
	text := "TNM AEIOU\n0010101100011101000100111011001111000\nTNM AEIOB\n0010101100011101000100111011001111000"
	newText := ParseMessageString(text)
	assert.Equal(t, newText[0].Header, "TNM AEIOU")
	assert.Equal(t, newText[0].Message, "0010101100011101000100111011001111000")
	assert.Equal(t, newText[1].Header, "TNM AEIOB")
	assert.Equal(t, newText[1].Message, "0010101100011101000100111011001111000")
}

func Test_Parse_Message_File(t *testing.T){
	filename := "testfile.txt"
	newText := ParseMessageTextFile(filename)
	assert.Equal(t, newText[0].Header, "TNM AEIOU")
	assert.Equal(t, newText[0].Message, "0010101100011101000100111011001111000")
	assert.Equal(t, newText[1].Header, `$#**\`)
	assert.Equal(t, newText[1].Message, "0100000101101100011100001000")
}
