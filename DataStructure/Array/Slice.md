# Slice

## Dynamic Array
- length is automatically adjusted
- create
```
mySlice := []string{"a", "b", "c", "d"}
```
or
```
mySlice := make([]string, 15, 20)
```
- length & capacity
- zero value of slice is nil
- len, cap will return 0
- sub slice
```
a := []byte{'a','b','c','d','e'}
b := a[:2]
fmt.Println(b) -> a b
c := a[2:]
fmt.Println(b) -> c d e
```


- Slice data structure
```
type Slice<T> struct {
  firstElement pointerToThe element
  len int      -> number of elements referred to by this slice
  cap int      -> capacity of the underlying array
}
```
- To grow the Capacity of a slice, you need to copy
- ^ can be handled with the built in append function