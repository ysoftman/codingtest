/*
https://www.educative.io/m/level-order-traversal-binary-tree

Level Order Traversal of Binary Tree
Problem Statement
Given the root of a binary tree, display the node values at each level. Node values for all levels should be displayed on separate lines. Let’s take a look at the below binary tree.


   100
   /  \
  50   200
 /  \     \
25   75    350

Level order traversal for this tree should look like:
100
50, 200
25, 75, 350
*/

#include <queue>
#include <vector>
using namespace std;
string level_order_traversal(BinaryTreeNode *root)
{
    string result = "";
    //TODO: Write - Your - Code
    queue<BinaryTreeNode *> q;
    q.push(root);
    while (not q.empty())
    {
        vector<BinaryTreeNode *> vec;
        while (not q.empty())
        {
            BinaryTreeNode *node = q.front();
            q.pop();
            vec.push_back(node);
            result += to_string(node->data) + ",";
        }
        // 라인 구분이 필요할때
        // if (result.length() > 0 and result[result.length() - 1] == ',')
        // {
        //     result[result.length() - 1] = '\n';
        // }
        for (auto v : vec)
        {
            if (v->left != nullptr)
            {
                q.push(v->left);
            }
            if (v->right != nullptr)
            {
                q.push(v->right);
            }
        }
        vec.clear();
    }
    return result;
}
