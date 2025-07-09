/*
https://www.educative.io/m/find-permutation
Find Kth Permutation
Problem Statement
Given a set of ‘n’ elements, find their Kth permutation. Consider the following set of elements:

1,2,3


All permutations of the above elements are (with ordering):
1st : 1,2,3
2nd : 1,3,2
3rd : 2,1,3
4th : 2,3,1
5th : 3,1,2
6th : 3,2,1

Here we need to find the Kth permutation

---
다음과 같은 방식으로 permutation(치환) 을 수행한다.
                                      123
0번째고정할값들    (0<->0)123                 (0<->1)213                   (0<->2)312
1번째고정할값들 (1<->1)123 (1<->2)132     (1<->1)213 (1<->2)231      (1<->1)312 (1<->2)321
*/

#include <iostream>
#include <vector>

using namespace std;

int factorial(int n) {
    if (n == 0 || n == 1)
        return 1;
    return n * factorial(n - 1);
}

void recursive_permuation(vector<char> &v, int a, int b, vector<string> &rv) {
    if (a == b) {
        string str(v.begin(), v.end());
        rv.push_back(str);
        return;
    }
    for (int i = a; i <= b; i++) {
        // swap
        char temp = v[a];
        v[a]      = v[i];
        v[i]      = temp;
        recursive_permuation(v, a + 1, b, rv);
        // 다음 루프를 위해 원복
        temp = v[a];
        v[a] = v[i];
        v[i] = temp;
    }
}

void find_kth_permutation(
    vector<char> &v,
    int           k,
    string       &result) {
    // TODO: Write - Your - Code
    vector<string> rv;
    recursive_permuation(v, 0, v.size() - 1, rv);

    result = rv[k - 1];
}

string get_permutation(int n, int k) {
    vector<char> v;
    for (char i = 1; i <= n; ++i) {
        v.push_back(i + '0');
    }

    string result;
    find_kth_permutation(v, k, result);
    return result;
}

int main(int argc, char *argv[]) {
    for (int i = 1; i <= factorial(4); ++i) {
        cout << i << "th permutation = " << get_permutation(4, i) << endl;
    }
}