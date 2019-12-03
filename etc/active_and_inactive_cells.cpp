// https://www.tutorialspoint.com/active-and-inactive-cells-after-k-days

/*
# 이웃하는 양쪽의 셀이 같으면 0, 아니면 1
# 양쪽 끝의 셀 2개는 끝에 0 이 있다고 가정
# 입력
day = 3
01010101
# 결과
10000000  # 1일
01000000  # 2일
10100000  # 3일(최종 결과)
*/

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void printCells(vector<int> vec)
{
    for (int i = 0; i < vec.size(); ++i)
        cout << vec[i];
    cout << endl;
}

vector<int> active_and_inactive_cells(int cells[], int len, int days)
{
    vector<int> result(cells, cells + len);
    vector<int> temp;
    // 이웃하는 양쪽의 셀이 같으면 0, 아니면 1
    // 양쪽 끝의 셀 2개는 끝에 0 이 있다고 가정
    for (int d = 0; d < days; ++d)
    {
        int left = -1, right = 1;
        for (int i = 0; i < len; ++i)
        {
            // 값이 같으면 0 인 xor 성질을 이용하자.
            if (i == 0)
            {
                temp.push_back(0 ^ result[right]);
            }
            else if (i == len - 1)
            {
                temp.push_back(result[left] ^ 0);
            }
            else
            {
                temp.push_back(result[left] ^ result[right]);
            }
            ++left, ++right;
        }
        result.clear();
        result.assign(temp.begin(), temp.end());
        // printCells(result);
        temp.clear();
    }
    return result;
}

int main()
{
    int days = 3;
    int len = 8;
    int cells[] = {0, 1, 0, 1, 0, 1, 0, 1};
    for (int i = 0; i < len; ++i)
    {
        cout << cells[i];
    }
    cout << endl;
    vector<int> r;
    r = active_and_inactive_cells(cells, len, days);
    printCells(r);
}
