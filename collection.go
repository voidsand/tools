package tools

type collection struct {
	data []interface{}
}

func NewCollection() collection {
	return collection{make([]interface{}, 0)}
}

func (c collection) IsEmpty() bool {
	return len(c.data) <= 0
}

func (c collection) Len() int {
	return len(c.data)
}

func (c collection) All() []interface{} {
	return c.data
}

func (c collection) Get(index int) interface{} {
	if index < 0 || index >= len(c.data) {
		return nil
	}
	return c.data[index]
}

func (c collection) Find(item interface{}) int {
	for i, v := range c.data {
		if v == item {
			return i
		}
	}
	return -1
}

func (c *collection) Push(item interface{}) {
	c.data = append(c.data, item)
}

func (c *collection) PopFront() {
	if len(c.data) <= 0 {
		return
	}
	c.data = append(c.data[:0], c.data[1:]...)
}

func (c *collection) PopBack() {
	if len(c.data) <= 0 {
		return
	}
	c.data = append(c.data[:len(c.data)-1], c.data[len(c.data):]...)
}

func (c *collection) Insert(index int, item interface{}) {
	if index < 0 || index >= len(c.data) {
		return
	}
	rear := append([]interface{}{}, c.data[index:]...)
	c.data = append(c.data[:index], item)
	c.data = append(c.data, rear...)
}

func (c *collection) Remove(index int) {
	if index < 0 || index >= len(c.data) {
		return
	}
	c.data = append(c.data[:index], c.data[index+1:]...)
}

func (c *collection) Clear() {
	c.data = append(c.data[:0], c.data[len(c.data):]...)
}
