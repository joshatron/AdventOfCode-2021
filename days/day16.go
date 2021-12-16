package days

import (
	"fmt"
	"math"
	"strconv"

	"joshatron.io/aoc2021/input"
)

func Day16Puzzle1() string {
	_, versionTotal := parsePackets(input.ReadDayInput("16"))

	return fmt.Sprint(versionTotal)
}

func parsePackets(hex string) (Packet, int64) {
	binary := hexToBinaryString(hex)
	return parsePacketRecursive(binary)
}

func parsePacketRecursive(binary string) (Packet, int64) {
	version, _ := strconv.ParseInt(binary[:3], 2, 64)
	typeId, _ := strconv.ParseInt(binary[3:6], 2, 64)
	var totalLength int64
	totalLength = 6

	remaining := binary[6:]
	if typeId == 4 {
		combined := ""
		for remaining[0] == '1' {
			combined += remaining[1:5]
			remaining = remaining[5:]
			totalLength += 5
		}
		combined += remaining[1:5]
		totalLength += 5
		num, _ := strconv.ParseInt(combined, 2, 64)
		return Packet{version, typeId, num, totalLength, []*Packet{}}, version
	} else {
		if remaining[0] == '0' {
			remaining = remaining[1:]
			bitsToParse, _ := strconv.ParseInt(remaining[:15], 2, 64)
			remaining = remaining[15:]
			versionTotal := version
			packet := Packet{version, typeId, 0, bitsToParse + 22, []*Packet{}}
			for bitsToParse > 0 {
				innerPacket, total := parsePacketRecursive(remaining)
				packet.children = append(packet.children, &innerPacket)
				versionTotal += total
				bitsToParse -= innerPacket.length
				remaining = remaining[innerPacket.length:]
			}
			return packet, versionTotal
		} else {
			remaining = remaining[1:]
			packetsToParse, _ := strconv.ParseInt(remaining[:11], 2, 64)
			remaining = remaining[11:]
			versionTotal := version
			packet := Packet{version, typeId, 0, 18, []*Packet{}}
			for packetsToParse > 0 {
				innerPacket, total := parsePacketRecursive(remaining)
				packet.children = append(packet.children, &innerPacket)
				versionTotal += total
				packet.length += innerPacket.length
				packetsToParse--
				remaining = remaining[innerPacket.length:]
			}
			return packet, versionTotal
		}
	}
}

func hexToBinaryString(hex string) string {
	binary := ""
	for _, r := range hex {
		switch r {
		case '0':
			binary += "0000"
		case '1':
			binary += "0001"
		case '2':
			binary += "0010"
		case '3':
			binary += "0011"
		case '4':
			binary += "0100"
		case '5':
			binary += "0101"
		case '6':
			binary += "0110"
		case '7':
			binary += "0111"
		case '8':
			binary += "1000"
		case '9':
			binary += "1001"
		case 'A':
			binary += "1010"
		case 'B':
			binary += "1011"
		case 'C':
			binary += "1100"
		case 'D':
			binary += "1101"
		case 'E':
			binary += "1110"
		case 'F':
			binary += "1111"
		}
	}

	return binary
}

type Packet struct {
	version  int64
	typeId   int64
	literal  int64
	length   int64
	children []*Packet
}

func Day16Puzzle2() string {
	packet, _ := parsePackets(input.ReadDayInput("16"))
	return fmt.Sprint(packetValue(packet))
}

func packetValue(packet Packet) int64 {
	switch packet.typeId {
	case 0:
		var sum int64
		for _, p := range packet.children {
			sum += packetValue(*p)
		}
		return sum
	case 1:
		var product int64
		product = 1
		for _, p := range packet.children {
			product *= packetValue(*p)
		}
		return product
	case 2:
		var min int64
		min = math.MaxInt
		for _, p := range packet.children {
			temp := packetValue(*p)
			if temp < min {
				min = temp
			}
		}
		return min
	case 3:
		var max int64
		max = math.MinInt
		for _, p := range packet.children {
			temp := packetValue(*p)
			if temp > max {
				max = temp
			}
		}
		return max
	case 4:
		return packet.literal
	case 5:
		if packetValue(*packet.children[0]) > packetValue(*packet.children[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		if packetValue(*packet.children[0]) < packetValue(*packet.children[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		if packetValue(*packet.children[0]) == packetValue(*packet.children[1]) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}
