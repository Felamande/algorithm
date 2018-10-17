#include "stack.h"
#include "time.h"
#include <malloc.h>
#include <stdio.h>

int calc(char *expr) {
    // t1_3_9("[][][]");
    stack *sym = (stack *)malloc(sizeof(stack));
    initStack(sym);

    for (char *p = expr; *p; p++) {
        char sym = *p;
        if (sym == '(') {
            continue;
        } else if ('0' <= sym && sym <= '9') {

        } else if (sym == ')') {
        }
    }

    return 0;
}

int main(void) {
    readTime();
    calc("(2-3)");
}