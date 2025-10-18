# https://www.acmicpc.net/problem/11478
import sys


# recursionError
def subString2(input, set):
    for i in range(len(input)):
        s.add(input[0 : i + 1])
        subString2(input[i + 1 :], s)


def subString(input, set):
    for i in range(len(input)):
        for j in range(i, len(input)):
            s.add(input[i : j + 1])


input = sys.stdin.readline().rstrip()
s = set()
subString(input, s)
print(len(s))
