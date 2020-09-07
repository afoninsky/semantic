package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var err error

	r, err := New("")
	assert.NoError(t, err)

	v1, _ := r.Current()
	fmt.Printf("current = %s\n", v1)

	tag, err := r.Tag()
	assert.NoError(t, err)
	fmt.Printf("tag = %s\n", tag)

	v2, _ := r.Current()
	fmt.Printf("current = %s\n", v2)

	next, err := r.Next()
	assert.NoError(t, err)
	fmt.Printf("next = %s\n", next)
}
