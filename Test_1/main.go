package main

import "fmt"

type SpacialStack struct {
	dataStack []int
	minStack  []int
}

// 入栈--压入栈
func (s *SpacialStack) Push(value int) {
	s.dataStack = append(s.dataStack, value)
	if len(s.minStack) == 0 || value <= s.minStack[len(s.dataStack)-1] {
		s.minStack = append(s.minStack, value)
	}
}

func (s *SpacialStack) Pop() int {
	if len(s.dataStack) == 0 {
		return -1
	}
	poppedValue := s.dataStack[len(s.dataStack)-1]
	s.dataStack = s.dataStack[:len(s.dataStack)-1]
	if poppedValue == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)]
	}
	return poppedValue

}

func (s *SpacialStack) GetMin() int {
	if len(s.minStack) == 0 {
		return -1

	}
	return s.minStack[len(s.minStack)-1]
}

func main() {
	s := SpacialStack{}
	s.Push(5)
	s.Push(3)
	s.Push(7)
	s.Push(7)
	s.Push(2)
	fmt.Println("Minimum element:", s.GetMin())
	s.Pop()
	fmt.Println("Minimum element:", s.GetMin())
}
