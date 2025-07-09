// https://www.hackerrank.com/challenges/contacts/problem?isFullScreen=true
/*
We're going to make our own Contacts application! The application must perform two types of operations:

add name, where  is a string denoting a contact name. This must store  as a new contact in the application.
find partial, where  is a string denoting a partial name to search the application for. It must count the number of contacts starting with  and print the count on a new line.
Given  sequential add and find operations, perform each operation in order.

Input Format

The first line contains a single integer, , denoting the number of operations to perform.
Each line  of the  subsequent lines contains an operation in one of the two forms defined above.

Constraints

It is guaranteed that  and  contain lowercase English letters only.
The input doesn't have any duplicate  for the  operation.
Output Format

For each find partial operation, print the number of contact names starting with  on a new line.

Sample Input

4
add hack
add hackerrank
find hac
find hak


Sample Output

2
0
Explanation

We perform the following sequence of operations:

Add a contact named hack.
Add a contact named hackerrank.
Find and print the number of contact names beginning with hac. There are currently two contact names in the application and both of them start with hac, so we print  on a new line.
Find and print the number of contact names beginning with hak. There are currently two contact names in the application but neither of them start with hak, so we print  on a new line.
*/
// #include <bits/stdc++.h>
#include <iostream>
#include <vector>

using namespace std;

/*
 * Complete the contacts function below.
 */

// 다음 에러는 g++ -std=c++11 을 사용하면 된다.
// space is required between consecutive right angle brackets (use '> >')
// vector<int> contacts(vector<vector<string>> queries)

// Trie 구성을 위한 노드 구조
struct Node {
    int   cnt;        // 현재 노드까지 왔을때 카운트증가
    Node *child[26];  // 26개의 알파벳 자식 노드를 가지고 있다.
    Node() {
        cnt = 0;
        // 포인터 변수 자체 주소값을 참조해서 초기화 해야 한다.
        for (auto &i : child) {
            i = NULL;
        }
    }
};

void add(Node *root, string str) {
    Node *node = root;
    for (auto s : str) {
        // cout << s << "  " << s - 97 << endl;
        // 26개의 알파벳 자식 노드에 카운트를 증가한다.
        if (node->child[s - 97] == NULL) {
            // Node new_node 로 로컬 스택에 생성하면
            // 메모리 해제로 인한 bus error 가 발생한다.
            Node *new_node      = new Node();
            node->child[s - 97] = new_node;
        }
        node = node->child[s - 97];
        ++node->cnt;
    }
}

int find(Node *root, string str) {
    Node *node = root;
    for (auto s : str) {
        // cout << s << "  " << s - 97 << endl;
        if (node->child[s - 97] == NULL) {
            return 0;
        }
        node = node->child[s - 97];
    }
    if (node == NULL) {
        return 0;
    }
    // 마지막 까지 도착하면 str 가 있는 전체 원드들의 카운트가 된다.
    return node->cnt;
}
vector<int> contacts(vector<vector<string>> queries) {
    /*
     * Write your code here.
     */
    // auto 사용을 위해선 g++ -std=c++11
    vector<int>              result;
    vector<string>::iterator iter;
    Node                     root;
    for (auto i : queries) {
        for (iter = i.begin(); iter != i.end(); ++iter) {
            // 어떤 명령을 입력했는지 파악
            if (*iter == "add") {
                ++iter;
                add(&root, *iter);
            } else if (*iter == "find") {
                ++iter;
                int cnt = find(&root, *iter);
                // cout << "cnt: " << cnt << endl;
                result.push_back(cnt);
            }
        }
    }

    return result;
}

int main() {
    // 로컬 컴파일시 사용하지 않음.
    // ofstream fout(getenv("OUTPUT_PATH"));

    int queries_rows;
    cin >> queries_rows;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    vector<vector<string>> queries(queries_rows);
    for (int queries_row_itr = 0; queries_row_itr < queries_rows; queries_row_itr++) {
        queries[queries_row_itr].resize(2);

        for (int queries_column_itr = 0; queries_column_itr < 2; queries_column_itr++) {
            cin >> queries[queries_row_itr][queries_column_itr];
        }

        cin.ignore(numeric_limits<streamsize>::max(), '\n');
    }

    vector<int> result = contacts(queries);

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
