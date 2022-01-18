package HashTable

import (
	"fmt"
	"hash/fnv"
)

// how many different keys are we expecting?
// if in hundreds, use uint8
// if much bigger use 65536 which is unit16

type HashTableItem struct {
	key   string
	value string
	next  *HashTableItem
}

type HashTable struct {
	data   [256]*HashTableItem
	hashFn func(string) uint8
}

type IHashTable interface {
	PutItem(key string, value string) string
	GetItem(key string) (string, error)
	DeleteItem(key string) error
}

func NewHashTable() *HashTable {
	var tbData [256]*HashTableItem

	return &HashTable{data: tbData, hashFn: getHash}
}

func (t *HashTable) PutItem(key string, value string) string {
	hash := t.hashFn(key)
	cur := t.data[hash]
	if cur == nil {
		t.data[hash] = &HashTableItem{key: key, value: value}
		return value
	}

	if cur.key == key {
		cur.value = value
		return value
	}

	for cur.next != nil {
		cur = cur.next
		if cur.key == key {
			cur.value = value
		}
	}

	cur.next = &HashTableItem{
		key:   key,
		value: value,
	}

	return value
}

func (t *HashTable) ForceHashFn(fn func(string) uint8) {
	t.hashFn = fn
}

func (t *HashTable) GetItem(key string) (string, error) {
	hash := t.hashFn(key)
	cur := t.data[hash]
	if cur == nil {
		return "", fmt.Errorf("Hash Table does not contain the item with the key %s", key)
	}

	if cur.key == key {
		return cur.value, nil
	}

	for cur.next != nil {
		cur = cur.next
		if cur.key == key {
			return cur.value, nil
		}
	}

	return "", fmt.Errorf("Hash Table does not contain the item with the key %s", key)
}

func (t *HashTable) DeleteItem(key string) (err error) {
	hash := t.hashFn(key)
	cur := t.data[hash]
	if cur == nil {
		return fmt.Errorf("Hash Table does not contain the Item with key %s", key)
	} else if cur.key == key {
		t.data[hash] = cur.next
		return nil
	}

	for cur.next != nil {
		prev := cur
		cur = cur.next
		if cur.key == key {
			prev.next = cur.next
			return nil
		}
	}

	return fmt.Errorf("Hash Table does not contain the Item with key %s", key)
}

func getHash(key string) uint8 {
	hash := fnv.New64a()
	hash.Write([]byte(key))

	return uint8(hash.Sum64() % 256)
}

func forceHashFn(t *HashTable, fn func(string) uint8) {
	t.hashFn = fn
}
