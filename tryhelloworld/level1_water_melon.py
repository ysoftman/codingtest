# ysoftman
# -*- coding: utf-8 -*-

# water_melon함수는 정수 n을 매개변수로 입력받습니다.
# 길이가 n이고, 수박수박수...와 같은 패턴을 유지하는 문자열을 리턴하도록 함수를 완성하세요.
# 예를들어 n이 4이면 '수박수박'을 리턴하고 3이라면 '수박수'를 리턴하면 됩니다.

def water_melon(n):
    # 함수를 완성하세요.
    out = ""
    for i in range(n):
        if i % 2 == 0:
            out += "수"
        else:
            out += "박"
    return out


# 실행을 위한 테스트코드입니다.
print("n이 3인 경우: " + water_melon(3));
print("n이 4인 경우: " + water_melon(4));
