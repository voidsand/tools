package tools

import (
	"reflect"
	"testing"
)

func TestNewCollection(t *testing.T) {
	got := NewCollection()
	want := collection{make([]interface{}, 0)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_IsEmpty(t *testing.T) {
	c := NewCollection()
	got := c.IsEmpty()
	if got != true {
		t.Errorf("got %v want true", got)
	}
}

func TestCollection_Len(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	got := c.Len()
	want := 3
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_All(t *testing.T) {
	c := NewCollection()
	got := c.All()
	want := c.data
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_Get(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	got := c.Get(1)
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_Find(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	got := c.Find(3)
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_Push(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	var got int
	for _, v := range c.All() {
		if n, ok := v.(int); ok {
			got += n
		}
	}
	want := 6
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_PopFront(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.PopFront()
	var got int
	for _, v := range c.All() {
		if n, ok := v.(int); ok {
			got += n
		}
	}
	want := 5
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_PopBack(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.PopBack()
	var got int
	for _, v := range c.All() {
		if n, ok := v.(int); ok {
			got += n
		}
	}
	want := 3
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_Insert(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.Insert(1, 4)
	var got int
	for _, v := range c.All() {
		if n, ok := v.(int); ok {
			got += n
		}
	}
	want := 10
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestCollection_Remove(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.Remove(1)
	var got int
	for _, v := range c.All() {
		if n, ok := v.(int); ok {
			got += n
		}
	}
	want := 4
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollection_Clear(t *testing.T) {
	c := NewCollection()
	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.Clear()
	var got int
	for _, v := range c.All() {
		if n, ok := v.(int); ok {
			got += n
		}
	}
	want := 0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
