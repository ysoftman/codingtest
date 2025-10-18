# https://www.acmicpc.net/problem/10872
import sys


def fact(n):
    if n <= 1:
        return 1
    return n * fact(n - 1)


input = sys.stdin.readline().rstrip()
print(fact(int(input)))
