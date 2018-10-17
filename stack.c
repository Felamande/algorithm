#include "stack.h"
#include <stdio.h>
#include <stdlib.h>

#ifdef STACK_LL
RESULT initStack(stack *s) {
    s->root = (node *)malloc(sizeof(node));
    if (!s->root) {
        return FAIL_MEM;
    }
    s->root->next = NULL;
    s->len = 0;
    return SUCCESS;
}

RESULT push(stack *s, char val) {
    struct node *n = (node *)malloc(sizeof(node));
    if (!n) {
        return FAIL_MEM;
    }
    n->next = NULL;
    n->val = val;

    n->next = s->root->next;
    s->root->next = n;
    s->len++;

    return SUCCESS;
}

RESULT pop(stack *s, char *val) {
    if (s->len == 0) {
        return EMPTY;
    }
    node *tmp = s->root->next;
    *val = tmp->val;
    s->root->next = tmp->next;
    s->len--;
    free(tmp);
    return SUCCESS;
}

char peek(stack *s) { return s->root->next->val; }

int printstack(stack *s) {
    node *p = s->root->next;
    while (p != NULL) {
        printf("|%d|\n", p->val);
        p = p->next;
    }
    printf("--\n");
    return 0;
}
#endif // STACK_LL

#ifdef STACK_ARR
#define STACK_ARR_SIZE 1024
int initStack(stack *s) {
    s->cap = STACK_ARR_SIZE;
    s->len = 0;
    s->elems = (char *)malloc(sizeof(char) * STACK_ARR_SIZE);
}
int isfull(stack *s) { return s->cap == s->len; }

RESULT resize(stack *s, int newcap) {
    char *backup = s->elems;
    s->elems = (char *)realloc(s->elems, sizeof(char) * newcap);
    if (!(s->elems)) {
        s->elems = backup;
        return FALSE;
    }

    s->cap = newcap;
    return TRUE;
}

RESULT push(stack *s, char val) {
    if (isfull(s)) {
        RESULT re = resize(s, 2 * s->cap);
        if (re != SUCCESS) {
            return re;
        }
    }
    s->elems[s->len] = val;
    s->len++;
    return SUCCESS;
}

RESULT pop(stack *s, char *val) {
    if (s->len > 0 && s->len == s->cap / 4) {
        RESULT re = resize(s, s->cap / 2);
        if (re != SUCCESS) {
            return re;
        }
    }
    s->len--;
    *val = s->elems[s->len];
    return SUCCESS;
}

char peek(stack *s) { return s->elems[s->len]; }

int printstack(stack *s) {
    for (int i = 0; i < s->len; i++) {
        printf("%d\n", s->elems[i]);
    }
}

int swap(stack *dst, stack *src) {
    char *tmp;
    tmp = dst->elems;
    dst->elems = src->elems;
    src->elems = tmp;

    int lentmp;
    lentmp = dst->len;
    dst->len = src->len;
    src->len = lentmp;

    int captmp;
    captmp = dst->cap;
    dst->cap = src->cap;
    src->cap = captmp;
}

#endif // STACK_ARR
