#include "time.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>

typedef struct node {
    struct node *next;
    int val;
} node;

typedef struct stack {
    struct node *root;
    int len;

} stack;

int InitStack(stack *s) {
    s->root = (node *)malloc(sizeof(node));
    s->root->next = NULL;
    s->len = 0;
    return 1;
}

int push(stack *s, int val) {
    struct node *n = (node *)malloc(sizeof(node));
    n->next = NULL;
    n->val = val;

    n->next = s->root->next;
    s->root->next = n;
    s->len++;

    return 1;
}

int pop(stack *s) {
    if (s->len == 0) {
        return 0;
    }
    node *tmp = s->root->next;
    int data = tmp->val;
    s->root->next = tmp->next;
    s->len--;
    free(tmp);
    return data;
}

int peek(stack *s){
    return s->root->next->val;
}

int printStack(stack *s) {
    node *p = s->root->next;
    while (p != NULL) {
        printf("|%d|\n", p->val);
        p = p->next;
    }
    printf(" ￣\n");
    return 0;
}

int parenthese(char *c) {
    stack *s = (stack *)malloc(sizeof(stack));
    InitStack(s);
}

int main(void) {

    stack *s = (stack *)malloc(sizeof(stack));
    InitStack(s);
    LONGLONG t1 = readTime();
    for (int i = 0; i < 20000000; i++) {
        push(s, i);
    }
    LONGLONG t2 = readTime();
    printf("push1=%.7fs\n", (t2 - t1) * 1.0 / 1e7);
    return 0;
}
