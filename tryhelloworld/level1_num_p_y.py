# ysoftman
# -*- coding: utf-8 -*-

# numPY함수는 대문자와 소문자가 섞여있는 문자열 s를 매개변수로 입력받습니다.
# s에 'p'의 개수와 'y'의 개수를 비교해 같으면 True, 다르면 False를 리턴하도록 함수를 완성하세요. 'p', 'y' 모두 하나도 없는 경우는 항상 True를 리턴합니다.
# 예를들어 s가 "pPoooyY"면 True를 리턴하고 "Pyy"라면 False를 리턴합니다.

def numPY(s):
    # 함수를 완성하세요
    cntp, cnty = 0, 0
    for i in range(0, len(s), 1):
        if s[i].lower() == 'y':
            cnty += 1
        elif s[i].lower() == 'p':
            cntp += 1
    
    if cntp != cnty:
        return False
    return True



# 아래는 테스트로 출력해 보기 위한 코드입니다.
print( numPY("pPoooyY") )
print( numPY("Pyy") )