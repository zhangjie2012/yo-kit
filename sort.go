package kit

import (
	"sort"

	"github.com/zhangjie2012/yo-kit/datastruct"
)

// SortByKeysSF64 <string, float64> map sort by keys
// [via]: https://www.geeksforgeeks.org/how-to-sort-golang-map-by-keys-or-values/
func SortByKeySF64(m map[string]float64, asc bool) []*datastruct.SF64Pair {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	if asc {
		sort.Strings(keys)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}

	results := make([]*datastruct.SF64Pair, 0, len(keys))
	for _, k := range keys {
		results = append(results, &datastruct.SF64Pair{
			Key: k,
			Val: m[k],
		})
	}

	return results
}

// SortByValueSF64 <string, float64> map sort by values
func SortByValueSF64(m map[string]float64, asc bool) []*datastruct.SF64Pair {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	if asc {
		sort.SliceStable(keys, func(i, j int) bool {
			return m[keys[i]] < m[keys[j]]
		})
	} else {
		sort.SliceStable(keys, func(i, j int) bool {
			return m[keys[i]] > m[keys[j]]
		})
	}

	results := make([]*datastruct.SF64Pair, 0, len(keys))
	for _, k := range keys {
		results = append(results, &datastruct.SF64Pair{
			Key: k,
			Val: m[k],
		})
	}

	return results
}
