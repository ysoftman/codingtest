/*
https://leetcode.com/problems/valid-anagram/
242. Valid Anagram
Easy
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
*/

// g++ -std=c++11 valid_anagram.cpp && ./a.out 
#include <iostream>
#include <unordered_map>

using std::cout;
using std::endl;
using std::string;
using std::unordered_map;

class Solution
{
public:
    bool isAnagram(string s, string t)
    {
        if (s.length() != t.length())
        {
            return false;
        }
        unordered_map<char, int> hashmap;
        for (auto c : s)
        {
            hashmap[c]++;
        }
        // for (auto m : hashmap) {
        //     cout << m.first << ":" << m.second;
        // }
        // cout << endl;
        // unordered_map<char, int>::iterator itera;
        // for (iter = hashmap.begin(); iter != hashmap.end(); iter++) {
        //     cout << iter->first << ":" << iter->second;
        // }
        // cout << endl;
        // cout << endl;

        for (auto c : t)
        {
            if (hashmap.find(c) == hashmap.end())
            {
                return false;
            }
            if (hashmap[c] <= 0)
            {
                return false;
            }
            hashmap[c]--;
        }
        return true;
    }
};

int main()
{
    Solution *sol = new Solution();
    cout << sol->isAnagram("anagram", "nagaram") << endl;
    cout << sol->isAnagram("rat", "car") << endl;
    cout << sol->isAnagram("aacc", "ccac") << endl;
    delete sol;
    return 0;
}
