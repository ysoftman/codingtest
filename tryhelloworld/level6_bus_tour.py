# ysoftman
# -*- coding: utf-8 -*-
'''
버스정류장 N개가 있습니다. 정류장에는 1번부터 N번까지의 번호가 매겨져 있습니다.
2차원 배열에서 A번 정류장에서 B번 정류장으로 가는 버스가 있다면 O, 없다면 X로 표시합니다.
예를들어 3개의 버스정류장이 있을 때
X O X
X X O
O X X
와 같이 표시하면
1번 정류장에서 2번정류장으로 갈 수 있음(첫 번째 줄의 O)
2번 정류장에서 3번정류장으로 갈 수 있음(두 번째 줄의 O)
3번 정류장에서 1번정류장으로 갈 수 있음(세 번째 줄의 O)
이 표를 이용해서 버스를 계속 갈아타서라도 갈 수 있는 정류장을 모두 표시하는 2차원 배열을 리턴하세요.
위의 표와 같이 1번에서 2번 정류장으로, 그리고 2번에서 3번 정류장으로 가는 버스가 있다면,
한 번 갈아타면 1번에서 3번 정류장으로 갈 수 있다는 뜻이 됩니다.
busTour 함수를 통해 모든 정류장에서 출발해서 도달할 수 있는지를 표시하는 배열을 리턴하세요.
입력이
X O X
X X O
O X X

라면 다음과 같이 리턴하면 됩니다.
O O O
O O O
O O O
'''

'''
현재 버스 정류장과 다음 버스 정류장의 리스트와 or 연산을 수행
이 과정을 버스정류장 개수 n 만큼 반복 한다.
'''


def logical_or(a, b):
    temp = []
    for i in range(0, len(a)):
        if a[i] == 'O' or b[i] == 'O':
            temp.append('O')
        else:
            temp.append('X')
    return temp


def busTour(map):
    for n in range(0, len(map)):
        for i in range(0, len(map)):
            for j in range(0, len(map[i])):
                if map[i][j] == 'O':
                    map[i] = logical_or(map[i], map[j])
    return map


# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(busTour([['X', 'O', 'X'], ['X', 'X', 'O'], ['O', 'X', 'X']]))
