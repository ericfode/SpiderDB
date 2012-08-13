package spiderDB

import "strconv"

func IntToBytes(in int) []byte {
	str := strconv.Itoa(in)
	return []byte(str)
}

func BytesToInt(in []byte) int {
	out, _ := strconv.Atoi(string(in))
	return out
}

func StringToInt(in string) int {
	out, _ := strconv.Atoi(in)
	return out
}
