/*
https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/546/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

// 원소는 한번씩만 보고, O(n) 으로 끝내기
vector<int> twoSum(vector<int>& nums, int target) {
    vector<int> result;
    // hashmap O(1)
    unordered_map<int, int> hashmap;
    for (int i = 0; i < nums.size(); ++i) {
        int difference = target - nums[i];
        // 현재 원소가 대상값이 되려면 필요한 값이 hashmap에서 찾아본다.
        if (hashmap.find(difference) != hashmap.end()) {
            result.push_back(hashmap[difference]);
            result.push_back(i);
            return result;
        }
        // 원소마다 대상값과의 차이를 key로 해서 hashmap 으로 저장한다.
        hashmap[nums[i]] = i;
    }

    return result;
}

// sort array, and use two pointer, O(logN)
vector<int> twoSum2(vector<int>& nums, int target) {
    struct n {
        int val;
        int idx;
        n(int a, int b) : val(a), idx(b) {}
    };
    vector<n> data;
    for (size_t i = 0; i < nums.size(); i++) {
        data.emplace_back(nums[i], i);
    }
    // for (auto& i : data) {
    //     cout << i.val << "--" << i.idx << endl;
    // }
    sort(data.begin(), data.end(), [](const n& a, const n& b) {
        return a.val < b.val;
    });
    int i = 0;
    int j = data.size() - 1;
    for (; i < j;) {
        if (data.at(i).val + data.at(j).val == target) {
            return vector<int>{data.at(i).idx, data.at(j).idx};
        }
        if (data.at(i).val + data.at(j).val < target) {
            i++;
        } else if (data.at(i).val + data.at(j).val > target) {
            j--;
        }
    }
    return vector<int>{};
}

int main() {
    vector<int> input = {2, 7, 11, 15};
    cout << "twoSum" << endl;
    vector<int> r = twoSum(input, 9);
    for (const auto& i : r) {
        cout << i << endl;
    }
    cout << "twoSum2" << endl;
    vector<int> r2 = twoSum2(input, 9);
    for (const auto& i : r2) {
        cout << i << endl;
    }
    return 0;
}
