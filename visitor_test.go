package visitor

import "testing"

func TestVisitBasicType(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	v := 1
	Visit(v, visitor)
	if len(visited) != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[0].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
}

func TestVisitPointer(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	v := 1
	Visit(&v, visitor)
	if len(visited) != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[0].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
}

func TestVisitSlice(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	v := []int{1, 2, 3, 4, 5}
	Visit(v, visitor)
	if len(visited) != 6 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]int); !ok {
		t.Fatal("visited")
	}
	for i := 1; i <= 5; i++ {
		if v, ok := visited[i].(int); !ok || v != i {
			t.Fatal("visited")
		}
	}
}

func TestVisitStruct(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	type Foo struct {
		Foo int
		Bar string
		Baz bool
	}
	v := Foo{
		42, "foo", true,
	}
	Visit(v, visitor)
	if len(visited) != 4 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].(Foo); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(int); !ok || v != 42 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].(string); !ok || v != "foo" {
		t.Fatal("visited")
	}
	if v, ok := visited[3].(bool); !ok || v != true {
		t.Fatal("visited")
	}
}

func TestVisitMap(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	type Foo map[int]int
	v := Foo{
		1: 1,
		2: 2,
		3: 3,
	}
	Visit(v, visitor)
	if len(visited) != 4 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].(Foo); !ok || len(v) != 3 {
		t.Fatal("visited")
	}
}

func TestPartialVisitSlice(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if v, ok := v.(int); ok && v == 2 {
			return nil
		}
		return visitor
	}
	v := []int{1, 2, 3, 4, 5}
	Visit(v, visitor)
	if len(visited) != 3 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]int); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].(int); !ok || v != 2 {
		t.Fatal("visited")
	}
}

func TestPartialVisitSlice2(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if _, ok := v.([]int); ok {
			return nil
		}
		return visitor
	}
	v := []int{1, 2, 3, 4, 5}
	Visit(v, visitor)
	if len(visited) != 1 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]int); !ok {
		t.Fatal("visited")
	}
}
