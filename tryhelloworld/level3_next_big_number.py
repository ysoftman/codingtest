# ysoftman
# -*- coding: utf-8 -*-

'''
어떤 수 N(1≤N≤1,000,000) 이 주어졌을 때, N의 다음 큰 숫자는 다음과 같습니다.
N의 다음 큰 숫자는 N을 2진수로 바꾸었을 때의 1의 개수와 같은 개수로 이루어진 수입니다.
1번째 조건을 만족하는 숫자들 중 N보다 큰 수 중에서 가장 작은 숫자를 찾아야 합니다.
예를 들어, 78을 2진수로 바꾸면 1001110 이며, 78의 다음 큰 숫자는 83으로 2진수는 1010011 입니다.
N이 주어질 때, N의 다음 큰 숫자를 찾는 nextBigNumber 함수를 완성하세요.
'''


def nextBigNumber(n):
    answer = 0
    strbin = "{0:b}".format(n)
    # print strbin
    one_cnt = strbin.count("1")
    while answer == 0:
        n += 1
        if ("{0:b}".format(n)).count("1") == one_cnt:
            answer = n
            break

    return answer

# 아래 코드는 테스트를 위한 출력 코드입니다.
print(nextBigNumber(78))
