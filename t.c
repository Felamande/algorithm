#include "time.h"
#include <stdio.h>

int main() {

    LONGLONG tt;
    LONGLONG t1 = readTime();
    for (int i = 0; i < 1100000000; i++) {
        // tt = readTime();
    }
    LONGLONG t2 = readTime();
    printf("%.7f\n", (t2 - t1) * 1.0 / 1e7);
}