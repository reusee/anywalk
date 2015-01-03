package anywalk

import "testing"

func TestVisitBasicType(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	v := 1
	Walk(v, visitor)
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
	Walk(&v, visitor)
	if len(visited) != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].(*int); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
}

func TestPartialVisitPointer(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if _, ok := v.(*int); ok {
			return nil
		}
		return visitor
	}
	v := 1
	Walk(&v, visitor)
	if len(visited) != 1 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].(*int); !ok {
		t.Fatal("visited")
	}
}

func TestPartialVisitPointer2(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if _, ok := v.(int); ok {
			return nil
		}
		return visitor
	}
	v := []int{1, 2, 3}
	Walk(&v, visitor)
	if len(visited) != 3 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].(*[]int); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[1].([]int); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[2].(int); !ok || v != 1 {
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
	Walk(v, visitor)
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

func TestVisitSliceOfStruct(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	type Foo struct {
		I int
	}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	v := []Foo{
		{1}, {2}, {3},
	}
	Walk(v, visitor)
	if len(visited) != 7 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]Foo); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(Foo); !ok || v.I != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[3].(Foo); !ok || v.I != 2 {
		t.Fatal("visited")
	}
	if v, ok := visited[4].(int); !ok || v != 2 {
		t.Fatal("visited")
	}
	if v, ok := visited[5].(Foo); !ok || v.I != 3 {
		t.Fatal("visited")
	}
	if v, ok := visited[6].(int); !ok || v != 3 {
		t.Fatal("visited")
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
	Walk(v, visitor)
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
	Walk(v, visitor)
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
	Walk(v, visitor)
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
	Walk(v, visitor)
	if len(visited) != 1 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]int); !ok {
		t.Fatal("visited")
	}
}

func TestPartialVisitStruct(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	type Foo struct {
		I int
	}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if v, ok := v.(Foo); ok && v.I == 2 {
			return nil
		}
		return visitor
	}
	v := []Foo{
		{1}, {2}, {3},
	}
	Walk(v, visitor)
	if len(visited) != 4 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]Foo); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(Foo); !ok || v.I != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[3].(Foo); !ok || v.I != 2 {
		t.Fatal("visited")
	}
}

func TestPartialVisitStruct2(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	type Foo struct {
		I int
	}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if v, ok := v.(int); ok && v == 1 {
			return nil
		}
		return visitor
	}
	v := []Foo{
		{1}, {2}, {3},
	}
	Walk(v, visitor)
	if len(visited) != 3 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]Foo); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(Foo); !ok || v.I != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
}

func TestPartialVisitMap(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if _, ok := v.(map[int]int); ok {
			return nil
		}
		return visitor
	}
	v := []map[int]int{
		{1: 1},
		{1: 2},
	}
	Walk(v, visitor)
	if len(visited) != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]map[int]int); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(map[int]int); !ok || v[1] != 1 {
		t.Fatal("visited")
	}
}

func TestPartialVisitMap2(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		if v, ok := v.(int); ok && v == 2 {
			return nil
		}
		return visitor
	}
	v := []map[int]int{
		{1: 1},
		{1: 2},
		{1: 3},
	}
	Walk(v, visitor)
	if len(visited) != 5 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].([]map[int]int); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].(map[int]int); !ok || v[1] != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].(int); !ok || v != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[3].(map[int]int); !ok || v[1] != 2 {
		t.Fatal("visited")
	}
	if v, ok := visited[4].(int); !ok || v != 2 {
		t.Fatal("visited")
	}
}

func TestVisitInvalidValue(t *testing.T) {
	var visitor Visitor
	var visited []interface{}
	visitor = func(v interface{}) Visitor {
		visited = append(visited, v)
		return visitor
	}
	type Foo struct {
		P *int
	}
	v := Foo{}
	Walk(v, visitor)
	if len(visited) != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].(Foo); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[1].(*int); !ok {
		t.Fatal("visited")
	}
}
