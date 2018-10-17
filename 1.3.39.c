#include "ringbuff.h"
#include <malloc.h>
#include <pthread.h>
int main(void) {
    ringbuf *rb = (ringbuf *)malloc(sizeof(ringbuf));
    initbuf(rb, 1024);

    return 0;
}