package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	items := []string{"foo", "bar", "wiz", "bang"}
	exclude := []string{"bar", "bang"}
	filtered := filter(items, exclude)
	expected := []string{"foo", "wiz"}
	for i, v := range filtered {
		assert.Equal(t, v, expected[i])
	}
}

func TestWrapStrings(t *testing.T) {
	mark := "mark"
	items := []string{"foo", "bar", "wiz", "bang"}
	wrapped := wrapStrings(items, mark)
	for i, v := range items {
		assert.Equal(t, mark+v+mark, wrapped[i])
	}
}

func TestPrependStrings(t *testing.T) {
	prefix := "mark"
	items := []string{"foo", "bar", "wiz", "bang"}
	prepended := prependStrings(items, prefix)
	for i, v := range items {
		assert.Equal(t, prefix+v, prepended[i])
	}
}