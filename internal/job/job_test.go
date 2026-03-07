package job

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCorn(t *testing.T) {
	NewCorn(time.Second*1, func() {
		fmt.Println(time.Now())
	})
}
