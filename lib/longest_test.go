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
