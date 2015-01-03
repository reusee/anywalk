package anywalk

import "testing"

func TestVisitWithKeyBasic(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	var visited []Visited
	visitor = func(v, key interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, key})
		return visitor
	}

	type Foo struct {
		IntPtr *int
		Slice  []int
		Map    map[int]string
	}
	i := 5
	v := Foo{
		IntPtr: &i,
		Slice:  []int{1, 2, 3},
		Map: map[int]string{
			1: "foo",
			2: "bar",
			3: "baz",
		},
	}
	WalkWithKey(v, visitor)

	if len(visited) != 11 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.(Foo); !ok || visited[0].Key != nil {
		t.Fatal("visited")
	}
	if _, ok := visited[1].Value.(*int); !ok {
		t.Fatal("visited")
	}
	if key, ok := visited[1].Key.(string); !ok || key != "IntPtr" {
		t.Fatal("visited")
	}
	if _, ok := visited[2].Value.(int); !ok || visited[2].Key != nil {
		t.Fatal("visited")
	}
	if _, ok := visited[3].Value.([]int); !ok {
		t.Fatal("visited")
	}
	if key, ok := visited[3].Key.(string); !ok || key != "Slice" {
		t.Fatal("visited")
	}
	if v, ok := visited[4].Value.(int); !ok || v != 1 {
		t.Fatal("visited")
	}
	if key, ok := visited[4].Key.(int); !ok || key != 0 {
		t.Fatal("visited")
	}
	if v, ok := visited[5].Value.(int); !ok || v != 2 {
		t.Fatal("visited")
	}
	if key, ok := visited[5].Key.(int); !ok || key != 1 {
		t.Fatal("visited")
	}
	if v, ok := visited[6].Value.(int); !ok || v != 3 {
		t.Fatal("visited")
	}
	if key, ok := visited[6].Key.(int); !ok || key != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[7].Value.(map[int]string); !ok {
		t.Fatal("visited")
	}
	if key, ok := visited[7].Key.(string); !ok || key != "Map" {
		t.Fatal("visited")
	}
	if _, ok := visited[8].Value.(string); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[8].Key.(int); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[9].Value.(string); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[9].Key.(int); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[10].Value.(string); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[10].Key.(int); !ok {
		t.Fatal("visited")
	}
}

func TestPartialVisitWithKey(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		if _, ok := v.(*int); ok {
			return nil
		}
		return visitor
	}
	i := 5
	v := []*int{&i}
	WalkWithKey(v, visitor)
	if len(visited) != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.([]*int); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
	if _, ok := visited[1].Value.(*int); !ok {
		t.Fatal("visited")
	}
	if key, ok := visited[1].Key.(int); !ok || key != 0 {
		t.Fatal("visited")
	}
}

func TestPartialVisitWithKey2(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		if _, ok := v.(int); ok {
			return nil
		}
		return visitor
	}
	i := 5
	v := []*int{&i}
	WalkWithKey(v, visitor)
	if len(visited) != 3 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.([]*int); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
	if _, ok := visited[1].Value.(*int); !ok {
		t.Fatal("visited")
	}
	if key, ok := visited[1].Key.(int); !ok || key != 0 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].Value.(int); !ok || v != 5 {
		t.Fatal("visited")
	}
	if visited[2].Key != nil {
		t.Fatal("visited")
	}
}

func TestPartialSliceVisitWithKey(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		if _, ok := v.([]*int); ok {
			return nil
		}
		return visitor
	}
	i := 5
	v := []*int{&i}
	WalkWithKey(v, visitor)
	if len(visited) != 1 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.([]*int); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
}

func TestPartialStructkVisitWithKey(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	type Foo struct {
		I int
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		if _, ok := v.(Foo); ok {
			return nil
		}
		return visitor
	}
	v := []Foo{
		{42},
	}
	WalkWithKey(v, visitor)
	if len(visited) != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.([]Foo); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
	if v, ok := visited[1].Value.(Foo); !ok || v.I != 42 {
		t.Fatal("visited")
	}
	if key, ok := visited[1].Key.(int); !ok || key != 0 {
		t.Fatal("visited")
	}
}

func TestPartialStructkVisitWithKey2(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	type Foo struct {
		I int
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		if _, ok := v.(int); ok {
			return nil
		}
		return visitor
	}
	v := []Foo{
		{42},
	}
	WalkWithKey(v, visitor)
	if len(visited) != 3 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.([]Foo); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
	if v, ok := visited[1].Value.(Foo); !ok || v.I != 42 {
		t.Fatal("visited")
	}
	if key, ok := visited[1].Key.(int); !ok || key != 0 {
		t.Fatal("visited")
	}
	if v, ok := visited[2].Value.(int); !ok || v != 42 {
		t.Fatal("visited")
	}
	if key, ok := visited[2].Key.(string); !ok || key != "I" {
		t.Fatal("visited")
	}
}

func TestPartialMapkVisitWithKey(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		if _, ok := v.(map[string]string); ok {
			return nil
		}
		return visitor
	}
	v := []map[string]string{
		{
			"foo": "FOO",
			"bar": "BAR",
		},
	}
	WalkWithKey(v, visitor)
	if len(visited) != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.([]map[string]string); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
	if _, ok := visited[1].Value.(map[string]string); !ok {
		t.Fatal("visited")
	}
	if key, ok := visited[1].Key.(int); !ok || key != 0 {
		t.Fatal("visited")
	}
}

func TestPartialMapkVisitWithKey2(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		if _, ok := v.(string); ok {
			return nil
		}
		return visitor
	}
	v := []map[string]string{
		{
			"foo": "FOO",
			"bar": "BAR",
		},
	}
	WalkWithKey(v, visitor)
	if len(visited) != 3 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.([]map[string]string); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
	if _, ok := visited[1].Value.(map[string]string); !ok {
		t.Fatal("visited")
	}
	if key, ok := visited[1].Key.(int); !ok || key != 0 {
		t.Fatal("visited")
	}
	if _, ok := visited[2].Value.(string); !ok {
		t.Fatal("visited")
	}
	if _, ok := visited[2].Key.(string); !ok {
		t.Fatal("visited")
	}
}

func TestVisitWithKeyInvalidValue(t *testing.T) {
	var visitor VisitorWithKey
	type Visited struct {
		Value, Key interface{}
	}
	var visited []Visited
	visitor = func(v, k interface{}) VisitorWithKey {
		visited = append(visited, Visited{v, k})
		return visitor
	}
	type Foo struct {
		P *int
	}
	v := Foo{}
	WalkWithKey(v, visitor)
	if len(visited) != 2 {
		t.Fatal("visited")
	}
	if _, ok := visited[0].Value.(Foo); !ok {
		t.Fatal("visited")
	}
	if visited[0].Key != nil {
		t.Fatal("visited")
	}
	if _, ok := visited[1].Value.(*int); !ok {
		t.Fatal("visited")
	}
	if v, ok := visited[1].Key.(string); !ok || v != "P" {
		t.Fatal("visited")
	}
}
