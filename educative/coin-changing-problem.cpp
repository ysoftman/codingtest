/*
https://www.educative.io/m/coin-changing-problem

Coin Changing Problem

Problem Statement
Suppose we have coin denominations of [1, 2, 5] and the total amount is 7.
We can make changes in the following 6 ways:

Denomination: 1,2,5
amount: 7

No. of ways to make the change
1,1,1,1,1,1,1
1,1,1,1,1,2
1,1,1,2,2
1,2,2,2
1,1,5
2,5

Total Methods: 6

----

[0] => methods: 1

[1] => methods: 1
1

[2] => methods: 2
1,1
2

[3] => methods: 2
1,1,1
1,2

[4] => methods: 3
1,1,1,1
1,1,2
2,2

[5] => methods: 4
1,1,1,1,1
1,1,1,2
1,2,2
5

[6] => methods: 5
1,1,1,1,1,1
1,1,1,1,2
1,1,2,2
2,2,2
1,5

[7] => methods: 6
1,1,1,1,1,1,1
1,1,1,1,1,2
1,1,1,2,2
1,2,2,2
1,1,5
2,5
---

1,2,5 동전을 하나씩 추가하가면서 총량을 만들 수 있는 경우의 파악
이전 작은 동전의 경우의 수는 계속 재사용될 수 있어 memo해둔다.(memoization=>dynamic programming)
0 1 2 3 4 5 6 7 ==> 만들고자 하는 총량(amount)
1 1 1 1 1 1 1 1 ==> 1 동전만을 사용해서 만들 수 있는 경우의 수(방법:methods)
1 1 2 2 3 3 4 4 ==> 1,2 동전을 사용해서 만들 수 있는 경우의 수(방법:methods)
1 1 2 2 3 4 5 6 ==> 1,2,5 동전을 사용해서 만들 수 있는 경우의 수(방법:methods)

ex)
methods[7] += methods[7-1] = 1 ==> 동전 1만 사용한 단계에서의 method[6] = 1
methods[7] += methods[7-2] = 1+3 ==> 동전 1,2 사용한 단계에서의 method[5] = 3
methods[7] += methods[7-5] = 1+3+2 ==> 동전 1,2,5 사용한 단계에서의 method[2] = 2

7-1 를 하는 이유 : 6의 경우의 수들 각각에 동전 1를 더하면 7이 되기 때문에, 1의 경우의 수를 그대로 가져와 쓸 수 있다.
7-2 를 하는 이유 : 5의 경우의 수들 각각에 동전 2를 더하면 7이 되기 때문에, 5의 경우의 수를 그대로 가져와 쓸 수 있다.
7-5 를 하는 이유 : 2의 경우의 수들 각각에 동전 5를 더하면 7이 되기 때문에, 2의 경우의 수를 그대로 가져와 쓸 수 있다.


dynamic programming 에 사용할 점화식
methods[amount] += methods[amount-coin]]


동전 1만 사용했을때
methods[0] += methods[0] = 1 ==> 동전 1은 사용할 수 없음, 0을 만들 수 있는 경우는 1로 설정
methods[1] += methods[1-1] = 1
methods[2] += methods[2-1] = 1
methods[3] += methods[3-1] = 1
methods[4] += methods[4-1] = 1
methods[5] += methods[5-1] = 1
methods[6] += methods[6-1] = 1
methods[7] += methods[7-1] = 1

동전 1,2 사용했을때 = 동전 1 사용했을때 + 동전 2까지 사용했을때
methods[0] += 0 = 1 ==> 동전 1,2는 사용할 수 없음, 0을 만들 수 있는 경우는 1로 설정
methods[1] += 0 = 1 ==> 동전 2는 사용할 수 없음, 동전 1만 사용한 것으로 끝
methods[2] += methods[2-2] = 1+1=2
methods[3] += methods[3-2] = 1+1=2
methods[4] += methods[4-2] = 1+2=3
methods[5] += methods[5-2] = 1+2=3
methods[6] += methods[6-2] = 1+3=4
methods[7] += methods[7-2] = 1+3=4

동전 1,2,5 사용했을때 = 동전 1 사용했을때 + 동전 2까지 사용했을때 + 동전 5까지 사용했을때
methods[0] += 0 = 1 ==> 동전 1,2,5는 사용할 수 없음, 0을 만들 수 있는 경우는 1로 설정
methods[1] += 0 = 1 ==> 동전 2,5는 사용할 수 없음, 동전 1만 사용한 것으로 끝
methods[2] += 0 = 2 ==> 동전 5는 사용할 수 없음, 동전 1,2만 사용한 것으로 끝
methods[3] += 0 = 2 ==> 동전 5는 사용할 수 없음, 동전 1,2만 사용한 것으로 끝
methods[4] += 0 = 3 ==> 동전 5는 사용할 수 없음, 동전 1,2만 사용한 것으로 끝
methods[5] += methods[5-5] = 3+1=4
methods[6] += methods[6-5] = 4+1=5
methods[7] += methods[7-5] = 4+2=6
*/

#include <iostream>
#include <vector>
using namespace std;

int solve_coin_change(vector<int> &denominations, int amount) {
    // TODO: Write - Your - Code
    vector<int> methods(amount + 1);
    methods[0] = 1;
    // time complexity : O(mxn)
    for (auto coin : denominations) {
        for (int i = coin; i <= amount; i++) {
            methods[i] += methods[i - coin];
        }
    }
    return methods[amount];
}

int main() {
    vector<int> denominations = {1, 2, 5};
    int         amount        = 7;
    int         result        = solve_coin_change(denominations, amount);
    // printing result
    cout << "solve_coin_change_dp([";
    for (int den : denominations) {
        cout << den << " ";
    }
    cout << "], " << amount << ") = " << result << endl;
}
