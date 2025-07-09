// https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree
/*
The height of a binary tree is the number of edges between the tree's root and its furthest leaf. For example, the following binary tree is of height :
     4
  2    6
1 3    5 7

Function Description

Complete the getHeight or height function in the editor. It must return the height of a binary tree as an integer.

getHeight or height has the following parameter(s):

root: a reference to the root of a binary tree.
Note -The Height of binary tree with single node is taken as zero.

Input Format

The first line contains an integer , the number of nodes in the tree.
Next line contains  space separated integer where i-th integer denotes node[i].data.

Note: Node values are inserted into a binary search tree before a reference to the tree's root node is passed to your function. In a binary search tree, all nodes on the left branch of a node are less than the node value. All values on the right branch are greater than the node value.

Constraints
1 <= node.data[i] <= 20
1 <= n <= 20

Output Format

Your function should return a single integer denoting the height of the binary tree.

Sample Input

     3
  2     5
1     4  6
           7

Sample Output
3
Explanation

The longest root-to-leaf path is shown below:
     3
  2     5
1     4  6
           7

There are  nodes in this path that are connected by  edges, meaning our binary tree's

*/

// 프로그램 콘테스트에서만 사용할 수 있는 모든 standard library
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
    /*The tree node has data, left child and right child
class Node {
    int data;
    Node* left;
    Node* right;
};

*/
    int height(Node *root) {
        // Write your code here.
        // BFS(너비 우선 탐색)
        queue<Node *> q;
        // cout << root->data << endl;
        q.push(root);
        int height         = 0;
        int previous_nodes = 0;
        while (not q.empty()) {
            // cout << q.front()->data << " height: " << height << endl;
            if (q.front()->left != NULL) {
                q.push(q.front()->left);
            }
            if (q.front()->right != NULL) {
                q.push(q.front()->right);
            }
            if (previous_nodes != 2 and (q.front()->left != NULL or q.front()->right != NULL)) {
                ++height;
            }
            // 이전(부모) 노드에서 left,right 둘다 노드가 존재하는 경우 2번 카운트 되는것을 막기위해
            previous_nodes = 0;
            if (q.front()->left != NULL)
                ++previous_nodes;
            if (q.front()->right != NULL)
                ++previous_nodes;

            q.pop();
        }
        return height;
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

    int height = myTree.height(root);

    std::cout << height;

    return 0;
}

/* 입력 예시
7 # 그래프를 구성할 노드 개수
3 2 5 1 4 6 7 # 각 노드의 데이터
# 현재(부모) 노드값을 기준으로 작으면 왼쪽, 크면 오른쪽으로 배치된다.
    3
  2    5
1    4   6
           7
높이(edge) : 3

5
3 1 7 5 4
    3
  1   7
    5
  4
높이(edge) : 3


15
1 3 2 5 4 6 7 9 8 11 13 12 10 15 14
1
    3
 2     5
     4    6
            7
              9
            8    11
              10    13
                  12    15
                      14
높이(edge) : 9

7
3 5 2 1 4 6 7
     3
  2     5
1     4   6
            7
높이(edge) : 3
*/
