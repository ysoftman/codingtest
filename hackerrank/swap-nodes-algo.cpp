// https://www.hackerrank.com/challenges/swap-nodes-algo/problem
/*
A binary tree is a tree which is characterized by one of the following
properties:

It can be empty (null).
It contains a root node only.
It contains a root node with a left subtree, a right subtree, or both. These
subtrees are also binary trees. In-order traversal is performed as

Traverse the left subtree.
Visit root.
Traverse the right subtree.
For this in-order traversal, start from the left child of the root node and keep
exploring the left subtree until you reach a leaf. When you reach a leaf, back
up to its parent, check for a right child and visit it if there is one. If there
is not a child, you've explored its left and right subtrees fully. If there is a
right child, traverse its left subtree then its right in the same manner. Keep
doing this until you have traversed the entire tree. You will only store the
values of a node as you visit when one of the following is true:

it is the first node visited, the first time visited
it is a leaf, should only be visited once
all of its subtrees have been explored, should only be visited once while this
is true it is the root of the tree, the first time visited Swapping: Swapping
subtrees of a node means that if initially node has left subtree L and right
subtree R, then after swapping, the left subtree will be R and the right
subtree, L.

For example, in the following tree, we swap children of node 1.

                                Depth
    1               1            [1]
   / \             / \
  2   3     ->    3   2          [2]
   \   \           \   \
    4   5           5   4        [3]
In-order traversal of left tree is 2 4 1 3 5 and of right tree is 3 5 1 2 4.

Swap operation:

We define depth of a node as follows:

The root node is at depth 1.
If the depth of the parent node is d, then the depth of current node will be
d+1. Given a tree and an integer, k, in one operation, we need to swap the
subtrees of all the nodes at each depth h, where h ∈ [k, 2k, 3k,...]. In other
words, if h is a multiple of k, swap the left and right subtrees of that level.

You are given a tree of n nodes where nodes are indexed from [1..n] and it is
rooted at 1. You have to perform t swap operations on it, and after each swap
operation print the in-order traversal of the current state of the tree.

Function Description

Complete the swapNodes function in the editor below. It should return a
two-dimensional array where each element is an array of integers representing
the node indices of an in-order traversal after a swap operation.

swapNodes has the following parameter(s):
- indexes: an array of integers representing index values of each node[i],
beginning with node[1], the first element, as the root.
- queries: an array of integers, each representing a k value.

Input Format
The first line contains n, number of nodes in the tree.

Each of the next n lines contains two integers, a b, where a is the index of
left child, and b is the index of right child of ith node.

Note: -1 is used to represent a null node.

The next line contains an integer, t, the size of queries.
Each of the next t lines contains an integer queries[i], each being a value k.

Output Format
For each k, perform the swap operation and store the indices of your in-order
traversal to your result array. After all swap operations have been performed,
return your result array for printing.

Constraints

Either  or
Either  or
The index of a non-null child will always be greater than that of its parent.
Sample Input 0

3 # 입력할 노드(left,right 자식노드) 개수, root 노드는 디폴틀 1로 존재
2 3 # 1의 자식노드 left:2 right:3 노드
-1 -1 # 2의 자식노드 left:-1 right:-1 노드(-1는 NULL)
-1 -1 # 3의 자식노드 left:-1 right:-1 노드(-1는 NULL)
2 # swap 수행 회수
1 # depth 1 의 배수마다 swap 수행
1 # depth 1 의 배수마다 swap 수행
Sample Output 0

3 1 2
2 1 3
Explanation 0

As nodes 2 and 3 have no children, swapping will not have any effect on them. We
only have to swap the child nodes of the root node.

    1   [s]       1    [s]       1
   / \      ->   / \        ->  / \
  2   3 [s]     3   2  [s]     2   3
Note: [s] indicates that a swap operation is done at this depth.

Sample Input 1

5
2 3
-1 4
-1 5
-1 -1
-1 -1
1
2
Sample Output 1

4 2 1 5 3
Explanation 1

Swapping child nodes of node 2 and 3 we get

    1                  1
   / \                / \
  2   3   [s]  ->    2   3
   \   \            /   /
    4   5          4   5
Sample Input 2

11
2 3
4 -1
5 -1
6 -1
7 8
-1 9
-1 -1
10 11
-1 -1
-1 -1
-1 -1
2
2
4
Sample Output 2

2 9 6 4 1 3 7 5 11 8 10
2 6 9 4 1 3 7 5 10 8 11
Explanation 2

Here we perform swap operations at the nodes whose depth is either 2 or 4 for
and then at nodes whose depth is 4 for .

         1                     1                          1
        / \                   / \                        / \
       /   \                 /   \                      /   \
      2     3    [s]        2     3                    2     3
     /      /                \     \                    \     \
    /      /                  \     \                    \     \
   4      5          ->        4     5          ->        4     5
  /      / \                  /     / \                  /     / \
 /      /   \                /     /   \                /     /   \
6      7     8   [s]        6     7     8   [s]        6     7     8
 \          / \            /           / \              \         / \
  \        /   \          /           /   \              \       /   \
   9      10   11        9           11   10              9     10   11
*/

