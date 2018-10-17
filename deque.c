#include "deque.h"
#include <malloc.h>
#include <stdio.h>

RESULT initDeque(deque *q, int cap) {
    q->first = cap / 2 - 1;
    q->last = cap / 2;
    q->elems = (int *)malloc(sizeof(int) * cap);
    if (!q->elems) {
        return FAIL_MEM;
    }
    q->cap = cap;
    q->initcap = cap;
    q->len = 0;
    return SUCCESS;
}

// int next(deque *q, int pos) {
//     if (pos == q->last) {
//         return q->first;
//     }
//     return pos + 1;
// }

//
int isfull(deque *q) {
    if (q->len == q->cap) {
        return TRUE;
    }
    return FALSE;
}

int resize(deque *q, int newcap) {
    int *oldelems = q->elems;
    q->elems = (int *)malloc(sizeof(int) * newcap);
    if (!(q->elems)) {
        q->elems = oldelems;
        return FAIL_MEM;
    }

    int oldfirst = q->first;
    int newfirst = newcap / 2 - 1 - q->len / 2;
    int newlast = newfirst + q->len + 1;

    for (int idx = 0; idx < q->len; idx++) {
        q->elems[newfirst + 1 + idx] = oldelems[oldfirst + 1 + idx];
    }

    q->first = newfirst;
    q->last = newlast;
    q->cap = newcap;

    free(oldelems);
    return SUCCESS;
}

int pushRight(deque *q, int val) {
    if (q->last == q->cap) {
        resize(q, 2 * q->cap);
    }
    q->elems[q->last] = val;
    q->len++;
    q->last++;
}

int pushLeft(deque *q, int val) {
    if (q->first == -1) {
        resize(q, 2 * q->cap);
    }
    q->elems[q->first] = val;
    q->first--;
    q->len++;
}

RESULT popLeft(deque *q, int *val) {
    if (q->len == 0) {
        return EMPTY;
    }
    q->first++;
    q->len--;
    *val = q->elems[q->first];
    if (q->len > q->initcap && q->len == q->cap / 4) {
        resize(q, q->cap / 2);
    }
    return SUCCESS;
}