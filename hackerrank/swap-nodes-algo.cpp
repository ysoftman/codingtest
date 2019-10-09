// https://www.hackerrank.com/challenges/swap-nodes-algo/problem
/*
3
2 3
-1 -1
-1 -1
2
1
1
*/

#include <bits/stdc++.h>
#include <iostream>
#include <queue>

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

// inorder traversal : 왼쪽 자식노드 substree -> 부모 -> 오른쪽 자식노드
// subtree 순으로
void printnode_inOrder(Node *root, vector<int> &result) {
  // -1 은 null 로 처리하지 않는다.
  if (root->data < 0)
    return;

  if (root->left != NULL)
    printnode_inOrder(root->left, result);
  cout << root->data << " ";
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
    Node *temp = node->left;
    node->left = node->right;
    node->right = temp;
  }
  swap_k_depth(node->left, depth + 1, k);
  swap_k_depth(node->right, depth + 1, k);
}

vector<vector<int>> swapNodes(vector<vector<int>> indexes,
                              vector<int> queries) {

  vector<vector<int>> result;

  Node root(1);
  Node *node = &root;

  queue<Node *> q;
  q.push(node);

  // 이진트리 생성
  cout << "indexes" << endl;
  for (auto i : indexes) {
    node = q.front();
    for (auto j : i) {
      cout << j << "  ";
      Node *nn = new Node(j);
      if (node->left == NULL) {
        node->left = nn;
      } else {
        node->right = nn;
      }
      q.push(nn);
    }
    cout << endl;
    q.pop();
  }
  cout << endl;

  // for debugging
  // vector<int> current_result;
  // printnode_inOrder(&root, current_result);

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
