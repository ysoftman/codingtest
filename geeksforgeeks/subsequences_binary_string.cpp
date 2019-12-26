/*
ysoftman

https://www.geeksforgeeks.org/number-of-subsequences-in-a-given-binary-string-divisible-by-2/

Given a binary string str of length N, the task is to find the count of subsequences of str which are divisible by 2. Leading zeros in a sub-sequence is allowed.

Examples:

Input: str = “101”
Output: 2
“0” and “10” are the only subsequences
which are divisible by 2.

Input: str = “10010”
Output: 22

풀이)
2진수 스트링에서 2로 나눌 수 있는 부분문자열(subsequences)의 개수 구하기

- 끝이 0으로 끝아야 2로 나눌 수 있다.
- 2진수 -> 10진수 변환할때 자리수가 높이지면 2의 배수로 증가
*/
#include <iostream>

using namespace std;

int countSubsequencesInBinaryString(string str)
{
    int result = 0;
    int mul = 1;
    for (int i = 0; i < str.length(); i++)
    {
        if (str[i] == '0')
        {
            result += mul;
        }
        mul *= 2;
    }
    return result;
}

int main()
{
    string input("101");
    cout << input << " --> " << countSubsequencesInBinaryString(input) << endl;
    input = "10010";
    cout << input << " --> " << countSubsequencesInBinaryString(input) << endl;
    input = "10000";
    cout << input << " --> " << countSubsequencesInBinaryString(input) << endl;
    input = "1111111111";
    cout << input << " --> " << countSubsequencesInBinaryString(input) << endl;
    return 0;
}
