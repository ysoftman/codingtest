# ysoftman
# -*- coding: utf-8 -*-

'''
행렬의 곱셈은, 곱하려는 두 행렬의 어떤 행과 열을 기준으로, 좌측의 행렬은 해당되는 행, 
우측의 행렬은 해당되는 열을 순서대로 곱한 값을 더한 값이 들어갑니다.
행렬을 곱하기 위해선 좌측 행렬의 열의 개수와 우측 행렬의 행의 개수가 같아야 합니다.
곱할 수 있는 두 행렬 A,B가 주어질 때, 행렬을 곱한 값을 출력하는 productMatrix 함수를 완성해 보세요`.
'''


def productMatrix(A, B):
    answer = []
    for i in range(len(A)):
        element = []
        for j in range(len(B[0])):
            temp = 0
            for k in range(len(A[i])):
                temp += A[i][k] * B[k][j]
            element.append(temp)
        answer.append(element)
    return answer


# 아래는 테스트로 출력해 보기 위한 코드입니다.
a = [[1, 2],
     [2, 3]]
b = [[3, 4],
     [5, 6]]
print("결과 : {}".format(productMatrix(a, b)));
a = [[1, 2, 1],
     [2, 3, 1]]
b = [[3, 4],
     [5, 6],
     [1, 1]]
print("결과 : {}".format(productMatrix(a, b)));
a = [[1],
     [2]]
b = [[3, 4]]
print("결과 : {}".format(productMatrix(a, b)));
