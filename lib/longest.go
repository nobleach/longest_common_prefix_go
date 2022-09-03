package lib

import (
    "bytes"
    "sort"
)

func LongestCommonPrefix(strs []string) string {
    sort.Strings(strs)
    pointer := 0
    firstStr := strs[0]
    lastStr := strs[len(strs) - 1]
    var longestPrefix bytes.Buffer
    longestPrefix.WriteString("")

    for pointer <= len(firstStr) - 1 {
        if (firstStr[pointer] != lastStr[pointer]) {
            break
        }


        longestPrefix.WriteByte(firstStr[pointer]) 
        pointer++
    }

    return longestPrefix.String()
}
