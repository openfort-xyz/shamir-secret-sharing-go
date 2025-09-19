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

	yk := make([]byte, 8)
	yk[0] = y
	for i := 1; i < 8; i++ {
		yk[i] = gfMul(yk[i-1], yk[i-1])
	}

	ret := yk[7]
	for i := 6; i > 0; i-- {
		ret = gfMul(ret, yk[i])
	}
	return ret
}

func gfDiv(x, y byte) byte {
	return gfMul(x, gfInv(y))
}
