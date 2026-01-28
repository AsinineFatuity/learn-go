package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IdentifyPrefixPostfix(userID, email string) bool {
	hasPrefix := strings.HasPrefix(email, userID)
	hasSuffix := strings.HasSuffix(email, userID)
	return hasPrefix || hasSuffix
}

func ContainsEducative(email string) bool {
	target := "educative"
	return strings.Contains(strings.ToLower(email), target)
}

func MaskUserName(email string) string {
	splitEmail := strings.Split(email, "@")
	if len(splitEmail) != 2 || len(splitEmail[0]) == 0 {
		return email //malformed email
	}
	username := splitEmail[0]
	if len(username) <= 2 {
		return email //nothing to mask
	}
	firstChar := string(username[0])
	lastChar := string(username[len(username)-1])
	unmasked := "@" + splitEmail[1]
	mask := strings.Repeat("*", len(username)-2)
	return firstChar + mask + lastChar + unmasked
}

func IndexOfAtSymbol(email string) int {
	return strings.Index(email, "@")
}

func TrimAndSplitUserID(userID string) string {
	trimmedStr := strings.TrimSpace(userID)
	splitId := strings.Split(trimmedStr, "-")
	return splitId[1]
}

func ConvertStringToInt(str string) int {
	valInt, _ := strconv.Atoi(str)
	return valInt
}

func StringManipulationRunner() {
	// Test your functions here
	fmt.Println(IdentifyPrefixPostfix(".io", "evangeline@educative.io")) // true
	fmt.Println(IdentifyPrefixPostfix("UID", "UID-0123"))                // true
	fmt.Println(IdentifyPrefixPostfix("UID", "evangeline@educative.io")) // false
	fmt.Println(ContainsEducative("evangeline@educative.io"))            // true
	fmt.Println(MaskUserName("evangeline@educative.io"))                 // e******e@educative.io
	fmt.Println(IndexOfAtSymbol("evangeline@educative.io"))              // 10
	fmt.Println(TrimAndSplitUserID("UID-0123"))                          // 0123
	fmt.Println(ConvertStringToInt("123"))                               // 123
}
