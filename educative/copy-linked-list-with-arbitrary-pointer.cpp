/*
https://www.educative.io/m/copy-linked-list-with-arbitrary-pointer

You are given a linked list where the node has two pointers. The first is the regular next pointer. The second pointer is called arbitrary_pointer and it can point to any node in the linked list. Your job is to write code to make a deep copy of the given linked list. Here, deep copy means that any operations on the original list (inserting, modifying and removing) should not affect the copied list.

Here’s an example of a linked list with arbitrary pointers connected.

7 --> 14 --> 21 --> NULL

arbitrary_pointer
7->21
14->NULL
21->7
*/

// LinkedListNode 내용이 없어 로컬에서 빌드하지 못함,
// 위 사이트내에서 테스트
typedef LinkedListNode *NodePtr;

#include <unordered_map>
using namespace std;

LinkedListNode *deep_copy_arbitrary_pointer(
    LinkedListNode *head) {
    // TODO: Write - Your - Code
    unordered_map<LinkedListNode *, LinkedListNode *> hashmap;
    LinkedListNode                                   *copy_list = new LinkedListNode(head->data);
    copy_list->arbitrary_pointer                                = head->arbitrary_pointer;
    LinkedListNode *copy_list_head                              = copy_list;
    head                                                        = head->next;
    while (head != NULL) {
        LinkedListNode *node    = new LinkedListNode(head->data);
        node->arbitrary_pointer = head->arbitrary_pointer;
        copy_list->next         = node;
        // 현재 원본 노드 주소를 복사된 신규 노드 주소와 맵핑
        hashmap[head] = node;
        head          = head->next;
        copy_list     = copy_list->next;
    }
    copy_list = copy_list_head;
    while (copy_list != NULL) {
        // arbitrary_poiter 를 신규 노드 주소에 맞게 변경
        if (hashmap.find(copy_list->arbitrary_pointer) != hashmap.end()) {
            copy_list->arbitrary_pointer = hashmap[copy_list->arbitrary_pointer];
        }
        copy_list = copy_list->next;
    }
    return copy_list_head;
}
