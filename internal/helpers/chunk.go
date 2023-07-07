package helpers

func ChunkByMaxLen[T any](arr []T, length int) (result [][]T) {
	start := 0
	for {
		if start+length < len(arr) {
			result = append(result, arr[start:start+length])
			start += length
		} else {
			if len(arr)%length > 0 {
				result = append(result, arr[start:])
			}
			break
		}
	}
	return
}
