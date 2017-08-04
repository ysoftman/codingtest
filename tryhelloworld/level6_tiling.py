# ysoftman
# -*- coding: utf-8 -*-
'''
1x1 정사각형 2개가 붙어 있는 타일이 있습니다.
이 타일을 이용하여 총 3xN 의 보드판을 채우려고 합니다.
타일은 가로, 세로 두 가지 방향으로 배치할 수 있습니다.
보드판의 길이 N이 주어질 때, 3xN을 타일로 채울 수 있는 경우의 수를 반환하는 tiling 함수를 완성하세요.
단, 리턴하는 숫자가 매우 커질 수도 있으므로 숫자가 5자리를 넘어간다면 맨 끝자리 5자리만 리턴하세요.
예를 들어 N = 2일 경우 3을 반환해 주면 됩니다.
하지만 만약 답이 123456789라면 56789만 반환해주면 됩니다.
리턴하는 숫자의 앞자리가 0일 경우 0을 제외한 숫자를 리턴하세요. 리턴타입은 정수형이어야 합니다.
참고: 이 문제는 2 x n 타일링 문제와 유사합니다. 문제이해가 어려우면 2 x n 타일링 문제를 먼저 풀어 보세요.
'''

'''
pattern
n -> 3x1 에서의 answer
2 -> 3 = 3
3 -> 0
4 -> 11 = 3*3+1*2
5 -> 0
6 -> 41 = 11*3+3*2+1*2
7 -> 0
8 -> 153 = 41*3+11*2+3*2+1*2
9 -> 0
10 -> 571 = 153*3+41*2+11*2+3*2+1*2
11 -> 0
12 -> 2131 = 571*3+153*2+41*2+11*2+3*2+1*2
13 -> 0
14 -> 7953 = 2131*3+571*2+153*2+41*2+11*2+3*2+1*2
'''


def tiling(n):
    if n < 2:
        return 0
    if n % 2 != 0:
        return 0

    pre_acc = 2
    results = [1, 0, 3, 0]
    for i in range(4, n + 1):
        temp = 0
        if i % 2 == 0:
            temp = 3 * results[i - 2] + pre_acc
            pre_acc += 2 * results[i - 2]
        results.append(temp)

    answer = str(results[n])
    if len(answer) > 5:
        answer = answer[-5:]

    return int(answer)

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(tiling(2))
print(tiling(4))
print(tiling(6))
print(tiling(8))
print(tiling(10))
print(tiling(12))
print(tiling(14))
