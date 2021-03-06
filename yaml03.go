/* RZFeeser | Alta3 Research
   Generating YAML files      */

package main

import (
     "fmt"
     "io/ioutil"
     "log"

     "gopkg.in/yaml.v3"
)

func main() {

     trees := [5]string{"elm", "oak", "maple", "sycamore", "chestnut"}

     data, err := yaml.Marshal(&trees)

     if err != nil {
          log.Fatal(err)
     }

     // does not work - permission denied: err2 := ioutil.WriteFile("trees.yaml", data, 0)
     err2 := ioutil.WriteFile("trees.yaml", data, 0700)

     if err2 != nil {

          log.Fatal(err2)
     }

     fmt.Println("data written")
}

