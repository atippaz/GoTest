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
		fmt.Println(fmt.Sprintf("------------------- %s ----------------", strconv.Itoa(index)))
		result := encoding(item.input)
		fmt.Println(fmt.Sprintf("code : %s", item.input))
		fmt.Println(fmt.Sprintf("answer : %s", item.expected))
		fmt.Println(fmt.Sprintf("encode result : %s",result ))
		fmt.Println(fmt.Sprintf("is correct : %s",strconv.FormatBool(result == item.expected) ))
	}
}


func encoding(code string)string{
	encoded:=""
	stringIsEncode := ""
	// สร้าง 1010101
	for i:=0;i<len(code) ;i++{
		symbol:=code[i:i+1]
		value:=""
		if symbol == "L" {
			value = "10"
		} else if symbol == "R" {
			value = "01"
		} else if symbol == "=" {
			value = "00"
		}
		stringIsEncode+=symbol
		encoded = addLastString(encoded,value)
		if(len(stringIsEncode)>1){
			encoded = backTestAndChangeValue(stringIsEncode,encoded)
		}
	}
	return encoded
}

func addLastString(started string,adding string)string{
	if len(started)==0 {
		return adding
	}
	endStartText, err := strconv.Atoi(started[len(started)-1:])
    if err != nil {
        panic(err)
    }
    startWithoutEnd := started[:len(started)-1]
    startAdd, err := strconv.Atoi(adding[0:1])
    if err != nil {
        panic(err) 
    }
    addString := adding[1:]    
    combinedNumber := endStartText + startAdd
    result := startWithoutEnd + strconv.Itoa(combinedNumber) + addString
	return result
}
func backTestAndChangeValue(code string,encode string)string{
	_encode:=encode
	for i:=len(code);i>0  ;i--{
		hasChange := false
		leftValue, err := strconv.Atoi(_encode[i-1: i])
        if err != nil {
			panic(err)
        }
        rightValue, err := strconv.Atoi(_encode[i : i+1])
        if err != nil {
			panic(err)
        }
		symbol:=code[i-1:i]
		if symbol == "L" && leftValue <= rightValue {
			leftValue+=1
			hasChange = true
		} else if symbol == "R" &&leftValue>=rightValue {
			rightValue+=1
			hasChange = true
		} else if symbol == "=" &&leftValue!=rightValue {
			rightValue = leftValue
			hasChange = true
		}
		_encode = _encode[0:i-1] + strconv.Itoa(leftValue)+strconv.Itoa(rightValue) +_encode[i+1:]
		if(hasChange){
			i++
		}
	}
	return _encode
}