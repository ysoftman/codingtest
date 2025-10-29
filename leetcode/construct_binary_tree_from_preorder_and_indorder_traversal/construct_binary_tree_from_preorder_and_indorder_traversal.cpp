/*
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
105. Construct Binary Tree from Preorder and Inorder Traversal
Medium
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]


Constraints:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
*/

/*
# preorder : 부모 -> 왼쪽 -> 오른쪽 순으로 순회
preorder = [3,9,20,15,7]
# inorder : 왼쪽 -> 부모 -> 오른쪽 순으로 순회
inorder = [9,3,15,20,7]

Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7

# 다음과 같은 포맷으로 입력
[3,9,20,15,7]
[9,3,15,20,7]

# 결과
[3,9,20,null,null,15,7]
*/
#include <iostream>
#include <queue>
#include <sstream>
#include <vector>

using namespace std;
struct TreeNode {
    int       val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

static int preorder_index = 0;

int        searchinorder(vector<int> inorder, int start, int end, int v) {
    for (int i = start; i <= end; ++i) {
        if (inorder[i] == v) {
            return i;
        }
    }
    return -1;
}
TreeNode* makeTree(vector<int> preorder, vector<int> inorder, int left, int right) {
    if (left > right) {
        return NULL;
    }

    // preorder 의 원소값을 inorder 에서 찾아 그 index 를 기준으로
    // 왼쪽부분에서는 left 오른쪽부분에서 right 노드를 재귀로 찾아간다.
    struct TreeNode* node = new TreeNode(preorder[preorder_index++]);
    if (left == right) {
        return node;
    }
    int i = searchinorder(inorder, left, right, node->val);
    // if (i < 0)
    // {
    //     return NULL;
    // }
    node->left  = makeTree(preorder, inorder, left, i - 1);
    node->right = makeTree(preorder, inorder, i + 1, right);

    return node;
}

class Solution {
   public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        preorder_index = 0;
        return makeTree(preorder, inorder, 0, inorder.size() - 1);
    }
};

void trimLeftTrailingSpaces(string& input) {
    input.erase(input.begin(), find_if(input.begin(), input.end(), [](int ch) {
                    return !isspace(ch);
                }));
}

void trimRightTrailingSpaces(string& input) {
    input.erase(find_if(input.rbegin(), input.rend(), [](int ch) {
                    return !isspace(ch);
                }).base(),
                input.end());
}

vector<int> stringToIntegerVector(string input) {
    vector<int> output;
    trimLeftTrailingSpaces(input);
    trimRightTrailingSpaces(input);
    input = input.substr(1, input.length() - 2);
    stringstream ss;
    ss.str(input);
    string item;
    char   delim = ',';
    while (getline(ss, item, delim)) {
        output.push_back(stoi(item));
    }
    return output;
}

string treeNodeToString(TreeNode* root) {
    if (root == nullptr) {
        return "[]";
    }

    string           output = "";
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        TreeNode* node = q.front();
        q.pop();

        if (node == nullptr) {
            output += "null, ";
            continue;
        }

        output += to_string(node->val) + ", ";
        q.push(node->left);
        q.push(node->right);
    }
    return "[" + output.substr(0, output.length() - 2) + "]";
}

int main() {
    string line;
    cout << "example input preorder and inorder, ctrl+c/ctrl+d/enter to exit" << endl;
    cout << "[3,9,20,15,7]\n[9,3,15,20,7]" << endl;

    while (getline(cin, line)) {
        vector<int> preorder = stringToIntegerVector(line);
        getline(cin, line);
        vector<int> inorder = stringToIntegerVector(line);

        TreeNode*   ret     = Solution().buildTree(preorder, inorder);

        string      out     = treeNodeToString(ret);
        cout << out << endl;
    }
    return 0;
}
