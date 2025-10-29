# https://leetcode.com/problems/transpose-file/
# 194. Transpose File
# Medium
# Given a text file file.txt, transpose its content.

# You may assume that each row has the same number of columns, and each field is separated by the ' ' character.

# Example:

# If file.txt has the following content:

# name age
# alice 21
# ryan 30
# Output the following:

# name alice ryan
# age 21 30

#!/bin/bash

cat <<'zzz' >file.txt
name age
alice 21
ryan 30
zzz

# using awk(c 와 비슷한 문법사용)
awk '
# BEGIN 입력 파일 라인 처리전 실행
BEGIN {
    # print "BEGIN (before reading input)"
}
# 입력 파일 라인 처리
{
    # 단어 " " 카운드 로 기록
    for (i = 1; i <= NF; i++) {
        if (NR == 1) {
            s[i] = $i;
        } else {
            s[i] = s[i] " " $i;
        }
    }
}
# END 입력 파일 라인 처리 후 실행
END {
    for (i = 1; s[i] != ""; i++) {
        print s[i];
    }
}' file.txt

rm -rf file.txt
