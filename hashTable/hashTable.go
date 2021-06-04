package main

import "fmt"

type hashTable struct {
	index map[int]int
	cap   int
}

func main() {
	hash := createHash(12)
	fmt.Println(hash)
	hash.insert(1, 10)
	hash.insert(12, 20)
	hash.insert(13, 30)
	hash.insert(11, 40)
	fmt.Println(hash)
	hash.remove(1)
	fmt.Println(hash.get(12))
	fmt.Println(hash.size())
}

func createHash(cap int) hashTable {
	new := hashTable{cap: cap, index: make(map[int]int)}
	new.checkCapacity()
	return new
}

func (h *hashTable) insert(key int, data int) {
	if len(h.index) < h.cap {
		hashValue := hashFn(key)
		h.index[hashValue] = data
		return
	}
	fmt.Println("The hash table is full")
}

func (h *hashTable) remove(key int) {
	hashValue := hashFn(key)
	delete(h.index, hashValue)
}

func (h *hashTable) get(key int) int {
	hashValue := hashFn(key)
	return h.index[hashValue]
}

func (h *hashTable) size() int {
	return len(h.index)
}

func hashFn(key int) int {
	return 2*key + 1
}

//to get a prime number for capaity
func (h *hashTable) checkCapacity() {
	if h.cap%2 == 0 {
		h.cap++
	}
	for !checkPrime(h.cap) {
		h.cap += 2
	}
}

//to check if the capacity is a prime number or not
func checkPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
