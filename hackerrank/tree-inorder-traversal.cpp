// https://www.hackerrank.com/challenges/tree-inorder-traversal/problem
/*
Complete the inOrder function in your editor below, which has  parameter: a
pointer to the root of a binary tree. It must print the values in the tree's
inorder traversal as a single line of space-separated values.

Input Format

Our hidden tester code passes the root node of a binary tree to your inOrder
function.

Constraints

1 Nodes in the tree  500

Output Format

Print the tree's inorder traversal as a single line of space-separated values.

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

1 2 3 4 5 6
*/
// #include <bits/stdc++.h>
#include <iostream>

using namespace std;

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

class Solution {
public:
  Node *insert(Node *root, int data) {
    if (root == NULL) {
      return new Node(data);
    } else {
      Node *cur;
      if (data <= root->data) {
        cur = insert(root->left, data);
        root->left = cur;
      } else {
        cur = insert(root->right, data);
        root->right = cur;
      }

      return root;
    }
  }

  /* you only have to complete the function given below.
  Node is defined as

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

  // inorder traversal : 왼쪽 자식노드 substree -> 부모 -> 오른쪽 자식노드
  // subtree 순으로
  void inOrder(Node *root) {
    if (root->left != NULL)
      inOrder(root->left);
    cout << root->data << " ";
    if (root->right != NULL)
      inOrder(root->right);
  }

}; // End of Solution

int main() {

  Solution myTree;
  Node *root = NULL;

  int t;
  int data;

  std::cin >> t;

  while (t-- > 0) {
    std::cin >> data;
    root = myTree.insert(root, data);
  }

  myTree.inOrder(root);
  return 0;
}

/*
// 테스트
6 # 노드개수 입력
1 2 5 3 6 4 # 노드값
# 결과
1 2 3 4 5 6

15 # 노드개수 입력
1 14 3 7 4 5 15 6 13 10 11 2 12 8 9 # 노드 값
# 결과
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
*/