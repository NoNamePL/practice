package main

import "fmt"

type StringIntMap struct {
	hash map[string]int
}

func (s *StringIntMap) Add(key string, value int) {
	s.hash[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.hash, key)
}

func (s *StringIntMap) Copy() map[string]int {
	res := make(map[string]int, len(s.hash))

	for key, val := range s.hash {
		res[key] = val
	}

	return res
}

func (s *StringIntMap) Exists(key string) bool {
	if _, ok := s.hash[key]; ok {
		return ok
	}
	return false
}

func (s *StringIntMap) Get(key string) (int, bool) {
	val, ok := s.hash[key]
	return val, ok
}

func main() {
	myMap := StringIntMap{hash: map[string]int{}}

	myMap.Add("1", 1)
	fmt.Println(myMap.Exists("1"))
	myNewMap := myMap.Copy()
	fmt.Println(myNewMap["1"])
	fmt.Println(myMap.Get("1"))
	myMap.Remove("1")
	fmt.Println(myMap.Exists("1"))
}
