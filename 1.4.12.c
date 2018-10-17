#include <stdio.h>
#include <stdlib.h>
int *intersec(int *a, int *b, int len) {
    int *re = (int *)malloc(sizeof(int) * len);

    int re_idx = 0;
    int a_idx = 0;
    int b_idx = 0;
    int last_val = a[a_idx];
    for (;;) {
        if (a_idx > len - 1 || b_idx > len - 1) {
            break;
        }

        if (a[a_idx] < b[b_idx]) {
            a_idx++;
            continue;
        } else if (a[a_idx] > b[b_idx]) {
            b_idx++;
            continue;
        }

        // a[a_idx]==b[b_idx]
        int dup_val = a[a_idx];
        if (dup_val == last_val && re_idx != 0) {
            a_idx++;
            b_idx++;
            continue;
        }
        re[re_idx] = dup_val;
        last_val = dup_val;

        re_idx++;
        a_idx++;
        b_idx++;
    }

    re = realloc(re, sizeof(int) * (re_idx + 1));
    return re;
}

int main(int argc, char const *argv[]) {
    int a[8] = {1, 1, 2, 3, 4, 6, 8, 9};
    int b[8] = {2, 2, 2, 4, 5, 6, 9, 11};
    int *c = intersec(a, b, 8);
    for (int i = 0; i < 8; i++) {
        printf("%d ", c[i]);
    }
    printf("%d\n", sizeof(a));
    return 0;
}
