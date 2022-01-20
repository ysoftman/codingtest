/*
https://www.educative.io/blog/0-1-knapsack-problem-dynamic-solution

        itemA  itemB itemC itemD 
value:    30     20    40   10
weight:   1      2     3    4 

capacity : 5

배낭 공간(capacity)이 5일때 아이템 넣을 때 최대 value 를 구해보자.

각 아이템을 넣었을떄 (true)/ 넣지 않았을떄(false) 로 생각
문제를 작은 문제로 나눠보자

                                                    ks("ABCD", weight:5)
D를 넣는 경우 D값+10, ks("ABC",1,(5-D무게=1)                                                    D를 넣지 않는 경우 값+0, ks("ABC",5)

C를 넣는 경우 C값+40 ks("AB", -2(종료))   C를 안넣는 경우 C값+0 ks("AB", 1-0)     C를 넣는 경우 C값+40 ks("AB", 2)          C를 안넣는 경우 C값+0 ks("AB", 5)

점화식 만들어 보자

w = weight
n = n개의 아이템, n는 몇번째 아이템까지 있는지(0~n)

case1) n번째 아이템이 있는 경우 = ks(n-1, w-w[n])+val[n] ==> 기존무게 - n번째아이템 무게를 빼줘야 한다.
case2) n번째 아이템이 없는 경우 = ks(n-1, w)+0

ks(n, w) = max (case1, case2)

DynamicProgramming 사용하기 표를 만들어 보자.
오른쪽 아래 대각선 방향은 물건이 있는 경우 +이전 value +추가한 물건 value
세로 아래 방향은 물건이 없는 경우 +이전 value 를 그대로 가져온다.(물건은 더 넣을 수 없는 경우도 이 경우의 값을 사용한다.)
n\w    0  1  2  3  4  5
0      0  0  0  0  0  0
1A     0  30 30 30 30 30
2AB    0  30 30 50 50 50
3ABC   0  30 30 50 70 70
4ABCD  0  30 30 50 70 70

O(nw)

*/

#include <iostream>
#include <string.h> // for memset

using namespace std;

const int maxLen = 100;
int recursiveKnapSack(int DP[][maxLen], int v[], int vLength, int w[], int cap, int idx)
{
    if (cap <= 0 || idx < 0 || idx >= vLength)
    {
        return 0;
    }
    if (DP[idx][cap] > 0)
    {
        return DP[idx][cap];
    }
    int case1 = 0;
    int case2 = 0;
    // 대각선 아래 방향(idx 아이템을 넣은 경우)
    if (w[idx] <= cap)
    {
        case1 = recursiveKnapSack(DP, v, vLength, w, cap - w[idx], idx + 1) + v[idx];
    }
    case2 = recursiveKnapSack(DP, v, vLength, w, cap, idx + 1);

    DP[idx][cap] = std::max(case1, case2);
    return DP[idx][cap];
}

int knapSack(int profits[], int profitsLength, int weights[], int capacity)
{
    int DP[maxLen][maxLen];
    memset(DP, 0, sizeof(int) * maxLen * maxLen);
    return recursiveKnapSack(DP, profits, profitsLength, weights, capacity, 0);
}

int main()
{
    int profits[] = {30, 20, 40, 10};
    int weights[] = {1, 2, 3, 4};
    cout << "Total knapsack profit = " << knapSack(profits, 4, weights, 5) << endl;

    int profits1[] = {1, 6, 10, 16};
    int weights1[] = {1, 2, 3, 5};
    cout << "Total knapsack profit = " << knapSack(profits1, 4, weights1, 7) << endl;
    cout << "Total knapsack profit = " << knapSack(profits1, 4, weights1, 6) << endl;
}
