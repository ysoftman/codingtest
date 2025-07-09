/*
https://leetcode.com/problems/swap-nodes-in-pairs/
24. Swap Nodes in Pairs
Medium
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:
Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:
Input: head = []
Output: []

Example 3:
Input: head = [1]
Output: [1]

Constraints:
The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
*/
/*
https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

// g++ -std=c++11 ./swap_nodes_in_pairs.cpp && ./a.out
#include <iostream>
using namespace std;

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

struct ListNode {
    int       val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
   public:
    // 재귀함수로 처리
    ListNode *swapPairs(ListNode *head) {
        if (head == NULL) {
            return NULL;
        }
        if (head->next == NULL) {
            return head;
        }
        ListNode *temp;
        temp       = head->next;
        head->next = swapPairs(head->next->next);
        temp->next = head;
        head       = temp;

        return head;
    }
    // 반복문으로 처리
    ListNode *swapPairs2(ListNode *head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        ListNode *temp    = NULL;
        ListNode *pre     = NULL;
        ListNode *newhead = NULL;
        ListNode *cur     = head;
        while (cur != NULL) {
            if (cur->next == NULL) {
                break;
            }
            temp       = cur->next;
            cur->next  = temp->next;
            temp->next = cur;
            if (newhead == NULL) {
                newhead = temp;
            } else {
                pre->next = temp;
            }
            pre = cur;
            cur = cur->next;
        }
        return newhead;
    }
};

void printlist(ListNode *node) {
    while (node != NULL) {
        cout << node->val << " ";
        node = node->next;
    }
    cout << endl;
}

int main() {
    Solution  sol  = Solution();
    ListNode *node = new ListNode(1);
    ListNode *head = node;
    for (int i = 2; i <= 10; ++i) {
        ListNode *nn = new ListNode(i);
        node->next   = nn;
        node         = node->next;
    }
    printlist(head);
    head = sol.swapPairs(head);
    printlist(head);

    head = new ListNode(1);
    printlist(head);
    head = sol.swapPairs(head);
    printlist(head);
    return 0;
}
