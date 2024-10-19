package woutils

func ListIntToListInt32(listInt []int) []int32 {
	listInt32 := make([]int32, len(listInt))
	for i, v := range listInt {
		listInt32[i] = int32(v)
	}
	return listInt32
}
