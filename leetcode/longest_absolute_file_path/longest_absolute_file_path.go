/*
https://leetcode.com/problems/longest-absolute-file-path/
388. Longest Absolute File Path
Medium
Suppose we have a file system that stores both files and directories. An example of one system is represented in the following picture:

Here, we have dir as the only directory in the root. dir contains two subdirectories, subdir1 and subdir2. subdir1 contains a file file1.ext and subdirectory subsubdir1. subdir2 contains a subdirectory subsubdir2, which contains a file file2.ext.

In text form, it looks like this (with ⟶ representing the tab character):
dir
⟶ subdir1
⟶ ⟶ file1.ext
⟶ ⟶ subsubdir1
⟶ subdir2
⟶ ⟶ subsubdir2
⟶ ⟶ ⟶ file2.ext
If we were to write this representation in code, it will look like this: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext". Note that the '\n' and '\t' are the new-line and tab characters.

Every file and directory has a unique absolute path in the file system, which is the order of directories that must be opened to reach the file/directory itself, all concatenated by '/'s. Using the above example, the absolute path to file2.ext is "dir/subdir2/subsubdir2/file2.ext". Each directory name consists of letters, digits, and/or spaces. Each file name is of the form name.extension, where name and extension consist of letters, digits, and/or spaces.

Given a string input representing the file system in the explained format, return the length of the longest absolute path to a file in the abstracted file system. If there is no file in the system, return 0.

Note that the testcases are generated such that the file system is valid and no file or directory name has length 0.

Example 1:
Input: input = "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
Output: 20
Explanation: We have only one file, and the absolute path is "dir/subdir2/file.ext" of length 20.

Example 2:
Input: input = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
Output: 32
Explanation: We have two files:
"dir/subdir1/file1.ext" of length 21
"dir/subdir2/subsubdir2/file2.ext" of length 32.
We return 32 since it is the longest absolute path to a file.

Example 3:
Input: input = "a"
Output: 0
Explanation: We do not have any files, just a single directory named "a".

Constraints:
1 <= input.length <= 104
input may contain lowercase or uppercase English letters, a new line character '\n', a tab character '\t', a dot '.', a space ' ', and digits.
All file and directory names have positive length.
*/
package main

import (
	"fmt"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dir
⟶ subdir1
⟶ ⟶ file1.ext
⟶ ⟶ subsubdir1
⟶ subdir2
⟶ ⟶ subsubdir2
⟶ ⟶ ⟶ file2.ext
각 라인 마다 path 를 구성(stack)
*/
func lengthLongestPath(input string) int {
	maxFilePathLen := 0
	paths := strings.Split(input, "\n")
	stack := []string{paths[0]}
	// 파일인 경우
	if strings.Contains(paths[0], ".") {
		maxFilePathLen = len(paths[0])
	}

	for i := 1; i < len(paths); i++ {
		cur := strings.Split(paths[i], "\t")
		curDepth := len(cur)

		// 상위 디렉토리인 경우
		// 같은 디렉토리인 경우
		for curDepth <= len(stack) {
			// pop
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, cur[len(cur)-1])

		// fmt.Println(strings.Join(stack, "/"))

		// 파일인 경우
		if strings.Contains(cur[curDepth-1], ".") {
			curFilePath := strings.Join(stack, "/")
			maxFilePathLen = max(maxFilePathLen, len(curFilePath))
		}

	}
	return maxFilePathLen
}

func main() {
	// test := strings.Split("\ta\t", "\t")
	// fmt.Println(len(test), test)
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
	fmt.Println(lengthLongestPath("a"))
	fmt.Println(lengthLongestPath("a.txt"))
	fmt.Println(lengthLongestPath("a\n\tb1\n\t\tf1.txt\n\taaaaa\n\t\tf2.txt"))
	fmt.Println(lengthLongestPath("a\n\taa\n\t\taaa\n\t\t\tfile1.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png"))
}
