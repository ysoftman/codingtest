// https://www.hackerrank.com/challenges/find-the-running-median/problem
/*
The median of a set of integers is the midpoint value of the data set for which an equal number of integers are less than and greater than the value. To find the median, you must first sort your set of integers in non-decreasing order, then:

If your set contains an odd number of elements, the median is the middle element of the sorted sample. In the sorted set ,  is the median.
If your set contains an even number of elements, the median is the average of the two middle elements of the sorted sample. In the sorted set ,  is the median.
Given an input stream of  integers, you must perform the following task for each  integer:

Add the  integer to a running list of integers.
Find the median of the updated list (i.e., for the first element through the  element).
Print the list's updated median on a new line. The printed value must be a double-precision number scaled to  decimal place (i.e.,  format).
Input Format

The first line contains a single integer, , denoting the number of integers in the data stream.
Each line  of the  subsequent lines contains an integer, , to be added to your list.

Constraints

Output Format

After each new integer is added to the list, print the list's updated median on a new line as a single double-precision number scaled to  decimal place (i.e.,  format).

Sample Input

6  #입력 개수
12
4
5
3
8
7
Sample Output

12.0
8.0
5.0
4.5
5.0
6.0
Explanation

There are  integers, so we must print the new median on a new line as each integer is added to the list:
list = {12}, median = 12.0
list = {12,4} -> {4,12} median = (12+4)/2 = 8.0
list = {12,4,5} -> {4,5,12} median = 5.0
list = {12,4,5,3} -> {3,4,5,12} median = (4+5)/2 = 4.5
list = {12,4,5,3,8} -> {3,4,5,8,12} median = 5.0
list = {12,4,5,3,8,7} -> {3,4,5,7,8,12} median = (5+7)/2 = 6.0
*/
// #include <bits/stdc++.h>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

/*
 * Complete the runningMedian function below.
 */
vector<double> runningMedian(vector<int> a) {
    /*
     * Write your code here.
     */
    // max-heap(부모노드값이 자식노드값보다 큰 경우)
    priority_queue<int, vector<int>, less<int>> pq_max;
    // min-heap(부모노드값이 자식노드값보다 작은 경우)
    priority_queue<int, vector<int>, greater<int>> pq_min;
    vector<double>                                 result;
    vector<int>                                    temp;
    // g++ -std=c++11
    for (auto b : a) {
        // (min-max)heap 을 이용하면 원소가 추가될때마다 전체 정렬을 하지 않고
        // O(logN) 만큼의 시간 복잡도로 정렬을 유지할 수 있다.
        // pq_max 를 왼쪽에 , pq_min 을 오른쪽에 있다고 생각하자.
        // pq_max 의 root 노드보다 크다면 pq_max(오른쪽)에 담는다.
        if (pq_max.size() > 0 && pq_max.top() < b)
            pq_min.push(b);
        else
            pq_max.push(b);

        // 한쪽 힙에 너무 많은 원소가 들어가 있지 않도록
        // 양쪽 힙의 원소 개수의 균형을 유지한다.
        // 양쪽 힙의 원소 개수가 2개이상 차이가 나지 않도록 한다.
        if (pq_max.size() > pq_min.size() + 1) {
            pq_min.push(pq_max.top());
            pq_max.pop();
        } else if (pq_max.size() + 1 < pq_min.size()) {
            pq_max.push(pq_min.top());
            pq_min.pop();
        }

        if (pq_max.size() == pq_min.size()) {
            // 2개의 root 노드는 전체의 중간 위치의 값들이 된다.
            result.push_back((pq_max.top() + pq_min.top()) / 2.0);
        } else {
            // 홀수의 경우 많은쪽의 루트 값이 중간값이다.
            if (pq_max.size() > pq_min.size())
                result.push_back(pq_max.top());
            else
                result.push_back(pq_min.top());
        }
    }
    return result;
}

int main() {
    // ofstream fout(getenv("OUTPUT_PATH"));

    int a_count;
    cin >> a_count;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    vector<int> a(a_count);

    for (int a_itr = 0; a_itr < a_count; a_itr++) {
        int a_item;
        cin >> a_item;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        a[a_itr] = a_item;
    }

    vector<double> result = runningMedian(a);

    for (int result_itr = 0; result_itr < result.size(); result_itr++) {
        // fout << result[result_itr];
        cout << result[result_itr];

        if (result_itr != result.size() - 1) {
            // fout << "\n";
            cout << endl;
        }
    }

    // fout << "\n";
    cout << endl;

    // fout.close();

    return 0;
}
