# ysoftman
# -*- coding: utf-8 -*-
'''
단어 퍼즐은 주어진 단어 조각들을 이용해서 주어진 문장을 완성하는 퍼즐입니다.
이때, 주어진 각 단어 조각들은 각각 무한개씩 있다고 가정합니다.
예를 들어 주어진 단어 조각이 [“ba”, “na”, “n”, “a”]인 경우 "ba", "na", "n", "a" 단어 조각이 각각 무한개씩 있습니다.
이때, 만들어야 하는 문장이 “banana”라면 “ba”, “na”, “n”, “a”의 4개를 사용하여 문장을 완성할 수 있지만,
“ba”, “na”, “na”의 3개만을 사용해도 “banana”를 완성할 수 있습니다.
사용 가능한 단어 조각들을 담고 있는 배열 strs와 완성해야 하는 문자열 t가 매개변수로 주어질 때,
주어진 문장을 완성하기 위해 사용해야 하는 단어조각 개수의 최솟값을 return 하도록 solution 함수를 완성해 주세요.
만약 주어진 문장을 완성하는 것이 불가능하면 -1을 return 하세요.

제한사항
strs는 사용 가능한 단어 조각들이 들어있는 배열로, 길이는 1 이상 100 이하입니다.
strs의 각 원소는 사용 가능한 단어조각들이 중복 없이 들어있습니다.
사용 가능한 단어 조각들은 문자열 형태이며, 모든 단어 조각의 길이는 1 이상 5 이하입니다.
t는 완성해야 하는 문자열이며 길이는 1 이상 20,000 이하입니다.
모든 문자열은 알파벳 소문자로만 이루어져 있습니다.
입출력 예
strs	t	result
["ba","na","n","a"]	"banana"	3
["app","ap","p","l","e","ple","pp"]	"apple"	2
["ba","an","nan","ban","n"]	"banana"	-1
입출력 예 설명
입출력 예 #1
문제의 예시와 같습니다.

입출력 예 #2
"ap" 1개, "ple" 1개의 총 2개로 "apple"을 만들 수 있으므로 필요한 단어 개수의 최솟값은 2를 return 합니다.

입출력 예 #3
주어진 단어로는 "banana"를 만들 수 없으므로 -1을 return 합니다.

'''

'''
["ba", "na", "n", "a"] 단어가 있는 경우
각 자리에서가장 적은 매치수를 dp 에 기록한다.
dp[i] = min(curmin , dp[i+1]+1)
'''


BIGNUM = 99999


def matchWord(strs, t, dp, pos):
    size = len(t)
    if size == 0 or pos >= size:
        return 0
    if dp[pos] != BIGNUM:
        return dp[pos]

    for j in range(pos, pos + 5):
        if t[pos:j + 1] in strs:
            nextmatch = matchWord(strs, t, dp, j + 1)
            if dp[pos] > nextmatch + 1:
                dp[pos] = nextmatch + 1

    return dp[pos]


def solution(strs, t):
    size = len(t)
    dp = [BIGNUM for i in range(size)]
    matchWord(strs, t, dp, 0)
    # print dp
    if dp[0] == BIGNUM:
        return -1
    return dp[0]

print(solution(["bc", "abde", "ab", "a", "de"], "abcde"))  # 3
print(solution(["ba", "na", "n", "a"], "banana"))  # 3
print(solution(["app", "ap", "p", "l", "e", "ple", "pp"], "apple"))  # 2
print(solution(["ba", "an", "nan", "ban", "n"], "banana"))  # -1
