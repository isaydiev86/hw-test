package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	it := ListItem{Value: v}
	temp := l.front
	it.Next = temp
	l.front = &it
	if l.len == 0 {
		l.back = l.front
	} else {
		temp.Prev = &it
	}
	l.len++
	return &it
}

func (l *list) PushBack(v interface{}) *ListItem {
	it := ListItem{Value: v}
	temp := l.back
	it.Prev = temp
	l.back = &it
	if l.len == 0 {
		l.front = l.back
	} else {
		temp.Next = &it
	}
	l.len++
	return &it
}

func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next
	if prev != nil {
		prev.Next = next
	} else {
		l.front = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		l.back = prev
	}
	l.len--
	i.Next = nil
	i.Prev = nil
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i {
		return
	}
	l.PushFront(i.Value)
	prev := i.Prev
	next := i.Next
	if prev != nil {
		prev.Next = next
	} else {
		l.front = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		l.back = prev
	}
}

func NewList() List {
	return new(list)
}
