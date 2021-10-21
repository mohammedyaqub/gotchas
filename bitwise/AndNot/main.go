//&^ (AND NOT): This is a bit clear operator.
package main

import "fmt"

func main() {  
    var a uint8 = 0x4//0000 0100
    var b uint8 = 0x02//0000 0010
    fmt.Printf("%08b [A]\n",a)
    fmt.Printf("%08b [B]\n",b)

    fmt.Printf("%08b (NOT B)\n",^b)
    fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n",b,0xff,b ^ 0xff)

    fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n",a,b,a ^ b)
    fmt.Printf("%08b & %08b = %08b [A AND B]\n",a,b,a & b)
    fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n",a,b,a &^ b)
    fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n",a,b,a & (^b))
	//output
// 	00000100 [A]
// 00000010 [B]
// 11111101 (NOT B)
// 00000010 ^ 11111111 = 11111101 [B XOR 0xff]
// 00000100 ^ 00000010 = 00000110 [A XOR B]
// 00000100 & 00000010 = 00000000 [A AND B]
// 00000100 &^00000010 = 00000100 [A 'AND NOT' B]
// 00000100&(^00000010)= 00000100 [A AND (NOT B)]
}