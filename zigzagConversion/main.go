package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ret := []rune{}
	n := len(s)
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret = append(ret, rune(s[j+i]))
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				ret = append(ret, rune(s[j+cycleLen-i]))
			}
		}
	}
	return string(ret)
}

// JAVA
// class Solution {
//     public String convert(String s, int numRows) {

//         if (numRows == 1) return s;

//         StringBuilder ret = new StringBuilder();
//         int n = s.length();
//         int cycleLen = 2 * numRows - 2;

//         for (int i = 0; i < numRows; i++) {
//             for (int j = 0; j + i < n; j += cycleLen) {
//                 ret.append(s.charAt(j + i));
//                 if (i != 0 && i != numRows - 1 && j + cycleLen - i < n)
//                     ret.append(s.charAt(j + cycleLen - i));
//             }
//         }
//         return ret.toString();
//     }
// }
