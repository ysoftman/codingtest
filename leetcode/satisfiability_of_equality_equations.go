/*
https://leetcode.com/problems/satisfiability-of-equality-equations/
990. Satisfiability of Equality Equations
Medium

You are given an array of strings equations that represent relationships between variables where each string equations[i] is of length 4 and takes one of two different forms: "xi==yi" or "xi!=yi".Here, xi and yi are lowercase letters (not necessarily different) that represent one-letter variable names.

Return true if it is possible to assign integers to variable names so as to satisfy all the given equations, or false otherwise.

Example 1:
Input: equations = ["a==b","b!=a"]
Output: false
Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.
There is no way to assign the variables to satisfy both equations.

Example 2:
Input: equations = ["b==a","a==b"]
Output: true
Explanation: We could assign a = 1 and b = 1 to satisfy both equations.


Constraints:
1 <= equations.length <= 500
equations[i].length == 4
equations[i][0] is a lowercase letter.
equations[i][1] is either '=' or '!'.
equations[i][2] is '='.
equations[i][3] is a lowercase letter.
*/

package main

import "fmt"

/*
ex) a=b b=c c=a
a -> b -> c -> a

ex) a=b b!=a
a -> b -> b
*/
func find(m map[byte]byte, k byte) byte {
	if m[k] != k {
		// fmt.Println(string(k), string(m[k]))
		m[k] = find(m, m[k])
	}
	return m[k]
}

func equationsPossible(equations []string) bool {
	m := make(map[byte]byte)
	// a~z = a~z 로 m 초기화
	for i := 'a'; i <= 'z'; i++ {
		m[byte(i)] = byte(i)
	}
	for i := 0; i < len(equations); i++ {
		left := equations[i][0]
		right := equations[i][3]
		// == 인 경우만 m 업데이트
		if equations[i][1] == '=' {
			// m 업데이트시에도 find 로 recursive 하게 최종 값을 찾아 설정한다.
			m[find(m, left)] = find(m, right)
		}
	}

	for i := 0; i < len(equations); i++ {
		left := equations[i][0]
		right := equations[i][3]
		// != 인 경우 left, right 값이 같으면 안된다.
		if equations[i][1] == '!' {
			// find 로 recursive 하게 최종 값을 찾는다.
			if find(m, left) == find(m, right) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))
	fmt.Println(equationsPossible([]string{"a==b", "b==c", "a==c"}))
	fmt.Println(equationsPossible([]string{"a==b", "e==c", "b==c", "a!=e"}))
	fmt.Println(equationsPossible([]string{"f==a", "a==b", "f!=e", "a==c", "b==e", "c==f"}))
}
