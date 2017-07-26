# ysoftman
# -*- coding: utf-8 -*-
'''
물건 값보다 많은 돈을 낼 경우 거스름돈을 돌려주게 됩니다.
거스름돈은 한 금액의 돈으로 줄 수도 있지만, 여러 금액의 동전을 섞어 거슬러 줄 수 있습니다.
거스름돈이 N원일 때, 1원, 2원, 5원 동전이 있다면 몇 가지 방법으로 돈을 거슬러 줄 수 있을까요?
change 함수를 통해 경우의 수를 반환해주는 함수를 만들어 보세요. K에는 사용할 수 있는 동전의 종류가 들어 있습니다.
예를 들어, N = 5이고 K = [1, 2, 5]이면 1원, 2원, 5원 동전을 가지고 5원을 맞추는 경우의 수를 찾으면 됩니다.
1원 5개
1원 3개, 2원 1개
1원 1개, 2원 2개
5원 1개
이렇게 총 4가지 경우가 있으면, 4를 리턴해 주면 됩니다.
'''

'''
dynamic programming 이용
1 2 3 4 5 
1 1 1 1 1  -> 동전 1 이하로 만들 수 있는 경우의 수
  1 1 1 2  -> 동전 2 이하로 만들 수 있는 경우의 수
        1  -> 동전 5 이하로 만들 수 있는 경우의 수
5 의 경우의 수 = 1 + 2 + 1
'''


def change(total, coin):
    coin.sort()
    acc = {}
    for i in range(total + 1):
        acc[i] = 0
        if i % coin[0] == 0:
            acc[i] = 1

    for i in range(1, len(coin)):
        for j in range(coin[i], total + 1):
            # print acc
            acc[j] += acc[j - coin[i]]
    # print acc
    return acc[total]


print(change(5, [1, 2, 5]))
print(change(1000, [20, 17, 33, 26, 8]))  # 21876
