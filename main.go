package main

import (
	"fmt"
	"math"
)

func main() {
	//var i uint = 1 << 63
	//fmt.Printf("i : %064b \n",i)
	//fmt.Printf("i : %064b \n",math.MaxInt)
	////k := 64 - bits.LeadingZeros64(i)
	////fmt.Printf("k: %d, pow2:%064b \n",k ,1 << uint(k))
	//
	//
	//fmt.Printf("LeadingZeros64(%064b) = %d\n", 0, bits.LeadingZeros64(0))
	//s := []int64{1,3,3,4,2,2,3}
	//fmt.Println(len(s)," , ",cap(s))
	//s = s[:2]
	//fmt.Printf("%v ",s)
	//fmt.Println(len(s)," , ",cap(s))
	fmt.Println(math.Ceil(0.0001))
	fmt.Println(uint32(math.MaxUint64 >> 32))
	// 15625
}
