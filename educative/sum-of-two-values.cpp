/*
https://www.educative.io/m/sum-of-two-values

Given an array of integers and a value, determine if there are any two integers in the array whose sum is equal to the given value. Return true if the sum exists and return false if it does not.

Consider this array and the target sums:

5 7 1 2 8 4 3
target sum 10 => 7+3=10, 2+8=10
target sum 19 => no 2 values sum upto 19
*/

#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

bool find_sum_of_two(vector<int> &A, int val) {
    // TODO: Write - Your - Code
    unordered_set<int> hashset;
    for (auto a : A) {
        if (hashset.find(val - a) != hashset.end()) {
            return true;
        }
        hashset.insert(a);
    }
    return false;
}

int main() {
    vector<int> v    = {5, 7, 1, 2, 8, 4, 3};
    vector<int> test = {3, 20, 1, 2, 7};

    for (int i = 0; i < test.size(); i++) {
        bool output = find_sum_of_two(v, test[i]);
        cout << "find_sum_of_two(v, " << test[i] << ") = " << (output ? "true" : "false") << endl;
    }
    return 0;
}
