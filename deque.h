#include "common.h"

typedef struct deque_ {
    int first;
    int last;
    int len;
    int cap;
    int initcap;
    int *elems;
} deque;

RESULT initDeque(deque *q, int cap);
int isfull(deque *q);
int resize(deque *q, int newcap);
int pushRight(deque *q, int val);
int pushLeft(deque *q, int val);
RESULT popLeft(deque *q, int *val);