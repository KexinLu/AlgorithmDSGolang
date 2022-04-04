# DataStructure_GO

## Getting started
Clone this project and put it under your gopath
This project employs gomod for package management

## File Structure
For every topic the file structure will look like this:

```
data_structure_name
    reading
        file
    implementation
        file
        test
    leetcode
        lc001
            file
            test
        lc002
            file
            test
        lc212
            file
            test
```

## Naming Convention

under the algorithm path, you will find folders of problemsets grouped by datastructure or other type tags
the way to name these problems should be

codeplatformID_problemname_freq[HML]_[easy|med|hard]
eg. lc001_twosum_freqH_easy

note the problemname should be one word

Following this naming convention is important to enable us to use some bash scripting to filter through our problems

## Purpose of data structure?

- organizing data into a format which enables some coding pattern to store/retrieve data
- helping us write cleaner code

## Abstract Data Types

- ADT is "interface" of data structures
- ADT only provides high level contracts but no specific details
eg.
     List
     Queue
     Map


## Big O

- Big O is an estimation, we should only keep the highest order item.
eg.

O(2n) ~> O(n)
O(n^2 + 2n) ~> O(n^2)

- Binary-search-like operations would turn an O(n) operation to O(logN). (In cs log is base 2 not 10)
- For comparasion search, the lower limit of O is O(nlogn), to have a better performance you need to use counting sort or radix sort or other sorting method


## Data Structures in this repo

* Array
    - Array(Static Array) [x]
    - Slice(Dynamic Array) [x]
    - String [x]
    - Suffix Array [x]
* List
    - Linked List [x]
    - Doubly Linked List [x]
* HashMap [x]
* Stack [x]
* Queue
    - simple queue [x]
    - PriorityQueue [x] (Only the Heap Version)
* Graph
    - Unidirectional Graph []
    - Bidirectional Graph []
    - BFS []
    - DFS []
    - MST(minimum spanning tree) []
    - 2d matrix(2d arrays) []
    - Tree(https://www.youtube.com/watch?v=xVka6z1hu-I)
        - Heap [x]
        - Trie [x]
        - Binary Tree [x]
        - Binary Search Tree [x]
        - Fenwick Tree (Binary Indexd Tree) []
        - Red-black Tree []
        - AVL Tree []
* Union Find []
