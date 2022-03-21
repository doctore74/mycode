/* Alta3 Research - Slice and copy function */

package main

import "fmt"

func main() {

        // Copy from one slice to another
        var s = make([]int, 3)
        n := copy(s, []int{0, 1, 2, 3}) // n == 3, s == []int{0, 1, 2}
        fmt.Println(s)                  // [0 1 2]
        fmt.Println(n)                  // 3

        // Copy from a slice to itself
        d := []int{0, 1, 2}
        g := copy(d, d[1:])     // g == 2, d == []int{1, 2, 2}
        fmt.Println(d)          // [1 2 2]
        fmt.Println(g)          // 2
        
}

