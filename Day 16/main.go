package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type packet struct {
	Version    uint64
	Type       uint64
	Subpackets []*packet
	Value      uint64
}

func main() {
	start := time.Now()
	f, err := os.Open("Day16Input.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var hexString string
	for scanner.Scan() {
		hexString = scanner.Text()
	}
	totalVersionSum, value := partOne(hexString)
	fmt.Printf("Total Sum of Versions Part One: %d \n", totalVersionSum)
	fmt.Printf("Value of Packet: %d \n", value)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(hexString string) (uint64, uint64) {
	binaryString := hexToBinaryString(hexString)
	packetList, versionSum := getPackets(binaryString)
	return versionSum, packetList[0].Value
}

func getPackets(binaryString string) ([]*packet, uint64) {
	var versionSum uint64
	var packetList []*packet
	for i := 0; i < len(binaryString); i++ {
		if len(binaryString[i:]) < 7 {
			break
		}
		packet := packet{}
		packet.Version, _ = strconv.ParseUint(binaryString[i:i+3], 2, 64)
		i += 3
		versionSum += packet.Version
		packet.Type, _ = strconv.ParseUint(binaryString[i:i+3], 2, 64)
		i += 3
		if fmt.Sprint(packet.Type) == "4" {
			var value string
			for string(binaryString[i]) == "1" {
				value += string(binaryString[i+1 : i+5])
				i += 5
			}
			value += string(binaryString[i+1 : i+5])
			i += 4
			packet.Value, _ = strconv.ParseUint(value, 2, 64)
			packetList = append(packetList, &packet)
		} else {
			if string(binaryString[i]) == "0" && len(binaryString[i+1:]) > 15 {
				length, _ := strconv.ParseUint(binaryString[i+1:i+16], 2, 64)
				i += 16
				packet.Subpackets, versionSum = getSubPacketsByLength(binaryString[i:i+int(length)], versionSum)
				packet.Value = getOperatorPacketValue(packet.Type, packet)
				i += int(length) - 1
				packetList = append(packetList, &packet)
			} else if string(binaryString[i]) == "1" && len(binaryString[i+1:]) > 11 {
				subPackets, _ := strconv.ParseUint(binaryString[i+1:i+12], 2, 64)
				i += 12
				var addToi int
				packet.Subpackets, addToi, versionSum = getSubPacketsByNumber(binaryString[i:], int(subPackets), versionSum)
				packet.Value = getOperatorPacketValue(packet.Type, packet)
				i += addToi
				packetList = append(packetList, &packet)
			}
		}
	}
	return packetList, versionSum
}

func getSubPacketsByNumber(binaryString string, subPackets int, versionSum uint64) ([]*packet, int, uint64) {
	var packetList []*packet
	var i int
	var j int
	for i, j = 0, 0; j < subPackets; i++ {
		packet := packet{}
		packet.Version, _ = strconv.ParseUint(binaryString[i:i+3], 2, 64)
		i += 3
		versionSum += packet.Version
		packet.Type, _ = strconv.ParseUint(binaryString[i:i+3], 2, 64)
		i += 3
		if fmt.Sprint(packet.Type) == "4" {
			var value string
			for string(binaryString[i]) == "1" {
				value += string(binaryString[i+1 : i+5])
				i += 5
			}
			value += string(binaryString[i+1 : i+5])
			i += 4
			packet.Value, _ = strconv.ParseUint(value, 2, 64)
			packetList = append(packetList, &packet)
			j++
		} else {
			if string(binaryString[i]) == "0" && len(binaryString[i+1:]) > 15 {
				length, _ := strconv.ParseUint(binaryString[i+1:i+16], 2, 64)
				i += 16
				packet.Subpackets, versionSum = getSubPacketsByLength(binaryString[i:i+int(length)], versionSum)
				i += int(length) - 1
				packet.Value = getOperatorPacketValue(packet.Type, packet)
				packetList = append(packetList, &packet)
				j++
			} else if string(binaryString[i]) == "1" && len(binaryString[i+1:]) > 11 {
				subPackets, _ := strconv.ParseUint(binaryString[i+1:i+12], 2, 64)
				i += 12
				var addToi int
				packet.Subpackets, addToi, versionSum = getSubPacketsByNumber(binaryString[i:], int(subPackets), versionSum)
				packet.Value = getOperatorPacketValue(packet.Type, packet)
				i += addToi - 1
				packetList = append(packetList, &packet)
				j++
			}
		}
	}
	return packetList, i, versionSum
}

func getSubPacketsByLength(binaryString string, versionSum uint64) ([]*packet, uint64) {
	var packetList []*packet
	for i := 0; i < len(binaryString); i++ {
		packet := packet{}
		packet.Version, _ = strconv.ParseUint(binaryString[i:i+3], 2, 64)
		i += 3
		versionSum += packet.Version
		packet.Type, _ = strconv.ParseUint(binaryString[i:i+3], 2, 64)
		i += 3
		if fmt.Sprint(packet.Type) == "4" {
			var value string
			for string(binaryString[i]) == "1" {
				value += string(binaryString[i+1 : i+5])
				i += 5
			}
			value += string(binaryString[i+1 : i+5])
			i += 4
			packet.Value, _ = strconv.ParseUint(value, 2, 64)
			packetList = append(packetList, &packet)
		} else {
			if string(binaryString[i]) == "0" && len(binaryString[i+1:]) > 15 {
				length, _ := strconv.ParseUint(binaryString[i+1:i+16], 2, 64)
				i += 16
				packet.Subpackets, versionSum = getSubPacketsByLength(binaryString[i:i+int(length)], versionSum)
				i += int(length) - 1
				packet.Value = getOperatorPacketValue(packet.Type, packet)
				packetList = append(packetList, &packet)
			} else if string(binaryString[i]) == "1" && len(binaryString[i+1:]) > 11 {
				subPackets, _ := strconv.ParseUint(binaryString[i+1:i+12], 2, 64)
				i += 12
				var addToi int
				packet.Subpackets, addToi, versionSum = getSubPacketsByNumber(binaryString[i:], int(subPackets), versionSum)
				packet.Value = getOperatorPacketValue(packet.Type, packet)
				i += addToi - 1
				packetList = append(packetList, &packet)
			}
		}
	}
	return packetList, versionSum
}

func getOperatorPacketValue(packetType uint64, packet packet) uint64 {
	var value uint64
	switch packetType {
	case uint64(0):
		for _, val := range packet.Subpackets {
			value += val.Value
		}
	case uint64(1):
		value++
		for _, val := range packet.Subpackets {
			value *= val.Value
		}
	case uint64(2):
		value = math.MaxInt64
		for _, val := range packet.Subpackets {
			if val.Value < value {
				value = val.Value
			}
		}
	case uint64(3):
		value = 0
		for _, val := range packet.Subpackets {
			if val.Value > value {
				value = val.Value
			}
		}
	case uint64(5):
		if packet.Subpackets[0].Value > packet.Subpackets[1].Value {
			value = 1
		} else {
			value = 0
		}
	case uint64(6):
		if packet.Subpackets[0].Value < packet.Subpackets[1].Value {
			value = 1
		} else {
			value = 0
		}
	case uint64(7):
		if packet.Subpackets[0].Value == packet.Subpackets[1].Value {
			value = 1
		} else {
			value = 0
		}
	}
	return value
}

func hexToBinaryString(hexString string) string {
	var binaryString string
	for _, v := range hexString {
		switch string(v) {
		case "0":
			binaryString += "0000"
		case "1":
			binaryString += "0001"
		case "2":
			binaryString += "0010"
		case "3":
			binaryString += "0011"
		case "4":
			binaryString += "0100"
		case "5":
			binaryString += "0101"
		case "6":
			binaryString += "0110"
		case "7":
			binaryString += "0111"
		case "8":
			binaryString += "1000"
		case "9":
			binaryString += "1001"
		case "A":
			binaryString += "1010"
		case "B":
			binaryString += "1011"
		case "C":
			binaryString += "1100"
		case "D":
			binaryString += "1101"
		case "E":
			binaryString += "1110"
		case "F":
			binaryString += "1111"
		}
	}
	return binaryString
}

func partTwo() {

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
