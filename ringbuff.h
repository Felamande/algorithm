#include "common.h"
// #define MAX 1024

typedef struct ringbuf_ {
    int read;
    int write;
    int len;
    int cap;
    char *buf;
} ringbuf;

int initbuf(ringbuf *rb, int cap);
int next(ringbuf *rb, int idx);
RESULT read(ringbuf *rb, char *val);
RESULT write(ringbuf *rb, char val);