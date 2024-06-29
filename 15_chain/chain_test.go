package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensitiveWordFilterChain_Filter(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.Equal(t, false, chain.Filter("test"))

	chain.AddFilter(&PoliticalWordFilter{})
	assert.Equal(t, true, chain.Filter("test"))
}

func TestCheck(t *testing.T) {
	chain := &Present{}
	chain.AddCheck(&UserCheck{})
	chain.AddCheck(&GroupCheck{})
	chain.Check()
}
