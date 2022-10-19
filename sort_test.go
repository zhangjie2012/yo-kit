package kit

import "testing"

func TestSortByKeySF64(t *testing.T) {
	m := map[string]float64{
		"orange":     55.6,
		"apple":      77.7,
		"mango":      33.3,
		"strawberry": 99.9,
	}

	result1 := SortByKeySF64(m, true)
	for _, p := range result1 {
		t.Log(p.Key, p.Val)
	}

	result2 := SortByKeySF64(m, false)
	for _, p := range result2 {
		t.Log(p.Key, p.Val)
	}
}

func TestSortBValueSF64(t *testing.T) {
	m := map[string]float64{
		"orange":     55.6,
		"apple":      77.7,
		"mango":      33.3,
		"mango2":     33.3,
		"strawberry": 99.9,
	}

	result1 := SortByValueSF64(m, true)
	for _, p := range result1 {
		t.Log(p.Key, p.Val)
	}

	result2 := SortByValueSF64(m, false)
	for _, p := range result2 {
		t.Log(p.Key, p.Val)
	}
}
