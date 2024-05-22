package main
import (
	"fmt"
	"strconv"
)
type TestCase struct {
    input string
	expected string
}
func main(){
	testCase := [...]TestCase{
		{
			input:"LLRR=",
			expected:"210122",
		},
		{
			input:"==RLL",
			expected:"000210",
		},
		{
			input:"=LLRR",
			expected:"221012",
		},
		{
			input:"RRL=R",
			expected:"012001",
		},
	}
	for index, item := range testCase {
		result := encode(item.input)
		fmt.Println(fmt.Sprintf("------------------- %s ----------------", strconv.Itoa(index)))
		fmt.Println(fmt.Sprintf("code : %s", item.input))
		fmt.Println(fmt.Sprintf("answer : %s", item.expected))
		fmt.Println(fmt.Sprintf("encode result : %s",result ))
		fmt.Println(fmt.Sprintf("is correct : %s",strconv.FormatBool(result == item.expected) ))
	}
}

func encode(code string)string{
	currentNumber := 0
    if code[0] == 'R' {
        currentNumber = 0
    } else if code[0] == '=' {
		if(len(code)<=1){
			currentNumber = 0
		} else if code[1] == 'L' {
			currentNumber  = 2
		} else if code[1] == 'R' || code[1] == '=' {
			currentNumber = 0
		}
    } else if code[0] == 'L'{
		currentNumber  = 2
	}
	result:=""
	firstIter := true
	for i:=0;i<len(code) || firstIter;i++{
		if firstIter {
			currentNumber = createAlphabet('=',currentNumber)
			result+=strconv.Itoa(currentNumber)
			firstIter = false
			i-=0
		}
		currentNumber = createAlphabet(code[i],currentNumber)
		result+=strconv.Itoa(currentNumber)
	}
	return result
}
func createAlphabet(symbol byte,currentValue int)int{
	if symbol == 'L' {
        return currentValue -1
    } else if symbol == 'R' {
        return currentValue + 1
    } else if symbol == '=' {
        return currentValue
    }
    return 0
}