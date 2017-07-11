# ysoftman
# -*- coding: utf-8 -*-

'''
효진이는 멀리 뛰기를 연습하고 있습니다. 효진이는 한번에 1칸, 또는 2칸을 뛸 수 있습니다. 칸이 총 4개 있을 때, 효진이는
(1칸, 1칸, 1칸, 1칸)
(1칸, 2칸, 1칸)
(1칸, 1칸, 2칸)
(2칸, 1칸, 1칸)
(2칸, 2칸)
의 5가지 방법으로 맨 끝 칸에 도달할 수 있습니다. 멀리뛰기에 사용될 칸의 수 n이 주어질 때,
효진이가 끝에 도달하는 방법이 몇 가지인지 출력하는 jumpCase 함수를 완성하세요.
예를 들어 4가 입력된다면, 5를 반환해 주면 됩니다.
'''
import re


def jumpCase(num):
    if num == 0:
        return 1
    elif num < 0:
        return 0
    out = 0
    for i in range(1, 3):
        out += jumpCase(num - i)
    return out


def ternary(num):
    if num == 0:
        return [0]
    out = []
    while num > 0:
        # 0 이 포함되면 이 수에 대허서는 진행하지 않는다.
        if num % 3 == 0:
            return [0]
        out.append(num % 3)
        num = int(num / 3)
    out = out[::-1]
    # return int("".join(str(x) for x in out))
    return out


# 느려서 안됨
def jumpCase2(num):
    answer = 0
    endnum = ""
    for i in range(num):
        endnum += "2"

    endnum = int(endnum)
    dec = 0
    ter = 0
    real_cnt = 0
    while ter < endnum:
        ter_list = ternary(dec)
        if ter_list[0] == 0:
            dec += 1
            continue

        # print ter, sum(ter_list)
        if sum(ter_list) == num:
            # print ter, sum(ter_list)
            answer += 1

        ter = int("".join(str(x) for x in ter_list))

        real_cnt += 1
        dec += 1
    return answer, dec


# 무지 느려서 안됨
def jumpCase3(num):
    end_state = ""
    for i in range(num):
        end_state += "2"
    end_state = int(end_state)
    result_cnt = 0
    cnt = 0
    while cnt <= end_state:
        cnt += 1
        if len(re.findall("[0|3-9]", str(cnt))) > 0:
            continue
        list_cnt = [int(i) for i in str(cnt)]
        # print list_cnt, sum(list_cnt)
        if sum(list_cnt) == num:
            # print list_cnt, sum(list_cnt)
            result_cnt += 1

    return result_cnt


# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(jumpCase(4))
print(jumpCase(5))
print(jumpCase(6))
print(jumpCase(8))
print(jumpCase(9))
print(jumpCase(10))
print(jumpCase(20))
