package main

import "fmt"

func print_array(array []int) {
  for i := 0; i < len(array); i++ {
    fmt.Printf("array[%d] = %d\n", i, array[i])
  }
}

func merge_sort(array []int) []int {
  //base cases
  if (len(array) == 1) {
    return array
  }

  // merge step
  left  := array[:len(array)/2]
//#fmt.Printf("Left\n")
//#print_array(left)
  right := array[len(array)/2:]
//#fmt.Printf("Right\n")
//#print_array(right)

  left  = merge_sort(left)
  right = merge_sort(right)

  return merge(left, right)
}

func merge(left []int, right []int) []int {

  //var result = make([]int, len(left) + len(right))
  var result = make([]int, 1)
//fmt.Printf("Length of result ( %d )\n", len(result))

  // keep going if either have elements
  for len(left) > 0 || len(right) > 0 {

    // if both have elements see which is >
    if len(left) > 0 && len(right) > 0 {

      // left great than right
      if left[0] > right[0] {
        result = append(result, right[0])
        if len(right) == 0 {
          right = []int {}
        } else {
          right = right[1:]
        }
      } else {
        result = append(result, left[0])
        if len(left) == 0 {
          left = []int {}
        } else {
          left = left[1:]
        }
      }
    } else if len(left) > 0 {
        result = append(result, left[0])
        if len(left) == 0 {
          left = []int {}
        } else {
          left = left[1:]
        }
    } else if len(right) > 0 {
        result = append(result, right[0])
        if len(right) == 0 {
          right = []int {}
        } else {
          right = right[1:]
        }
    }
  }


//#fmt.Printf("Result:\n")
//#print_array(result)
  return result[1:]
}


func main() {
  x := []int {1, 8, 7, 5, 2, 4, 10, 3, 9, 6}

  var result = merge_sort(x)
  for i := 0; i < len(result); i++ {
    fmt.Printf("Array[%d] = %d\n", i, result[i])
  }

}
