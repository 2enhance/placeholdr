package providers_test

import (
	"testing"

	"github.com/2enhance/placeholdr/config/providers"

	"github.com/stretchr/testify/assert"
)

func TestColors_BackgroundColor(t *testing.T) {
	bg := providers.BackgroundColor("foo")

	assert.Equal(t, "rgb(25,100,100)", bg.Format())
}
