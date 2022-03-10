package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMobileValidate(t *testing.T) {
	assert.Equal(t, true, MobileValidate("13600009181"))
	assert.Equal(t, false, MobileValidate("1334"))
}
