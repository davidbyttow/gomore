package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFound(t *testing.T) {
	assert.Equal(t, 4.0, Round(4.0))
	assert.Equal(t, 4.0, Round(4.1))
	assert.Equal(t, 4.0, Round(4.49999))
	assert.Equal(t, 5.0, Round(4.50))
	assert.Equal(t, 5.0, Round(4.51))
	assert.Equal(t, 5.0, Round(4.999))
}
