package util

import (
	"fmt"
	"ql/interfaces"
)

// AreStmtListsEqual returns a bool indicatiing if two passed stmt lists are equal in content
func AreStmtListsEqual(a, b interfaces.StmtList) bool {
	// slices don't support equality checking by default
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}
