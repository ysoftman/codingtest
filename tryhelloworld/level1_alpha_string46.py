# ysoftman
# -*- coding: utf-8 -*-

# alpha_string46함수는 문자열 s를 매개변수로 입력받습니다.
# s의 길이가 4혹은 6이고, 숫자로만 구성되있는지 확인해주는 함수를 완성하세요.
# 예를들어 s가 "a234"이면 False를 리턴하고 "1234"라면 True를 리턴하면 됩니다

def alpha_string46(s):
    #함수를 완성하세요
    if not(len(s) == 4 or len(s) == 6):
        return False

    for i in range(0, len(s), 1):
        if (s[i] >= '0' and s[i] <= '9'):
            continue
        else:
            return False
            
    return True


# 아래는 테스트로 출력해 보기 위한 코드입니다.
print( alpha_string46("a234") )
print( alpha_string46("1234") )
print( alpha_string46("52527") )