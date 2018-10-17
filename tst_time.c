#include "time.h"
#include <stdio.h>
#include <stdlib.h>
int main() {
    long long t1 = readTime();
    for (long i = 0; i < 207455; i++) {
    }
    long long t2 = readTime();
    printf("%.7fms\n", (t2 - t1) * 1.0 / 1e4);
    return 0;
}