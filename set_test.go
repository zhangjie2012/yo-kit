package kit

import "testing"

func TestSliceUniq(t *testing.T) {
	a := []string{"a", "b", "c", "d", "a", "b", "d"}
	u := SliceUniq(a)
	t.Log(u)
}

func TestSliceUniqInt64(t *testing.T) {
	a := []int64{1, 1, 3, 1, 5}
	u := SliceUniqInt64(a)
	t.Log(u)
}

func TestSliceUniqInt(t *testing.T) {
	a := []int{1, 1, 3, 1, 5}
	u := SliceUniqInt(a)
	t.Log(u)
}

func TestSliceUnion(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	b := []string{"f", "e", "c", "d", "e", "f"}
	u := SliceUnion(a, b)

	t.Log(a)
	t.Log(b)
	t.Log(u)
}

func TestSliceIntersection(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	b := []string{"b", "c", "d", "e"}
	inter := SliceIntersection(a, b)
	t.Log(inter)
}

func TestSliceDifference(t *testing.T) {
	a := []string{"a", "b", "c", "d", "f"}
	b := []string{"b", "c", "d", "e"}
	diff := SliceDifference(a, b)
	t.Log(diff) // expect ["a", "f"]
}
