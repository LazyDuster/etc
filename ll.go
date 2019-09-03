package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	length int
	first *node
}

func main() {
	fmt.Printf("linked lists lol\n")
	l := list{}

	for i := 0; i < 100; i++ {
		l.append(i)
	}

	for j := 0; j < 100; j++ {
		if j % 2 == 0 {
			l.delete(j)
			l.prepend(j)
		}
	}

	l.print()
}

func (l *list) print() {
	list := l.first
	for list != nil {
		fmt.Printf("%d\n", list.data)
		list = list.next
	}
	fmt.Println()
}

func (l *list) prepend(value int) {
	if l.first == nil {
		l.first = &node{value, nil}
	} else {
		temp := node{value, l.first}
		l.first = &temp
	}
	l.length++
}

func (l *list) append(value int) {
	if l.first == nil {
		l.first = &node{value, nil}
	} else {
		temp := l.first
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &node{value, nil}
	}
	l.length++
}

func (l *list) delete(value int) {
	if l.first == nil {
		fmt.Printf("ay i'm walkin here\n")
		return
	} else {
		temp := l.first
		if temp.data == value {
			l.first = l.first.next
			l.length--
			return
		} else {
			for temp.next != nil {
				if temp.next.data == value {
					temp.next = temp.next.next
					l.length--
					return
				} else {
					temp = temp.next
				}
			}
		}
	}
}