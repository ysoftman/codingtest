# ysoftman
# -*- coding: utf-8 -*-

# 두 수를 입력받아 두 수의 최대공약수와 최소공배수를 반환해주는 gcdlcm 함수를 완성해 보세요. 배열의 맨 앞에 최대공약수, 그 다음 최소공배수를 넣어 반환하면 됩니다. 예를 들어 gcdlcm(3,12) 가 입력되면, [3, 12]를 반환해주면 됩니다.

def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a%b)

def lcm(a, b):
    big = max(a,b)
    small = min(a,b)
    while True:
        if big % small == 0:
            break
        big += max(a,b)
    
    return big

def gcdlcm(a, b):
    answer = []
    val1 = gcd(a,b)
    val2 = lcm(a, b)
    # print val1
    answer.append(val1)
    answer.append(val2)
    return answer

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(gcdlcm(3,12))
print(gcdlcm(12,3))
print(gcdlcm(46,73))
