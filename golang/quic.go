func ReadVarint(data []byte) (uint64, int) {
	var n int
	var x uint64
	for shift := uint(0); ; shift += 7 {
		if n >= len(data) {
			return 0, 0
		}
		b := data[n]
		n++
		x |= uint64(b&0x7F) << shift
		if b < 0x80 {
			break
		}
	}
	return x, n
}

func EncodePacketNumber(data []byte, p *PacketNumber) int {
	if p.IsLeastUnacked() {
		data[0] = 0
		return 1
	}
	if p.IsAckEliciting() {
		data[0] = 1
	}
	binary.LittleEndian.PutUint64(data[1:], p.Value)
	return 1 + 8
}

func ReadPacketNumber(data []byte) (PacketNumber, int) {
	if data[0] == 0 {
		return PacketNumber(0), 1
	}
	return PacketNumber(binary.LittleEndian.Uint64(data[1:])), 1 + 8
}

func DecodePacketNumber(data []byte) (PacketNumber, int) {
	if data[0] == 0 {
		return PacketNumber(0), 1
	}
	return PacketNumber(binary.LittleEndian.Uint64(data[1:])), 1 + 8
}