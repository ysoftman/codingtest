/*
ysoftman

https://www.geeksforgeeks.org/flatten-bst-to-sorted-list-increasing-order/

Given a binary search tree, the task is to flatten it to a sorted list. Precisely, the value of each node must be lesser than the values of all the nodes at its right, and its left node must be NULL after flattening. We must do it in O(H) extra space where ‘H’ is the height of BST.

Examples:

Input:
          5
        /   \
       3     7
      / \   / \
     2   4 6   8
Output: 2 3 4 5 6 7 8
Input:
      1
       \
        2
         \
          3
           \
            4
             \
              5
Output: 1 2 3 4 5
*/

#include <algorithm>
#include <iostream>
#include <queue>

using namespace std;

struct node {
    int   data;
    node *left;
    node *right;
    node(int data) {
        this->data = data;
        left       = NULL;
        right      = NULL;
    }
};

void printBST(node *root) {
    cout << "BST" << endl;
    if (root == NULL) {
        cout << "null node!" << endl;
        return;
    }

    queue<node *> q;
    q.push(root);
    while (not q.empty()) {
        node *cur = q.front();
        cout << cur->data << endl;
        if (cur->left != NULL) {
            q.push(cur->left);
        }
        if (cur->right != NULL) {
            q.push(cur->right);
        }
        q.pop();
    }
}

void printVector(vector<int> result) {
    for (auto i : result) {
        cout << i << " ";
    }
    cout << endl;
}

void printnode_inOrder(node *root, vector<int> &result) {
    // -1 은 null 로 처리하지 않는다.
    if (root->data < 0)
        return;

    if (root->left != NULL)
        printnode_inOrder(root->left, result);
    // cout << root->data << " ";
    result.push_back(root->data);
    if (root->right != NULL)
        printnode_inOrder(root->right, result);
}

void flattenBST_using_sort(node *root) {
    cout << "flattenBST_using_sort" << endl;
    if (root == NULL) {
        cout << "null node!" << endl;
        return;
    }

    vector<int>   v;
    queue<node *> q;
    q.push(root);
    while (not q.empty()) {
        node *cur = q.front();
        v.push_back(cur->data);
        if (cur->left != NULL) {
            q.push(cur->left);
        }
        if (cur->right != NULL) {
            q.push(cur->right);
        }
        q.pop();
    }

    std::sort(v.begin(), v.end());
    printVector(v);
}

int main() {
    node *root         = new node(5);
    root->left         = new node(3);
    root->right        = new node(7);
    root->left->left   = new node(2);
    root->left->right  = new node(4);
    root->right->left  = new node(6);
    root->right->right = new node(8);

    printBST(root);
    flattenBST_using_sort(root);
    cout << "flattenBST_using_inorder" << endl;
    vector<int> result;
    printnode_inOrder(root, result);
    printVector(result);
}
