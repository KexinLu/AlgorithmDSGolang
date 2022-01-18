# Array
- An array type definition specifies a **length** and an **element type**
- Arrays **do not need to be initialized** explicitly
- zero value of an array == ready-to-use array whose elements are themselves zeroed
    eg. var myArray [5]int
    ```
     for _, elem := range myArray {
         fmt.Println(elem)
     }

     will Print
     0
     0
     0
     0
     0
    ```
- in memory representation of [4]int
  ![image](assets/slice-array.png)
- Go's Array is Value, not pointer. Thus, assign/pass-as-argument would result in a copy of its content with different memory address; thus, passing by pointer is preferred, but you need to consider multi-thread operation
- [pointers are not thread safe, even for goroutine](https://stackoverflow.com/questions/18116224/go-concurrent-access-to-pointers-methods#:~:text=Any%20pointer%20is%20considered%20not,go%20routines%20as%20you%20want.)
- You can think of Array like a struct with numbers from 0 -> n as the keys
  ```
  type [5]string struct {
      0  ""
      1  ""
      2  ""
      3  ""
      4  ""
  }
  ```

- Creating a new array in golang
```
a := [2]string{"abc","def"}
or
a := [...]string{"abc", "def"}
```
- both will result in a [2]string being the type of a
- Note: [...]string{"abc", "def"} and []string{"abc","def"} are different 