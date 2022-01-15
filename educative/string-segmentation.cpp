/*
https://www.educative.io/m/string-segmentation

Problem Statement
You are given a dictionary of words and a large input string. You have to find out whether the input string can be completely segmented into the words of a given dictionary. The following two examples elaborate on the problem further.

Given dictionary of words 
apple pear pier pie

Input string of "applepie" can be segmented into dictionary words (apple and pie)
appel pie

Input string OF "applepeer" cannot be segmented into dictionary words.
apple peer(x)

*/
#include <iostream>
#include <unordered_set>
using std::cout;
using std::endl;
using std::string;
using std::unordered_set;

unordered_set<string> dpFound;
bool can_segment_string(string s, unordered_set<string> &dictionary)
{
    //TODO: Write - Your - Code
    if (dpFound.find(s) != dpFound.end())
    {
        cout << "found:" << s << endl;
        return true;
    }
    for (int i = 0; i < s.length(); i++)
    {
        string temp = s.substr(0, i);
        if (dictionary.find(temp) != dictionary.end())
        {
            // 사전에 존재하는 단어 기록
            dpFound.insert(temp);
            if (dictionary.find(s.substr(i)) != dictionary.end())
            {
                return true;
            }
            return can_segment_string(s.substr(i), dictionary);
        }
    }
    return false;
}

int main()
{
    string s = "hellonow";
    unordered_set<string> dictonary = {"hello", "hell", "on", "now"};
    if (can_segment_string(s, dictonary))
        cout << "String Can be Segmented" << endl;
    else
        cout << "String Can NOT be Segmented" << endl;

    s = "applepie";
    dictonary = {"apple", "pear", "pier", "pie"};
    if (can_segment_string(s, dictonary))
        cout << "String Can be Segmented" << endl;
    else
        cout << "String Can NOT be Segmented" << endl;

    s = "applepeer";
    dictonary = {"apple", "pear", "pier", "pie"};
    if (can_segment_string(s, dictonary))
        cout << "String Can be Segmented" << endl;
    else
        cout << "String Can NOT be Segmented" << endl;
    return 0;
}
