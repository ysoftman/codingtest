/*
https://leetcode.com/problems/letter-tile-possibilities
1079. Letter Tile Possibilities
Medium
You have n  tiles, where each tile has one letter tiles[i] printed on it.
Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles.

Example 1:
Input: tiles = "AAB"
Output: 8
Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".

Example 2:
Input: tiles = "AAABBC"
Output: 188

Example 3:
Input: tiles = "V"
Output: 1

Constraints:
1 <= tiles.length <= 7
tiles consists of uppercase English letters.
*/
package main

import "fmt"

var m map[string]bool

func letterTile(curTiles, restTiles string) {
	for i := range restTiles {
		m[curTiles+string(restTiles[i])] = true
		rt := restTiles[:i]
		rt += restTiles[i+1:]
		letterTile(curTiles+string(restTiles[i]), rt)
	}
}

func numTilePossibilities(tiles string) int {
	m = make(map[string]bool)
	letterTile("", tiles)
	// fmt.Println(m)
	return len(m)
}

func main() {
	fmt.Println(numTilePossibilities("AAB"))
	fmt.Println(numTilePossibilities("AAABBC"))
	fmt.Println(numTilePossibilities("AAABBCC"))
	fmt.Println(numTilePossibilities("V"))
}
