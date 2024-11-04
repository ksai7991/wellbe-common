package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToInternationalPhoneNumber(t *testing.T) {
	tell1 := ConvertToInternationalPhoneNumber("080-1234-5678", "81")
	assert.Equal(t, "+818012345678", tell1)
	tell2 := ConvertToInternationalPhoneNumber("80-1234-5678", "81")
	assert.Equal(t, "+818012345678", tell2)
	tell3 := ConvertToInternationalPhoneNumber("8012345678", "81")
	assert.Equal(t, "+818012345678", tell3)
}
