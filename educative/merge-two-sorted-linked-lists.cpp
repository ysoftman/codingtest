/*
https://www.educative.io/m/merge-two-sorted-linked-lists

Given two sorted linked lists, merge them so that the resulting linked list is also sorted.

Consider two sorted linked lists as an example.

head1 -> 4 -> 8 -> 15 -> 19 -> NULL
head2 -> 7 -> 9 -> 10 -> 16 -> NULL

The merged linked list should look like this:

head1 -> 4 -> 7 -> 8 -> 9 -> 10 -> 15 -> 16 -> 19 -> NULL
*/

#include <iostream>
#include <vector>

using namespace std;

// LinkedListNode 내용이 없어 로컬에서 빌드하지 못함,
// 위 사이트내에서 테스트
typedef LinkedListNode *NodePtr;

NodePtr merge_sorted(NodePtr head1, NodePtr head2)
{
    //TODO: Write - Your - Code
    NodePtr mergehead;
    if (head1 != NULL && head2 != NULL)
    {
        if (head1->data < head2->data)
        {
            mergehead = head1;
            head1 = head1->next;
        }
        else
        {
            mergehead = head2;
            head2 = head2->next;
        }
    }
    NodePtr mergeheadstart = mergehead;

    while (head1 != NULL || head2 != NULL)
    {
        if (head1 != NULL && head2 != NULL)
        {
            if (head1->data < head2->data)
            {
                mergehead->next = head1;
                head1 = head1->next;
            }
            else
            {
                mergehead->next = head2;
                head2 = head2->next;
            }
        }
        else if (head1 != NULL && head2 == NULL)
        {
            mergehead->next = head1;
            head1 = head1->next;
        }
        else if (head1 == NULL && head2 != NULL)
        {
            mergehead->next = head2;
            head2 = head2->next;
        }
        mergehead = mergehead->next;
    }
    return mergeheadstart;
}
