package stack

import "fmt"

//Stack is an array of string that represents stack
type Stack []string

//IsEmpty checks whether stack is empty
func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

//Push pushes ele on stack
func (s *Stack) Push(str string) {

	*s = append(*s, str)

}

//Pop pops top ele
func (s *Stack) Pop() bool {
	if len(*s) == 0 {
		return false
	}
	*s = (*s)[:len(*s)-1]
	return true
}

//Top reterns stacktop
func (s *Stack) Top() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

//PrintStack prints complete stack
func (s *Stack) PrintStack() {
	fmt.Println((*s))
}
