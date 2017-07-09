# ysoftman
# -*- coding: utf-8 -*-

'''
1부터 입력받은 숫자 n 사이에 있는 소수의 개수를 반환하도록 numberOfPrime 메소드를 만들어 보세요.

소수는 1과 자기 자신으로만 나누어지는 수를 의미합니다.
(1은 소수가 아닙니다.)

10을 입력받았다면, 1부터 10 사이의 소수는 [2,3,5,7] 4개가 존재하므로 4를 반환
5를 입력받았다면, 1부터 5 사이의 소수는 [2,3,5] 3개가 존재하므로 3를 반환
'''


def numberOfPrime(n):
    result = []
    for i in range(2, n + 1):
        for j in range(2, i + 1):
            if j == i:
                result.append(i)
            if int(i % j) == 0:
                break

    return len(result)


# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(numberOfPrime(10))
print(numberOfPrime(5))
