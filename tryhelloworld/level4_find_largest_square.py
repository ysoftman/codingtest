# ysoftman
# -*- coding: utf-8 -*-

'''
O와 X로 채워진 표가 있습니다. 표 1칸은 1 x 1 의 정사각형으로 이루어져 있습니다.
표에서 O로 이루어진 가장 큰 정사각형을 찾아 넓이를 반환하는 findLargestSquare 함수를 완성하세요.
예를 들어
1	2	3	4	5
X	O	O	O	X
X	O	O	O	O
X	X	O	O	O
X	X	O	O	O
X	X	X	X	X
가 있다면 정답은

1	2	3	4	5
X	O	O	O	X
X	O	_	_	_
X	X	_	_	_
X	X	_	_	_
X	X	X	X	X
가 되며 넓이는 9가 되므로 9를 반환해 주면 됩니다.
'''


# 아래 슬라이딩 윈도우를 하나씩 대입하는것 이외도
# board 전체를 x=0, O=1 로 설정한다음 (x-1,y-1),(x,y-1),(x-1,y) 가 > 0 이면
# (x,y) = (x-1,y-1) + 1 를 설정하는것으로도 최대 정사각형 크기를 구할 수 있다.
def findLargestSquare(board):
    height = len(board)
    width = len(board[0])
    if height == 1:
        return 0
    sw_h = sw_w = max(height, width)

    while sw_w > 1:
        for y in range(height):
            if y + sw_w > height:
                break
            for x in range(width):
                if x + sw_w > width:
                    break

                find = True
                # sliding windows
                for j in range(y, y + sw_h):
                    for i in range(x, x + sw_w):
                        if board[j][i] == 'X':
                            find = False
                            break
                    if not find:
                        break
                # print (find, x, y, sw_h, sw_w)
                if find:
                    return sw_h * sw_w
        sw_w -= 1
        sw_h -= 1
    return 0

# 아래 코드는 출력을 위한 테스트 코드입니다.
testBoard = [['X', 'O', 'O', 'O', 'X'], ['X', 'O', 'O', 'O', 'O'], [
    'X', 'X', 'O', 'O', 'O'], ['X', 'X', 'O', 'O', 'O'], ['X', 'X', 'X', 'X', 'X']]
print(findLargestSquare(testBoard))
testBoard = [['X', 'O', 'O', 'O', 'X'], ['X', 'O', 'O', 'O', 'O'], [
    'X', 'X', 'O', 'O', 'O'], ['X', 'X', 'O', 'O', 'O']]
print(findLargestSquare(testBoard))
