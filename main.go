package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {

	if ransomNote > magazine {
		return false
	}

	m := make(map[int32]int)

	for i, v := range magazine {
		m[v]++
		m[int32(ransomNote[i])]--
	}
	return false
}
