package woutils

func CartesianToIsometric(x, y int32) (int32, int32) {
	x = x / 2
	return x - y, (x + y) / 2
}
