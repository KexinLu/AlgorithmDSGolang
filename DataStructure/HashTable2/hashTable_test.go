package hashTable2

import (
	"testing"
)

func TestCanCreateHashTable(t *testing.T) {
	ht := NewHashTable()
	if ht.size != 0 {
		t.Error("hash table size is wrong")
	}

	if len(ht.data) != 256 {
		t.Error("hash table size is wrong")
	}

}

//func TestGetHash(t *testing.T) {
//	if 11 != getHash("abc") {
//		t.Error("wrong hash value")
//	}
//	if 189 != getHash("abcd") {
//		t.Errorf("wrong hash value %v", getHash("abcd"))
//	}
//}

func TestCanPut(t *testing.T) {
	ht := NewHashTable()
	idx := ht.Put("hello", 15)
	item := ht.data[idx]
	if item.key != "hello" {
		t.Errorf("item key %v is different from expected %v", item.key, "hello")
	}

	if item.value != 15 {
		t.Errorf("item value %v is different from expected %v", item.value, 15)

	}

	if ht.size != 1 {
		t.Error("hash table size should be 1")
	}
}

func TestPut(t *testing.T) {
	ht := NewHashTable()
	idx := ht.Put("hello", 15)

	t.Run("Can update existing value", func(t *testing.T) {
		idx = ht.Put("hello", 18)
		helpCheckItem(t, ht, idx, "hello", 18)
	})

	t.Run("Can insert multiple value", func(t *testing.T) {
		ht.hashFn = func(key string) uint8 {
			if key == "another" {
				return 20
			}
			return idx
		}
		idx = ht.Put("no", 19)
		newID := ht.Put("another", 25)

		helpCheckItem(t, ht, idx, "hello", 18)
		helpCheckItem(t, ht, idx, "no", 19)
		helpCheckItem(t, ht, newID, "another", 25)
	})

}

func TestGet(t *testing.T) {
	ht := NewHashTable()
	idx := ht.Put("hello", 15)

	t.Run("Should return false when key is not in the table", func(t *testing.T) {
		val, exist := ht.Get("no")
		if val != -1 || exist != false {
			t.Errorf("should have returned %v and %v but got %v and %v", -1, false, val, exist)
		}
	})

	t.Run("Can get first node", func(t *testing.T) {
		val, exist := ht.Get("hello")
		if val != 15 || exist != true {
			t.Errorf("should have returned %v and %v but got %v and %v", 15, true, val, exist)
		}
	})

	t.Run("Can get node in linked list", func(t *testing.T) {
		ht.hashFn = func(key string) uint8 {
			if key == "another" {
				return 20
			}
			return idx
		}
		idx = ht.Put("no", 19)
		ht.Put("another", 25)

		testCases := []struct {
			inpKey  string
			expVal  int
			expBool bool
		}{
			{"hello", 15, true},
			{"another", 25, true},
			{"no", 19, true},
		}
		for _, tc := range testCases {
			val, exist := ht.Get(tc.inpKey)
			if val != tc.expVal || exist != tc.expBool {
				t.Errorf("should have returned %v and %v but got %v and %v", tc.expVal, tc.expBool, val, exist)
			}
		}
	})
}

func helpCheckItem(t *testing.T, ht *HashTable, idx uint8, key string, value int) {
	if ht.size == 0 {
		t.Error("got an empty hash table")
		return
	}
	item := ht.data[idx]
	if item == nil {
		t.Error("item does not exist for given index")
		return
	}
	cur := item
	for cur != nil {
		if cur.key == key {
			if cur.value != value {
				t.Errorf("item value %v is different from expected %v", cur.value, value)
				return
			} else {
				return
			}
			break
		}
		cur = cur.next
	}

	t.Errorf("item with key %s does not exist", key)
}
