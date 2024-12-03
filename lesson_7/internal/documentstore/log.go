package documentstore

import (
	"fmt"
	"time"
)

func logAction(action string, name string) {
	fmt.Printf("[%s] %s: %s\n", time.Now().Format(time.RFC3339), action, name)
}
