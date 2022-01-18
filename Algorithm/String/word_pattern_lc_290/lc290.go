package lc290

import (
	"bytes"
)

/**
https://leetcode.com/problems/word-pattern/
Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.



Example 1:

Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false


Constraints:

1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.

**/

func wordPattern(pattern string, s string) bool {
	pP, pN, sP, sN := 0, len(pattern), 0, len(s)
	// pP : pattern Pointer
	//      eg: "abba"
	//           ^
	// pN : count of pattern char
	//      eg: "abba" -> pN = 4
	//
	// sP : stinrg pointer
	//      eg: "cat dog cat"
	//           ^

	// ** only Lowercase English letters thus we have a-z which has 26 characters **
	// thus if we use an array(slice) with length of 26, the 0th position would represent "a"'s pattern, 1th position would represent "b"
	// eg.
	//       "abazz"          "cat dog cat goat goat"
	//       ["cat", "dog", "", "", "", ..., "", "goat"]
	record := make([]string, 26)

	// we also know pattern and string has 1-to-1 relationship
	// thus we can keep a map to check if the pattern is already mapped to another pattern character
	// Since the time complexity of the lookup is O(1), we are sacrificing memory for time
	/**
	eg.       "abazz"          "cat dog cat goat goat"
	{
	    "cat": 'a',  note: actual value here is 97 which is a byte
	    "dog": 'b',
	    "goat": 'z',
	}
	**/
	reverseMap := make(map[string]byte)

	// use a buffer to walk the string, when we encounter ' ' we know it's the end of a word, or when sp == sN-1, we are at the end of the input string
	var buffer bytes.Buffer

	// loop invariant is: we are not at the end of pattern and we are not at the end of input string
	for pP < pN && sP < sN {

		// current pattern char we are looking at
		pChar := pattern[pP]

		// recordIndex is not necessary, it's here to help you understand the position of pChar in the "record" slice
		// if pChar is 'b'
		// 'b' - 'a' yields 1
		// if pChar is 'a'
		// 'a' - 'a' yields 0
		// which represents the position of current pChar in our record slice
		recordIdx := pChar - 'a'

		// we encountered end of a word
		if sP == sN-1 || s[sP] == ' ' {

			// if at end of input string, add last character into buffer
			if sP == sN-1 {
				buffer.WriteByte(s[sP])
			}

			// eg. "cat"
			curString := buffer.String()

			// check the map to see if the pattern is already mapped to another value
			// the use of recordIdx is not necessary, we can use pChar instead since we are not using array
			//
			// another way to write this is
			// if reverseMap[curString] != 0 && reverseMap[curString] != pChar
			// when recording we just do reverseMap[pChar] =curString
			if cRecord, ok := reverseMap[curString]; ok && cRecord != recordIdx {
				return false
			}

			// check record array, if no string is assigned to pattern, assign the string to current index
			if record[recordIdx] == "" {
				record[recordIdx] = curString
			} else if record[recordIdx] != curString {
				// !! current pattern already points to another string
				//arrayhelper.PrintStringArray(record, int(recordIdx))
				return false
			}

			// record the current value in our reverseMap
			reverseMap[curString] = recordIdx

			// move pattern pointer forward
			pP++

			// reset the string buffer
			buffer.Reset()
		} else {
			// not end of a word
			buffer.WriteByte(s[sP])
		}
		// always move the pointer of string forward
		sP++
	}

	// if pattern is done but we still have leftover string, return false
	if pP < pN {
		return false
	}

	// if string is done but we still have leftover pattern, return false
	if sP < sN {
		return false
	}

	// everything checkedout return True
	return true
}
