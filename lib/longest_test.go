package lib

import (
    "testing"
)

func TestReturnsFl(t *testing.T) {
    strs := []string{"flower", "flow", "flight", "fly", "floor", "flu", "flag"}
    result := LongestCommonPrefix(strs)

    if result != "fl" {
        t.Errorf("Expected fl, got %s", result)
    }
}

func TestReturnsEmpty(t *testing.T) {
    strs := []string{"flower", "glow", "flight", "fly", "floor", "flu", "flag"}
    result := LongestCommonPrefix(strs)

    if result != "" {
        t.Errorf("Expected '', got %s", result)
    }
}
