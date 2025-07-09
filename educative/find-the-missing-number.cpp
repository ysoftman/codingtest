/*
https://www.educative.io/m/find-the-missing-number
You are given an array of positive numbers from 1 to n, such that all numbers from 1 to n are present except one number ‘x’. You have to find ‘x’. The input array is not sorted.

For example, let’s look at the below array.
n = 8
3 7 1 2 8 4 5
missing number = 6
*/

#include <iostream>
#include <vector>

using namespace std;

int find_missing(const vector<int> &input) {
    // TODO: Write - Your - Code
    int sum = 0;
    int max = 0;
    for (auto i : input) {
        if (max < i) {
            max = i;
        }
        sum += i;
    }
    int correct_sum = 0;
    for (int i = 1; i <= max; i++) {
        correct_sum += i;
    }
    cout << sum << endl;
    cout << correct_sum << endl;

    return correct_sum - sum;
}

void test(int n) {
    int         missing_element = rand() % n + 1;
    vector<int> v;
    for (int i = 1; i <= n; ++i) {
        if (i == missing_element) {
            continue;
        }
        v.push_back(i);
    }
    int actual_missing = find_missing(v);
    cout << "Expected Missing = " << missing_element << " Actual Missing = " << actual_missing << endl;
}

int main() {
    srand(time(NULL));
    for (int i = 1; i < 10; ++i)
        test(100000);

    return 0;
}
