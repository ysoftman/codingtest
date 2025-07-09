// https://www.hackerrank.com/challenges/tree-level-order-traversal/problem
/*
You are given a pointer to the root of a binary tree. You need to print the level order traversal of this tree. In level order traversal, we visit the nodes level by level from left to right. You only have to complete the function. For example:

     1
      \
       2
        \
         5
        /  \
       3    6
        \
         4
For the above tree, the level order traversal is 1 -> 2 -> 5 -> 3 -> 6 -> 4.

Input Format

You are given a function,

void levelOrder(Node * root) {

}
Constraints

1 Nodes in the tree  500

Output Format

Print the values in a single line separated by a space.

Sample Input

     1
      \
       2
        \
         5
        /  \
       3    6
        \
         4
Sample Output

1 2 5 3 6 4

Explanation

We need to print the nodes level by level. We process each level from left to right.
Level Order Traversal: 1 -> 2 -> 5 -> 3 -> 6 -> 4.
*/
// #include <bits/stdc++.h>
#include <iostream>
#include <queue>

using namespace std;

class Node {
   public:
    int   data;
    Node *left;
    Node *right;
    Node(int d) {
        data  = d;
        left  = NULL;
        right = NULL;
    }
};

class Solution {
   public:
    Node *insert(Node *root, int data) {
        if (root == NULL) {
            return new Node(data);
        } else {
            Node *cur;
            if (data <= root->data) {
                cur        = insert(root->left, data);
                root->left = cur;
            } else {
                cur         = insert(root->right, data);
                root->right = cur;
            }

            return root;
        }
    }
    /*
class Node {
    public:
        int data;
        Node *left;
        Node *right;
        Node(int d) {
            data = d;
            left = NULL;
            right = NULL;
        }
};
*/

    void levelOrder(Node *root) {
        // BFS(너비 우선 탐색)
        queue<Node *> q;
        q.push(root);
        while (not q.empty()) {
            // 공백으로 출력값들 구분
            if (q.front()->data != root->data)
                cout << " ";
            cout << q.front()->data;
            if (q.front()->left != NULL)
                q.push(q.front()->left);
            if (q.front()->right != NULL)
                q.push(q.front()->right);
            q.pop();
        }
        cout << endl;
    }

};  // End of Solution

int main() {
    Solution myTree;
    Node    *root = NULL;

    int      t;
    int      data;

    std::cin >> t;

    while (t-- > 0) {
        std::cin >> data;
        root = myTree.insert(root, data);
    }

    myTree.levelOrder(root);
    return 0;
}

/*
6 # 노드 개수 입력
1 2 5 3 6 4 # 노드값 입력
# 출력
1 2 5 3 6 4

*/