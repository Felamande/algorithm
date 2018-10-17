#include "common.h"

typedef struct _Lnode {
    struct _Lnode *next;
    struct _Lnode *pre;
    int val;
} Lnode;

typedef struct _LinkList {
    Lnode *head;
    Lnode *tail;
    int len;
} LinkList;

int initList(LinkList *l) ;
RESULT insertHead(LinkList *l, int val);
RESULT insertTail(LinkList *l, int val);
int delHead(LinkList *l,int *val);
int delTail(LinkList *l,int *val);
void printList(LinkList *l);