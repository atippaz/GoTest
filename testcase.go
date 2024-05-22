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
		fmt.Println(fmt.Sprintf("------------------- %s ----------------", strconv.Itoa(index+1)))
		fmt.Println(fmt.Sprintf("code : %s", item.input))
		fmt.Println(fmt.Sprintf("answer : %s", item.expected))
		result := encoding(item.input)
		fmt.Println(fmt.Sprintf("encode result : %s",result ))
		fmt.Println(fmt.Sprintf("is correct : %s",strconv.FormatBool(result == item.expected) ))
	}
}


func encoding(code string)string{
	stringIsEncode := ""
	// สร้าง 1010101
	coded :=""
	hasChange := false
	for i:=0;i<len(code) ;i++{
		symbol :=code[i:i+1]
		// if(hasChange){
		// 	coded = backTestAndChangeValue(stringIsEncode,coded)
		// 	hasChange = false
		// 	continue
		// }
		stringIsEncode+=symbol
		left := 0
		right:=0
		temp:=""
		if symbol == "L" {
			left = 1
		} else if symbol == "R" {
			right = 1
		} 
		temp = strconv.Itoa(left)+strconv.Itoa(right)
		if(len(coded)<2){
			coded += temp
			continue
		}
		leftValue, err := strconv.Atoi(coded[len(coded)-1:])
        if err != nil {
			panic(err)
        }
        rightValue, err := strconv.Atoi(temp[0 :1])
        if err != nil {
			panic(err)
        }
		startWithoutEnd := coded[:len(coded)-1]
		addString := temp[1:] 
		combinedNumber:=0   
		if(!hasChange){
			combinedNumber = leftValue + rightValue
			hasChange = false
		}
		coded = startWithoutEnd + strconv.Itoa(combinedNumber) + addString
		coded,hasChange = backTestAndChangeValue(stringIsEncode,coded)
	}
	coded,hasChange = backTestAndChangeValue(stringIsEncode,coded)
	return coded
}

// func addLastString(coded string,symbol string)string{
	
// }
func backTestAndChangeValue(code string,encode string)(string, bool){
	_encode:=encode
	isChange:=false
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
			isChange = true
		} else if symbol == "R" &&leftValue>=rightValue {
			rightValue+=1
			hasChange = true
			isChange = true
		} else if symbol == "=" &&leftValue!=rightValue {
			rightValue = leftValue
			hasChange = true
			isChange = true
		}
		_encode = _encode[0:i-1] + strconv.Itoa(leftValue)+strconv.Itoa(rightValue) +_encode[i+1:]
		if(hasChange){
			i+=1
		}
	}
	return _encode,isChange
}