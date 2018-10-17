#include "deque.h"
#include <stdio.h>
#include <stdlib.h>

int main(void) {
    deque *q = (deque *)malloc(sizeof(deque));
    initDeque(q, 128);
    for (int i = 0; i < 100; i++) {
        pushRight(q, i);
        pushLeft(q, -i);
        printf("first=%d, last=%d, len=%d, cap=%d\n", q->first, q->last, q->len,
               q->cap);
    }
    for (int i = 0; i < 200; i++) {
        int e;
        RESULT re = popLeft(q, &e);
        printf("first=%d, last=%d, len=%d, cap=%d, pop=%d\n", q->first, q->last,
               q->len, q->cap, e);
    }
    int e;

    RESULT re = popLeft(q, &e);
    printf("first=%d, last=%d, len=%d, cap=%d, pop=%d\n", q->first, q->last,
           q->len, q->cap, e);

    return 0;
}