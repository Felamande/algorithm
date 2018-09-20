#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>
#include "time.h"

#define STACKSIZE 1024
#define TRUE 1
#define FALSE 0

typedef struct _stack {
    int len;
    int cap;
    int *elems;
} stack;

int initStack(stack *s, int cap) {
    s->cap = cap;
    s->len = 0;
    s->elems = (int *)malloc(sizeof(int) * cap);
    // s->elems = (int *)malloc(STACKSIZE * sizeof(int));
}
int isfull(stack *s) { return s->cap == s->len; }

int resize(stack *s, int newcap) {
    int *backup = s->elems;
    s->elems = (int *)realloc(s->elems, sizeof(int) * newcap);
    if (!(s->elems)) {
        s->elems = backup;
        return FALSE;
    }

    s->cap = newcap;
    return TRUE;
}

void push(stack *s, int val) {
    if (isfull(s)) {
        resize(s, 2 * s->cap);
    }
    s->elems[s->len] = val;
    s->len++;
}

int pop(stack *s) {
    s->len--;
    int elem = s->elems[s->len];
    if (s->len > 0 && s->len == s->cap / 4) {
        resize(s, s->cap / 2);
    }
    return elem;
}

int peek(stack *s){
    return s->elems[s->len];
}

int printstack(stack *s) {
    for (int i = 0; i < s->len; i++) {
        printf("%d\n", s->elems[i]);
    }
}

int main() {
    stack *s = (stack *)malloc(sizeof(stack));
    initStack(s, STACKSIZE);

    LONGLONG t1 = readTime();
    for (int i = 0; i < 110000000; i++) {
        push(s, i);
        // printf("len=%d, cap=%d, addr=0x%x\n", s->len, s->cap, s->elems);
    }
    LONGLONG t2 = readTime();

    printf("push=%.7fs\n", (t2 - t1) * 1.0 / 1e7);
    return 0;
};