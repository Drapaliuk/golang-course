package increment

import (
	"fmt"
	"strconv"
)

func Increment(binary string) int64 {
	decimal, err := strconv.ParseInt(binary, 2, 64)

	if err != nil {
		fmt.Println(err)
	}

	return decimal + 1
}
