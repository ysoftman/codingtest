// ysoftman
// g++ -std=c++11 sort_mutli_list.cpp && ./a.out
//
// ex) [1, 3, 6, 10, 12], [2, 8, 11], [4, 7, 9], [1, 2, 3, 4]
// => [1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10, 11, 12]
#include <iostream>
#include <vector>
#include <list>
#include <map>

using namespace std;

vector<int> sort_sortedlist(vector<list<int>> &lists)
{
    vector<int> out;

    map<int, int> tempMap;
    for (auto i : lists)
    {
        for (auto j : i)
        {
            // cout << j << endl;
            tempMap[j]++;
        }
    }
    for (auto i : tempMap)
    {
        for (int j = 0; j < i.second; ++j)
        {
            // cout << i.first << endl;
            out.push_back(i.first);
        }
    }
    // for debugging...
    vector<int>::iterator outIter;
    for (outIter = out.begin(); outIter != out.end(); ++outIter)
    {
        cout << *outIter;
        if (outIter != out.end() - 1)
        {
            cout << ", ";
        }
        else
        {
            cout << endl;
        }
    }
    return out;
}
int main()
{
    // vector< list<int > > vec;
    list<int> tempList;
    list<int>::iterator iter;
    vector<list<int>> vec;
    tempList = {1, 3, 6, 10, 12};
    vec.push_back(tempList);
    tempList = {2, 8, 11};
    vec.push_back(tempList);
    tempList = {4, 7, 9};
    vec.push_back(tempList);
    tempList = {1, 2, 3, 4};
    vec.push_back(tempList);
    sort_sortedlist(vec);
    return 0;
}
