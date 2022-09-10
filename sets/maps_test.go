package sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapDifference(t *testing.T) {
	map1, map2 := make(map[int]int), make(map[int]int)
	mapLen := 10
	map1Start, map2Start := 0, mapLen/2
	for i := map1Start; i < mapLen+map1Start; i++ {
		map1[i] = i
	}
	for i := map2Start; i < mapLen+map2Start; i++ {
		map2[i] = i
	}
	t.Logf("map1: %v", map1)
	t.Logf("map2: %v", map2)
	testMap := MapDifference(map1, map2)
	t.Logf("difference: %v", testMap)
	for k := range map1 {
		_, notDifferent := map2[k]
		_, isInDiff := testMap[k]
		assert.NotEqual(t, notDifferent, isInDiff, "testMap contained %v", k)
	}
	for k := range map2 {
		_, notDifferent := map1[k]
		_, isInDiff := testMap[k]
		assert.NotEqual(t, notDifferent, isInDiff, "testMap contained %v", k)
	}
}

func TestMapIntersection(t *testing.T) {
	map1, map2 := make(map[int]int), make(map[int]int)
	mapLen := 10
	map1Start, map2Start := 0, mapLen/2
	for i := map1Start; i < mapLen+map1Start; i++ {
		map1[i] = i
	}
	for i := map2Start; i < mapLen+map2Start; i++ {
		map2[i] = i
	}
	t.Logf("map1: %v", map1)
	t.Logf("map2: %v", map2)
	testMap := MapIntersection(map1, map2)
	t.Logf("intersection: %v", testMap)
	for k := range map1 {
		_, isIntersect := map2[k]
		_, isInIntersect := testMap[k]
		assert.Equal(t, isIntersect, isInIntersect, "testMap contained %v", k)
	}
	for k := range map2 {
		_, isIntersect := map1[k]
		_, isInDiff := testMap[k]
		assert.Equal(t, isIntersect, isInDiff, "testMap contained %v", k)
	}
}

func TestMapUnion(t *testing.T) {
	map1, map2 := make(map[int]int), make(map[int]int)
	mapLen := 10
	map1Start, map2Start := 0, mapLen/2
	for i := map1Start; i < mapLen+map1Start; i++ {
		map1[i] = i
	}
	for i := map2Start; i < mapLen+map2Start; i++ {
		map2[i] = i
	}
	t.Logf("map1: %v", map1)
	t.Logf("map2: %v", map2)
	testMap := MapUnion(map1, map2)
	t.Logf("union: %v", testMap)
	for k := range map1 {
		assert.Contains(t, testMap, k, "union did not contain %v", k)
	}
	for k := range map2 {
		assert.Contains(t, testMap, k, "union did not contain %v", k)
	}
}
