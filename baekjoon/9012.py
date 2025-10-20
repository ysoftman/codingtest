# https://www.acmicpc.net/problem/9012
import sys


def parenthesisString(s):
    stack = []
    for c in s:
        if c == "(":
            stack.append(c)
        elif c == ")":
            if len(stack) == 0:
                return "NO"
            stack.pop()
    # print(stack)
    if len(stack) == 0:
        return "yes".upper()
    return "NO"


# print(parenthesisString("(())())"))
# print(parenthesisString("(((()())()"))
# print(parenthesisString("(()())((()))"))
# print(parenthesisString("((()()(()))(((())))()"))
# print(parenthesisString("()()()()(()()())()"))
# print(parenthesisString("(()((())()("))
# input = sys.stdin.readline().rstrip()
input = sys.stdin.readlines()
for i in range(1, int(input[0]) + 1):
    print(parenthesisString(str(input[i].rstrip())))
