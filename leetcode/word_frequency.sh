# https://leetcode.com/problems/word-frequency/
# 192. Word Frequency
# Medium
# Write a bash script to calculate the frequency of each word in a text file words.txt.

# For simplicity sake, you may assume:

# words.txt contains only lowercase characters and space ' ' characters.
# Each word must consist of lowercase characters only.
# Words are separated by one or more whitespace characters.
# Example:

# Assume that words.txt has the following content:

# the day is sunny the the
# the sunny is is
# Your script should output the following, sorted by descending frequency:

# the 4
# is 3
# sunny 2
# day 1
# Note:

# Don't worry about handling ties, it is guaranteed that each word's frequency count is unique.
# Could you write it in one-line using Unix pipes?

#!/bin/bash

cat << 'zzz' > words.txt
the day is sunny the the
the sunny is is
zzz

# 공백과 줄바꿈은 중복 없이 한번만 표시
# 오름 차순 정렬
# 같은 필드명 카운트하고 하나로 표시
# 내림 차순 정렬
# 1열 2열 위치 바꿔서 출력
# cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -r | awk '{print $2,$1}'
# 또는
cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -r | awk '{print $2" "$1}'

rm -rf words.txt
