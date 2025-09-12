package sss

func gfAdd(x, y byte) byte {
	return x ^ y
}

func gfMul(x, y byte) byte {
	var ret byte
	for i := 0; i < 8; i++ {
		m := -(y & 1) & 0xff
		ret ^= x & m
		hi := x & 0x80
		x <<= 1
		x &= 0xff
		rm := (-(hi >> 7)) & 0xff
		x ^= (0x1b & rm)
		y >>= 1
	}
	return ret
}

func gfInv(y byte) byte {
	if y == 0 {
		panic("division by zero")
	}
	y2 := gfMul(y, y)
	y4 := gfMul(y2, y2)
	y8 := gfMul(y4, y4)
	y16 := gfMul(y8, y8)
	y32 := gfMul(y16, y16)
	y64 := gfMul(y32, y32)
	y128 := gfMul(y64, y64)

	ret := y128
	ret = gfMul(ret, y64)
	ret = gfMul(ret, y32)
	ret = gfMul(ret, y16)
	ret = gfMul(ret, y8)
	ret = gfMul(ret, y4)
	ret = gfMul(ret, y2)
	return ret
}

func gfDiv(x, y byte) byte {
	return gfMul(x, gfInv(y))
}
