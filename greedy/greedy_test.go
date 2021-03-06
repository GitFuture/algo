package greedy

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestFindGreedy(t *testing.T) {
	s := FindGreedy(132)
	fmt.Println(s)
}

func TestGetMaxNumber(t *testing.T) {
	src.Asset(1, t, 4329, GetMaxNumber(1432219, 3))
	src.Asset(2, t, 439, GetMaxNumber(1432219, 4))
	src.Asset(3, t, 49, GetMaxNumber(1432219, 5))
	src.Asset(4, t, 21, GetMaxNumber(111121, 4))
}

