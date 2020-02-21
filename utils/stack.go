package utils

import (
	"bytes"
	"fmt"
)

type stack struct {
	data []interface{}
	top  int
	len  int
}

func NewStack(size int) *stack {
	return &stack{
		data: make([]interface{}, size),
		top:  -1,
	}
}

func (self *stack) Len() int {
	return self.len
}

func (self *stack) IsEmpty() bool {
	return self.top == -1
}

func (self *stack) Push(val interface{}) error {
	if self.top+1 == len(self.data) {
		self.expansion()
	}
	self.top += 1
	self.len += 1
	self.data[self.top] = val
	return nil
}

func (self *stack) Pop() (interface{}, bool) {
	if self.IsEmpty() {
		return -1, false
	}
	val := self.data[self.top]
	self.data[self.top] = nil
	self.top -= 1
	self.len -= 1
	return val, true
}

func (self *stack) expansion() {
	size := len(self.data)/2 + len(self.data)
	newArr := make([]interface{}, size)
	for idx := range self.data {
		newArr[idx] = self.data[idx]
	}
	self.data = newArr
}

func (self *stack) String() string {
	buf := bytes.NewBufferString("[")
	fmt.Println("len", len(self.data))
	for i := 0; i < len(self.data); i++ {
		if self.data[i] == nil {
			continue
		}
		if i != 0 {
			buf.WriteString(fmt.Sprintf(", %v", self.data[i]))
		} else {
			buf.WriteString(fmt.Sprintf("%v", self.data[i]))
		}
	}
	buf.WriteRune(']')
	return buf.String()
}
