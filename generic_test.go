package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	items := []string{"foo", "bar", "wiz", "bang"}
	for _, v := range items {
		assert.True(t, in(v, items))
	}
	assert.False(t, in("nothere", items))
}
