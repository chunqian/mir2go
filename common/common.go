// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

import (
	"strconv"
	"strings"
)

func TagCount(source string, tag rune) int {
	tcount := 0
	for _, c := range source {
		if c == tag {
			tcount++
		}
	}
	return tcount
}

func ArrestStringEx(source, searchAfter, arrestBefore string) (string, string) {
	arrestStr := ""
	if source == "" {
		return "", ""
	}
	
	var goodData bool
	srcLen := len(source)
	
	if srcLen >= len(searchAfter) {
		if strings.HasPrefix(source, searchAfter) {
			source = source[len(searchAfter):]
			goodData = true
		} else {
			n := strings.Index(source, searchAfter)
			if n >= 0 {
				source = source[n+len(searchAfter):]
				goodData = true
			}
		}
	}

	if goodData {
		n := strings.Index(source, arrestBefore)
		if n >= 0 {
			arrestStr = source[:n]
			return source[n+len(arrestBefore):], arrestStr
		} else {
			return searchAfter + source, arrestStr
		}
	} else {
		n := strings.Index(source, searchAfter)
		if n >= 0 {
			return source[n:], arrestStr
		}
		return "", arrestStr
	}
}

func StrToInt(str string, def int) int {
	// 将截取的字符串转换为整数
	v, err := strconv.Atoi(str)
	if err != nil {
		v = def
	}
	return v
}

func GetValidStr3(str string, divider []rune) (string, string) {
	var dest strings.Builder
	var count, srcLen int
	srcLen = len(str)
	dest.Grow(srcLen) // pre-allocate memory for better performance

	dividerMap := make(map[rune]bool)
	for _, ch := range divider {
		dividerMap[ch] = true
	}

	for count < srcLen {
		ch := rune(str[count])
		if _, exists := dividerMap[ch]; exists {
			break
		} else {
			dest.WriteRune(ch)
		}
		count++
	}

	if count < srcLen {
		return dest.String(), str[count+1:]
	}
	return dest.String(), ""
}