// g++ -std=c++11 ./swap-nodes-algo.cpp
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

// preorder traversal : 부모 -> 왼쪽 자식 -> 오른쪽 자식 우선순위로
void printnode_preOrder(Node *root, vector<int> &result) {
    // -1 은 null 로 처리하지 않는다.
    if (root->data < 0)
        return;

    cout << root->data << " ";
    result.push_back(root->data);
    if (root->left != NULL)
        printnode_preOrder(root->left, result);
    if (root->right != NULL)
        printnode_preOrder(root->right, result);
}

// inorder traversal : 왼쪽 자식 -> 부모 -> 오른쪽 자식 우선순위로
// subtree 순으로
void printnode_inOrder(Node *root, vector<int> &result) {
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

void swap_k_depth(Node *node, int depth, int k) {
    if (node == NULL) {
        return;
    }
    // k depth 배수때 마다 자식노드 left, right 를 스왑
    if (depth % k == 0) {
        Node *temp  = node->left;
        node->left  = node->right;
        node->right = temp;
    }
    swap_k_depth(node->left, depth + 1, k);
    swap_k_depth(node->right, depth + 1, k);
}

vector<vector<int>> swapNodes(vector<vector<int>> indexes,
                              vector<int>         queries) {
    vector<vector<int>> result;
    Node                root(1);
    Node               *node = &root;
    queue<Node *>       q;
    q.push(node);

    // 이진트리 생성
    cout << "indexes" << endl;
    for (auto i : indexes) {
        node = q.front();
        for (auto j : i) {
            cout << j << " ";
            Node *nn = new Node(j);
            if (node->left == NULL) {
                node->left = nn;
            } else {
                node->right = nn;
            }
            // -1 left 노드는 queue 에 넣지 않는다.
            if (j != -1)
                q.push(nn);
        }
        cout << endl;
        q.pop();
    }
    cout << endl;

    // for debugging
    // vector<int> temp_result;
    // printnode_preOrder(&root, temp_result);

    // queries : 스왑을 실행한 트리 depth 값들
    cout << "queries" << endl;
    for (auto k : queries) {
        cout << k << "  " << endl;
        swap_k_depth(&root, 1, k);
        vector<int> current_result;
        printnode_inOrder(&root, current_result);
        result.push_back(current_result);
    }
    cout << endl;

    return result;
}

int main() {
    // ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    vector<vector<int>> indexes(n);
    for (int indexes_row_itr = 0; indexes_row_itr < n; indexes_row_itr++) {
        indexes[indexes_row_itr].resize(2);

        for (int indexes_column_itr = 0; indexes_column_itr < 2;
             indexes_column_itr++) {
            cin >> indexes[indexes_row_itr][indexes_column_itr];
        }

        cin.ignore(numeric_limits<streamsize>::max(), '\n');
    }

    int queries_count;
    cin >> queries_count;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    vector<int> queries(queries_count);

    for (int queries_itr = 0; queries_itr < queries_count; queries_itr++) {
        int queries_item;
        cin >> queries_item;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        queries[queries_itr] = queries_item;
    }

    vector<vector<int>> result = swapNodes(indexes, queries);

    for (int result_row_itr = 0; result_row_itr < result.size();
         result_row_itr++) {
        for (int result_column_itr = 0;
             result_column_itr < result[result_row_itr].size();
             result_column_itr++) {
            // fout << result[result_row_itr][result_column_itr];
            cout << result[result_row_itr][result_column_itr];

            if (result_column_itr != result[result_row_itr].size() - 1) {
                // fout << " ";
                cout << " ";
            }
        }

        if (result_row_itr != result.size() - 1) {
            // fout << "\n";
            cout << endl;
        }
    }

    // fout << "\n";
    cout << endl;

    // fout.close();

    return 0;
}
