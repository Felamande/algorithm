#include "ringbuff.h"
#include <malloc.h>
int initbuf(ringbuf *rb, int cap) {
    rb->read = 0;
    rb->write = 0;
    rb->len = 0;
    rb->cap = cap;
    rb->buf = (char *)malloc(sizeof(char) * cap);
    return TRUE;
}

int next(ringbuf *rb, int idx) { return idx == rb->cap - 1 ? 0 : idx + 1; }

RESULT read(ringbuf *rb, char *val) {
    if (rb->len == 0) {
        return EMPTY;
    };

    *val = rb->buf[rb->read];
    rb->read = next(rb, rb->read);
    rb->len--;
    return SUCCESS;
}

RESULT write(ringbuf *rb, char val) {
    if (rb->len == MAX) {
        return FULL;
    }

    rb->buf[rb->write] = val;
    rb->write = next(rb, rb->write);
    rb->len++;
    return SUCCESS;
}
