package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberSystems(t *testing.T) {
	systems := Systems()
	assert.Equal(t, len(systems), 4)
}
