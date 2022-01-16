package hashTable2

import (
	"hash/fnv"
)

type IhashTable interface {
	Put(key string, val int)
	Get(key string) (int, bool)
	Size() int
	Delete(key string) error
}

type Item struct {
	key   string
	value int
	next  *Item
}

type HashTable struct {
	size   int
	data   [256]*Item
	hashFn func(string) uint8
}

func NewHashTable() *HashTable {
	var tableData [256]*Item
	return &HashTable{
		data:   tableData,
		hashFn: getHash,
	}
}

func getHash(s string) uint8 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return uint8(hash.Sum32() % 256)
}

func (ht *HashTable) Put(key string, val int) uint8 {
	hash := ht.hashFn(key)
	firstNode := ht.data[hash]
	if firstNode == nil {
		ht.data[hash] = &Item{
			key:   key,
			value: val,
		}
		ht.size++
		return hash
	}

	var prevItem *Item
	curItem := firstNode

	for curItem != nil {
		prevItem = curItem
		// update
		if curItem.key == key {
			curItem.value = val
			return hash
		}

		curItem = curItem.next
	}

	prevItem.next = &Item{
		key:   key,
		value: val,
	}
	ht.size++
	return hash

}

func (ht *HashTable) Get(key string) (int, bool) {
	hash := ht.hashFn(key)
	firstNode := ht.data[hash]
	if firstNode == nil {
		return -1, false
	}

	curItem := firstNode
	for curItem != nil {
		if curItem.key == key {
			return curItem.value, true
		}
		curItem = curItem.next
	}

	return -1, false
}
