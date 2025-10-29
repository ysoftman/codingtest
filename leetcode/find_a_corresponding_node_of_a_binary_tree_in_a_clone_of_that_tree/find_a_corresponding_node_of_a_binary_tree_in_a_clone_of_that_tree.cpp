/*
https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/
1379. Find a Corresponding Node of a Binary Tree in a Clone of That Tree
Easy
Given two binary trees original and cloned and given a reference to a node target in the original tree.
The cloned tree is a copy of the original tree.
Return a reference to the same node in the cloned tree.
Note that you are not allowed to change any of the two trees or the target node and the answer must be a reference to a node in the cloned tree.

Example 1:
Input: tree = [7,4,3,null,null,6,19], target = 3
Output: 3
Explanation: In all examples the original and cloned trees are shown. The target node is a green node from the original tree. The answer is the yellow node from the cloned tree.

Example 2:
Input: tree = [7], target =  7
Output: 7

Example 3:
Input: tree = [8,null,6,null,5,null,4,null,3,null,2,null,1], target = 4
Output: 4


Constraints:
The number of nodes in the tree is in the range [1, 104].
The values of the nodes of the tree are unique.
target node is a node from the original tree and is not null.

Follow up: Could you solve the problem if repeated values on the tree are allowed?
*/
#include <iostream>

// Definition for a binary tree node.
struct TreeNode {
    int       val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
   public:
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        // DFS - prerder
        if (original == nullptr) {
            std::cout << "-nullptr-" << std::endl;
            return nullptr;
        }
        if (original->val == target->val) {
            std::cout << cloned->val << std::endl;
            return cloned;
        }
        TreeNode* left = getTargetCopy(original->left, cloned->left, target);
        if (left != nullptr) {
            return left;
        }
        TreeNode* right = getTargetCopy(original->right, cloned->right, target);
        if (right != nullptr) {
            return right;
        }
        std::cout << "---last nullptr---" << std::endl;
        return nullptr;
    }
};

int main() {
    // [7,4,3,null,null,6,19], target = 3
    TreeNode* a = new TreeNode(3);
    a->left     = new TreeNode(6);
    a->right    = new TreeNode(19);
    TreeNode org(7);
    org.left    = a;
    org.right   = new TreeNode(3);

    TreeNode* b = new TreeNode(3);
    b->left     = new TreeNode(6);
    b->right    = new TreeNode(19);
    TreeNode cloned(7);
    cloned.left  = b;
    cloned.right = new TreeNode(3);

    TreeNode  target(3);
    Solution  sol;
    TreeNode* ret = sol.getTargetCopy(&org, &cloned, &target);
    if (ret != nullptr) {
        std::cout << ret->val << std::endl;
    }
}
