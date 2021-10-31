func bytetoqrcode(data []byte) *qrcode.QRCode {
	qr, _ := qrcode.New(string(data), qrcode.Medium)
	return qr
}

func BoolsToBytes(bools []bool) []byte {
	var bytes []byte
	for i := 0; i < len(bools); i += 8 {
		bytes = append(bytes, Bits2Byte(bools[i:i+8]))
	}
	return bytes


func BitsToBytes(bits []bool) []byte {
	var bytes []byte
	var tmp byte
	for _, b := range bits {
		if b {
			tmp |= 1
		} else {
			tmp |= 0
		}
		if len(bytes) == 0 {
			bytes = append(bytes, tmp)
		} else {
			bytes[len(bytes)-1] |= tmp << 1
			bytes = append(bytes, tmp>>7)
		}
		tmp = tmp << 1
	}
	return bytes
}

func Bits2Int(bits []bool) int {
	var num int
	for i, b := range bits {
		if b {
			num |= 1 << uint(i)
		}
	}
	return num
}

func Bits2Byte(bits []bool) byte {
	var num byte
	for i, b := range bits {
		if b {
			num |= 1 << uint(i)
		}
	}
	return num
}

func Bytes2String(bytes []byte) string {
	var str string
	for _, b := range bytes {
		str += fmt.Sprintf("%02x", b)
	}
	return str
}

func Bools2String(bools []bool) string {
	return Bytes2String(BoolsToBytes(bools))
}

func Rectangle(group []Pos) (minX, maxX, minY, maxY int) {
	minX, maxX, minY, maxY = group[0].X, group[0].X, group[0].Y, group[0].Y
	for _, p := range group {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y	
		}
	}
	return	
}

type Pos struct {
	X int
	Y int
}



