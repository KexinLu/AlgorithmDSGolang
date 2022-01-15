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
    - Array(Static Array) []
    - Slice(Dynamic Array) []
    - String []
    - Suffix Array []
* List 
    - Linked List []
    - Doubly Linked List []
* HashMap []
* Stack []
* Queue
    - simple queue []
    - PriorityQueue []
* Graph
    - Unidirectional Graph []
    - Bidirectional Graph []
    - BFS []
    - DFS []
    - MST(minimum spanning tree) []
    - 2d matrix(2d arrays) []
    - Tree(https://www.youtube.com/watch?v=xVka6z1hu-I)
        - Heap []
        - Trie []
        - Binary Tree []
        - Binary Search Tree []
        - Fenwick Tree (Binary Indexd Tree) []
        - Red-black Tree []
        - AVL Tree []
* Union Find []
