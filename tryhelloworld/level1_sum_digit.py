# ysoftman
# -*- coding: utf-8 -*-

# sum_digit함수는 자연수를 전달 받아서 숫자의 각 자릿수의 합을 구해서 return합니다.
# 예를들어 number = 123이면 1 + 2 + 3 = 6을 return하면 됩니다.
# sum_digit함수를 완성해보세요.

def sum_digit(number):
    '''number의 각 자릿수를 더해서 return하세요'''
    sum = 0
    while number > 0:
        sum += number % 10
        number = int(number / 10)
    return sum

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print("결과 : {}".format(sum_digit(123)));
print("결과 : {}".format(sum_digit(9757)));