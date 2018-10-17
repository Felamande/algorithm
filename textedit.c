#include "stack.h"
#include "time.h"
#include <stdio.h>
#include <stdlib.h>

typedef struct textmodel_ {
    stack *beforecur;
    stack *aftercur;
} textmodel;

RESULT initmodel(textmodel *tm) {
    tm->beforecur = (stack *)malloc(sizeof(stack));
    if (!tm->beforecur) {
        return FAIL_MEM;
    }
    initStack(tm->beforecur);

    tm->aftercur = (stack *)malloc(sizeof(stack));
    if (!tm->aftercur) {
        return FAIL_MEM;
    }
    initStack(tm->aftercur);
    return SUCCESS;
}

RESULT insert(textmodel *tm, char c) {
    RESULT re = push(tm->beforecur, c);
    return re;
}

RESULT delete (textmodel *tm, char *c) { return pop(tm->beforecur, c); }
int cursor(textmodel *tm) { return tm->beforecur->len; }
RESULT right(textmodel *tm, int k) {
    srand(2);
    char c;
    RESULT re;
    for (int i = 0; i < k; i++) {
        re = pop(tm->aftercur, &c);
        if (re != SUCCESS) {
            break;
        }
        push(tm->beforecur, c);
    }
}
int readfile(textmodel *tm, FILE *fp) {

    char c = 0;
    RESULT re;
    int buflen = 100 * 1 << 20;

    char *buf = (char *)malloc(buflen);
    int readlen = -1;
    long long t1 = readTime();
    while (readlen != 0) {
        readlen = fread(buf, sizeof(char), buflen, fp);
        for (int i = 0; i < readlen; i++) {
            push(tm->beforecur, buf[i]);
        }
    }
    long long t2 = readTime();
    free(buf);
    // int a = strlen(buf);
    int filelen = tm->beforecur->len;
    for (int i = 0; i < filelen; i++) {
        re = pop(tm->beforecur, &c);
        re = push(tm->aftercur, c);
    }
    long long t3 = readTime();
    printf("read=%.7f, swap=%.7f\n", (t2 - t1) * 1.0 / 1e7,
           (t3 - t2) * 1.0 / 1e7);
}

int main(int argc, char const *argv[]) {
    if (argc <= 1) {
        printf("input file not found:\n");
    }

    FILE *fp = fopen(argv[1], "r");
    if (!fp) {
        printf("input file not found:%s\n", argv[1]);
        return -1;
    }

    textmodel *tm = (textmodel *)malloc(sizeof(textmodel));
    initmodel(tm);
    readfile(tm, fp);
    right(tm, 7);
    // getchar();
    return 0;
}
