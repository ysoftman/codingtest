# https://leetcode.com/problems/tenth-line/
# 195. Tenth Line
# Easy
# Given a text file file.txt, print just the 10th line of the file.

# Example:
# Assume that file.txt has the following content:

# Line 1
# Line 2
# Line 3
# Line 4
# Line 5
# Line 6
# Line 7
# Line 8
# Line 9
# Line 10
# Your script should output the tenth line, which is:

# Line 10
# Note:
# 1. If the file contains less than 10 lines, what should you output?
# 2. There's at least three different solutions. Try to explore all possibilities.


# Read from the file file.txt and output the tenth line to stdout.
#!/bin/bash


cat << zzz > file.txt
1.apple
2.123
3.aaa
4.aaaaa
5.bbb
6.yoon
7.qqq
8.okay
9.ccc
10.lemon
11.zzz
zzz


# using sed
# sed -n 10p file.txt

# using loop
cnt=0
while read line; do
    let cnt++;
    if [[ $cnt == 10 ]]; then
        echo $line
        break
    fi
    # echo $cnt $line
done < file.txt


rm -rf file.txt
