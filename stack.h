#include "common.h"

#ifdef STACK_LL
typedef struct node {
    struct node *next;
    char val;
} node;

typedef struct _stack {
    struct node *root;
    int len;

} stack;
#endif

#ifdef STACK_ARR

typedef struct _stack {
    int len;
    int cap;
    char *elems;
} stack;

int isfull(stack *s);
RESULT resize(stack *s, int newcap);

#endif // STACK_ARR

RESULT initStack(stack *s);

RESULT push(stack *s, char val);

RESULT pop(stack *s, char *val);

char peek(stack *s);

int printstack(stack *s);
int swap(stack *dst, stack *src);