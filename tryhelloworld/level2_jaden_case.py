# ysoftman
# -*- coding: utf-8 -*-

"""
Jaden_Case함수는 문자열 s을 매개변수로 입력받습니다.
s에 모든 단어의 첫 알파벳이 대문자이고, 그 외의 알파벳은 소문자인 문자열을 리턴하도록 함수를 완성하세요
예를들어 s가 "3people unFollowed me for the last week"라면
"3people Unfollowed Me For The Last Week"를 리턴하면 됩니다.
"""


def Jaden_Case(s):
    # 함수를 완성하세요
    words = s.split(" ")
    out = ""
    for i in words:
        i = i.lower()
        out += i[0].upper() + i[1:] + " "
    out = out[: len(out) - 1]
    return out


# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(Jaden_Case("3people unFollowed me for the last week"))

