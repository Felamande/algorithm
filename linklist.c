#include "linklist.h"
#include <malloc.h>
#include <stdio.h>

typedef int (INSERT_FUNC)(int);

int initList(LinkList *l) {
    l->head = (Lnode *)malloc(sizeof(Lnode));
    l->tail = (Lnode *)malloc(sizeof(Lnode));

    l->head->next = l->tail;
    l->head->pre = NULL;

    l->tail->pre = l->head;
    l->tail->next = NULL;

    l->len = 0;
    return TRUE;
}

RESULT insertHead(LinkList *l, int val) {
    Lnode *node = (Lnode *)malloc(sizeof(Lnode));
    if (!node) {
        return FAIL_MEM;
    }
    node->val = val;

    node->next = l->head->next;
    l->head->next = node;
    node->next->pre = node;
    node->pre = l->head;
    l->len++;

    return SUCCESS;
}

RESULT insertTail(LinkList *l, int val) {
    

    Lnode *node = (Lnode *)malloc(sizeof(Lnode));
    if (!node) {
        return FAIL_MEM;
    }
    node->val = val;

    node->next = l->tail;
    l->tail->pre->next = node;
    node->pre = l->tail->pre;
    l->tail->pre = node;

    l->len++;
    return SUCCESS;
}

RESULT delHead(LinkList *l, int *val){
    if (l->len == 0) {
        if (l->head->next != l->tail || l->tail->pre != l->head) {
            return FAIL_INTERNAL;
        }
        return EMPTY;
    }

    Lnode *node = l->head->next;
    *val = node->val;

    node->next->pre = node->pre;
    node->pre->next = node->next;

    free(node);
    l->len--;
    return SUCCESS;
}

RESULT delTail(LinkList *l, int *val){
     if (l->len == 0) {
        if (l->head->next != l->tail || l->tail->pre != l->head) {
            return FAIL_INTERNAL;
        }
        return EMPTY;
    }

    Lnode *node = l->tail->pre;
    *val = node->val;

    l->tail->pre = node->pre;
    l->tail->pre->next = l->tail;

    free(node);
    l->len--;
    return SUCCESS;
}

void printList(LinkList *l){
    printf("H-");
    for(Lnode *p = l->head->next;p != l->tail;p=p->next){
        printf("%d-",p->val);
    }
    printf("T\n");
    return;
}