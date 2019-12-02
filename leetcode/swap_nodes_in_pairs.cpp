// https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/
#include <iostream>
using namespace std;

//  Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

ListNode *swapPairs(ListNode *head)
{
    if (head == NULL)
    {
        return NULL;
    }
    if (head->next == NULL)
    {
        return head;
    }
    ListNode *temp;
    temp = head->next;
    head->next = swapPairs(head->next->next);
    temp->next = head;
    head = temp;

    return head;
}

void printlist(ListNode *node)
{
    while (node != NULL)
    {
        cout << node->val << " ";
        node = node->next;
    }
    cout << endl;
}

int main()
{
    ListNode *node = new ListNode(1);
    ListNode *head = node;
    for (int i = 2; i <= 10; ++i)
    {
        ListNode *nn = new ListNode(i);
        node->next = nn;
        node = node->next;
    }
    printlist(head);
    head = swapPairs(head);
    printlist(head);
    return 0;
}