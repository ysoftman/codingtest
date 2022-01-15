/*
https://www.educative.io/m/is-binary-tree-a-binary-search-tree

Is Binary Tree a Binary Search Tree?
Problem Statement
Given a Binary Tree, figure out whether it’s a Binary Search Tree. In a binary search tree, each node’s key value is smaller than the key value of all nodes in the right subtree, and are greater than the key values of all nodes in the left subtree i.e. L < N < RL<N<R.

Below is an example of binary tree that is a valid BST.

      100
    /     \
  50       200
 /  \     /   \
25   75  125  350

Below is an example of a binary tree that is not a BST

      100
    /     \
  50       200
 /  \     /   \
25   75  90   350

*/

bool is_bst_recursive(BinaryTreeNode *root, int min, int max)
{
    if (root == nullptr)
    {
        return true;
    }
    if (root->data < min || root->data > max)
    {
        return false;
    }

    // inorder traversal (left -> root -> right)
    // left
    if (is_bst_recursive(root->left, min, root->data) == false)
    {
        return false;
    }
    // root
    // cout << root->data;
    // right
    if (is_bst_recursive(root->right, root->data, max) == false)
    {
        return false;
    }
    return true;
}

bool is_bst(BinaryTreeNode *root)
{
    //TODO: Write - Your - Code
    int max = (1 << 31 - 1);
    int min = (1 << 31 - 1) * -1;
    return is_bst_recursive(root, min, max);
}
