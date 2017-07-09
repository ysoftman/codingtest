# ysoftman
# -*- coding: utf-8 -*-

'''
앞뒤를 뒤집어도 똑같은 문자열을 palindrome이라고 합니다.
longest_palindrom함수는 문자열 s를 매개변수로 입력받습니다.
s의 부분문자열중 가장 긴 palindrom의 길이를 리턴하는 함수를 완성하세요.
예를들어 s가 "토마토맛토마토"이면 7을 리턴하고 "토마토맛있어"이면 3을 리턴합니다.
'''


# !!! 한글 글자는 환경에 따라 인코딩 필요 !!!
def longest_palindrom(s):
    # s = unicode(s, 'mbcs')
    result = ""
    for i in range(len(s)):
        j = len(s)
        while j > i:
            left_pivot = i + int(len(s[i:j]) / 2)
            right_pivot = left_pivot
            if int(len(s[i:j]) % 2) == 1:
                right_pivot += 1
            # print (i, left_pivot, right_pivot, j, "-->", s[i:left_pivot], s[right_pivot:j])
            temp = s[right_pivot:j]
            if s[i:left_pivot] == temp[::-1]:
                # print ("found", s[i:j])
                if len(result) < len(s[i:j]):
                    result = s[i:j]
                break
            j -= 1
    return len(result)


# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(longest_palindrom("토마토맛토마토"))
print(longest_palindrom("토마토맛있어"))
print(longest_palindrom("얌얌토마토"))
print(longest_palindrom("racecar"))
