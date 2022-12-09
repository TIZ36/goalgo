package algorithm

func SquareIsWhite(coordinates string) bool {
	heng := []uint8{1, 2, 3, 4, 5, 6, 7, 8}
	v := heng[coordinates[0]-'a'] + (coordinates[1] - '0')
	return !(v%2 == 0)
}
