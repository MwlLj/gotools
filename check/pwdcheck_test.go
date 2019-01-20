package check

import (
	"fmt"
	"testing"
)

func TestComplicatedPwdCheck(t *testing.T) {
	result := ComplicatedPwdCheck([]rune("123abcABC"), 8, 30)
	fmt.Println(result)
}
