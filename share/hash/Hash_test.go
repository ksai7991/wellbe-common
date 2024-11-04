package hash

import (
    "testing"
    "github.com/stretchr/testify/assert"
	
)

func TestGetHash(t *testing.T) {
	hash := GetHash([]string{"foo bar","hoge hoge"})
    assert.Equal(t, "586657bc74225807f60b045869bdbdc316e76cd72ffc55f91653b397d5c64f1d", hash)
}