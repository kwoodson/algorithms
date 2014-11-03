package main

import "fmt"
import "math/rand"

func print_array(array []int) {
   fmt.Printf("array = ")
   for i := 0; i < len(array); i++ {
      fmt.Printf("%d, ", array[i])
   }
   fmt.Printf("\n")
}


func quick(array []int) []int {
   if len(array) <= 1 {
      return array
   }

   left, right := 0, len(array) - 1

   pivotIndex := rand.Int() % len(array)
   fmt.Printf("Pivot=[%d]\n", pivotIndex)

   print_array(array)
   array[pivotIndex], array[right] = array[right], array[pivotIndex]
   fmt.Printf("After pivot\n")
   print_array(array)

   for i := range array {
      if array[i] < array[right] {
         array[i], array[left] = array[left], array[i]
         left++
      }
   }

   fmt.Printf("finished\n")
   print_array(array)

   fmt.Printf("swap left/right\n")
   array[left], array[right] = array[right],array[left]
   print_array(array)

   quick(array[:left])
   quick(array[left+1:])

   return array
}

func main() {
   x := []int {1, 8, 7, 5, 2, 4, 10, 3, 9, 6}

   var result = quick(x)
   print_array(result)
}
