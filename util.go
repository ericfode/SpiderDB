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

func ByteAAtoStringMap(temp [][]byte) map[string][]byte {

	result := make(map[string][]byte)
	for i := 0; i < len(temp); i += 2 {
		result[string(temp[i])] = temp[i+1]
	}

	return result
}
