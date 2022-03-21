/* Alta3 Research - Slice length & capacity */

package main

import "fmt"

func main() {

        // Slice - length and capacity
    	mySlice5 := []string{"a", "b", "c", "d"}

    fmt.Println("Contents of mySlice5:", mySlice5)
        fmt.Println("Length of mySlice5:",   len(mySlice5))    // returns length
    	fmt.Println("Capacity of mySlice5:", cap(mySlice5))    // returns capacity
    
    mySlice6 := make([]bool, 2, 4)    // the length (window) is set to 2
                      // the capacity (size of array backing Slice) is 4

    fmt.Println("Contents of mySlice6:", mySlice6)					
        fmt.Println("Length of mySlice6:",   len(mySlice6))    // returns length
    	fmt.Println("Capacity of mySlice6:", cap(mySlice6))    // returns capacity
    
}

