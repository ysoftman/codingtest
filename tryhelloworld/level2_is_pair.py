# ysoftman
# -*- coding: utf-8 -*-

'''
is_pair함수는 문자열 s를 매개변수로 입력받습니다.
s에 괄호가 알맞게 짝지어져 있으면 True를 아니면 False를 리턴하는 함수를 완성하세요.
예를들어 s가 "(hello)()"면 True이고, ")("이면 False입니다.
s가 빈 문자열("")인 경우는 없습니다.
'''


def is_pair(s):
    open_cnt = 0
    for ch in s:
        if ch == "(":
            open_cnt += 1
        elif ch == ")":
            open_cnt -= 1
        if open_cnt < 0:
            return False
    if open_cnt == 0:
        return True
    return False


# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(is_pair("(hello)()"))
print(is_pair(")("))
