package HashTable

import (
	"fmt"
	"hash/fnv"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	s := "abc"

	hash := fnv.New64a()
	hash.Write([]byte(s))

	hashVal := hash.Sum64()
	expected := uint8(hashVal % 256)
	rst := getHash(s)

	if rst != expected {
		t.Errorf("The hash Value %v is different from the expected hash value %v", rst, expected)
	}
}

func TestPutItem(t *testing.T) {
	key := "k1"
	val := "myValue"
	valReplace := "value2"

	key2 := "k2"
	val2 := "val2"
	ht := NewHashTable()

	ht.PutItem(key, val)
	ht.PutItem(key2, val2)

	tests := []struct {
		k, v     string
		expected string
	}{
		{key, val, val},
		{key, valReplace, valReplace},
		{key2, val2, val2},
	}

	for n, tt := range tests {
		ht.PutItem(tt.k, tt.v)
		if rst, _ := ht.GetItem(tt.k); rst != tt.expected {
			t.Errorf("test case %v failed, result %s, expected %s", n+1, rst, tt.expected)
		}
	}
}

func TestPutItemSameHash(t *testing.T) {
	t.Run("Test Putting Items with the same hash into the table and able to get", func(t *testing.T) {
		ht := NewHashTable()
		forceHashFn(ht, func(key string) uint8 {
			return 8
		})

		tests := []struct {
			k, v     string
			expected string
		}{
			{"key1", "val1", "val1"},
			{"key2", "val2", "val2"},
			{"key3", "val3", "val3"},
		}

		for _, tc := range tests {
			ht.PutItem(tc.k, tc.v)
			if v, err := ht.GetItem(tc.k); v != tc.expected || err != nil {
				t.Errorf("GetItem returned %s which is not expected %s, with error %s", v, tc.expected, err)
			}
		}
	})

}

func TestDeleteItem(t *testing.T) {
	t.Run("Test Normal Delete", func(t *testing.T) {
		ht := NewHashTable()

		ht.PutItem("key1", "val1")
		ht.PutItem("key2", "val2")

		if err := ht.DeleteItem("key1"); err != nil {
			t.Errorf("Failed to delete item with key %s", "key1")
		}

		if _, err := ht.GetItem("key1"); err == nil {
			t.Error("Get Item should throw error")
		}
	})

	t.Run("Test Delete on the same linked list", func(t *testing.T) {
		ht := NewHashTable()
		ht.PutItem("key1", "val1")
		ht.PutItem("key2", "val2")

		fmt.Println(ht.GetItem("key1"))

		if err := ht.DeleteItem("key1"); err != nil {
			t.Errorf("Failed to delete item with key %s", "key1")
		}

		if _, err := ht.GetItem("key1"); err == nil {
			t.Error("Get Item should throw error")
		}

	})

}
