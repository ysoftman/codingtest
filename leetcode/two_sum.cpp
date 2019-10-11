// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/546/
/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// 원소는 한번씩만 보고, O(n) 으로 끝내기
vector<int> twoSum(vector<int> &nums, int target)
{
    vector<int> result;
    // hashmap O(1)
    unordered_map<int, int> hashmap;
    for (int i = 0; i < nums.size(); ++i)
    {
        int difference = target - nums[i];
        // 현재 원소가 대상값이 되려면 필요한 값이 hashmap에서 찾아본다.
        if (hashmap.find(difference) != hashmap.end())
        {
            result.push_back(hashmap[difference]);
            result.push_back(i);
            return result;
        }
        // 원소마다 대상값과의 차이를 key로 해서 hashmap 으로 저장한다.
        hashmap[nums[i]] = i;
    }

    return result;
}

int main()
{
    vector<int> input = {2, 7, 11, 15};
    vector<int> r = twoSum(input, 9);
    for (auto i : r)
    {
        cout << i << endl;
    }
    return 0;
}