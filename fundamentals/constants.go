package fundamentals

import "fmt"

const (
	one = 1
	two = 2
)

// successive integer constants
const (
	zeroIota = iota
	oneIota
	twoIotaPlus10 = iota + 10
	lastIotaPlus10
)

// iota reset with every constant block
const (
	zeroIotaNewBlock = iota
	oneIotaNewBlock = iota
)

func Constants() {
	fmt.Println("******* Constants ******")

	fmt.Println(one, two)
	fmt.Println(zeroIota, oneIota, twoIotaPlus10, lastIotaPlus10) // 0 1 12 13
	fmt.Println(zeroIotaNewBlock, oneIotaNewBlock) // 0 1
}
