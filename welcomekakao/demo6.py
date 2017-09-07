# ysoftman
# -*- coding: utf-8 -*-
'''
N개의 스티커가 원형으로 연결되어 있습니다. 다음 그림은 N = 8인 경우의 예시입니다.
image
원형으로 연결된 스티커에서 몇 장의 스티커를 뜯어내어 뜯어낸 스티커에 적힌 숫자의 합이 최대가 되도록 하고 싶습니다.
단 스티커 한 장을 뜯어내면 양쪽으로 인접해있는 스티커는 찢어져서 사용할 수 없게 됩니다.

예를 들어 위 그림에서 14가 적힌 스티커를 뜯으면 인접해있는 10, 6이 적힌 스티커는 사용할 수 없습니다.
스티커에 적힌 숫자가 배열 형태로 주어질 때, 스티커를 뜯어내어 얻을 수 있는 숫자의 합의 최댓값을 return 하는 solution 함수를 완성해 주세요.
원형의 스티커 모양을 위해 배열의 첫 번째 원소와 마지막 원소가 서로 연결되어 있다고 간주합니다.

제한 사항
sticker는 원형으로 연결된 스티커의 각 칸에 적힌 숫자가 순서대로 들어있는 배열로, 길이(N)는 1 이상 100,000 이하입니다.
sticker의 각 원소는 스티커의 각 칸에 적힌 숫자이며, 각 칸에 적힌 숫자는 1 이상 100 이하의 자연수입니다.
원형의 스티커 모양을 위해 sticker 배열의 첫 번째 원소와 마지막 원소가 서로 연결되어있다고 간주합니다.
입출력 예
sticker	answer
[14, 6, 5, 11, 3, 9, 2, 10]	36
[1, 3, 2, 5, 4]	8
입출력 예 설명
입출력 예 #1
6, 11, 9, 10이 적힌 스티커를 떼어 냈을 때 36으로 최대가 됩니다.

입출력 예 #2
3, 5가 적힌 스티커를 떼어 냈을 때 8로 최대가 됩니다.
'''


'''
dynamic programming 알고리즘을 이용한다.
dp 자료구조
dp[리스트 크기][0] = dp 로 선택된 누적값
dp[리스트 크키][1] = 마지막 숫자부터 누적(시작된)것인지 여부

점화식
dp[i][0] = max(dp[i+1][0], dp[i+2][0]+sticker[i])

다음 4가지 경우에 대해서 처리하고 이 중 가장 큰 값을 선택하다.
14  6  5 11  3  9  2 10
10 14  6  5 11  3  9  2
2  10 14  6  5 11  3  9
9   2 10 14  6  5 11  3
'''


def getdp(sticker):
    size = len(sticker)

    dp = [[0 for j in range(2)] for i in range(size)]
    dp[size - 2][0] = sticker[-2]
    dp[size - 2][1] = 0
    dp[size - 1][0] = sticker[-1]
    dp[size - 1][1] = 1
    for i in range(size - 3, -2, -1):

        if dp[i + 2][0] + sticker[i] > dp[i + 1][0]:
            dp[i][0] = dp[i + 2][0] + sticker[i]
            dp[i][1] = dp[i + 2][1]
        else:
            dp[i][0] = dp[i + 1][0]
            dp[i][1] = dp[i + 1][1]

    # print dp
    if dp[0][1] == 0:
        return dp[0][0]
    else:
        return dp[1][0]


def solution(sticker):
    size = len(sticker)
    if size <= 3:
        return max(sticker)

    n1 = getdp(sticker)
    sticker.insert(0, sticker[-1])
    sticker.pop(-1)
    n2 = getdp(sticker)
    sticker.insert(0, sticker[-1])
    sticker.pop(-1)
    n3 = getdp(sticker)
    sticker.insert(0, sticker[-1])
    sticker.pop(-1)
    n4 = getdp(sticker)

    return max([n1, n2, n3, n4])


print(solution([1, 2]))
print(solution([1, 3, 2]))
print(solution([14, 6, 5, 11, 3, 9, 2, 10]))
print(solution([1, 3, 2, 1]))
print(solution([1, 3, 2, 5, 4]))
print(solution([100, 3, 5, 5, 4]))
print(solution([100, 3, 5, 200, 100]))
print(solution([100, 3, 5, 100, 500]))
