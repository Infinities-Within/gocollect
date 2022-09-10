package sets

// MapDifference returns the difference between two maps by creating a new
// map which contains all entries of map1 not present in map2, and all entries of
// map2 not present in map1.
func MapDifference[K comparable, V any](map1, map2 map[K]V) map[K]V {
	result := make(map[K]V)

	for k, v := range map1 {
		if _, ok := map2[k]; !ok {
			result[k] = v
		}
	}

	for k, v := range map2 {
		if _, ok := map1[k]; !ok {
			result[k] = v
		}
	}

	return result
}

// MapUnion returns a new map containing all entries of map1 as well as all entries of map2.
//
// If an entry exists in both maps, the value from map1 is used.
func MapUnion[K comparable, V any](map1, map2 map[K]V) map[K]V {
	result := make(map[K]V)

	// Load all items from map2
	for k, v := range map2 {
		result[k] = v
	}

	// Now load all items from map1, so that any values also present in map2
	// will be overwritten with values from map1, thus fulfilling our contract
	for k, v := range map1 {
		result[k] = v
	}

	return result
}

// MapIntersection returns the intersection between two maps by creating a new map which contains all entries
// which are present in BOTH map1 and map2.
//
// Since type parameter V is not guaranteed to be comparable, the values in the returned map will be the values
// present in map1.
func MapIntersection[K comparable, V any](map1, map2 map[K]V) map[K]V {
	result := make(map[K]V)

	for k, v := range map1 {
		if _, ok := map2[k]; ok {
			result[k] = v
		}
	}

	return result
}
