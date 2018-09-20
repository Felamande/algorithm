#include <stdio.h>
#include <windows.h>

typedef struct _KSYSTEM_TIME {
    ULONG LowPart;
    LONG High1Time;
    LONG High2Time;
} KSYSTEM_TIME;

#define KUSER_SHARED_DATA 0x7ffe0000
#define INTERRUPT_TIME ((KSYSTEM_TIME volatile *)(0x7ffe0008))

LONGLONG readTime() {
    LONG timeHigh;
    ULONG timeLow;
    do {
        timeHigh = INTERRUPT_TIME->High1Time;
        timeLow = INTERRUPT_TIME->LowPart;

    } while (timeHigh != INTERRUPT_TIME->High2Time);

    LONGLONG now = ((LONGLONG)timeHigh << 32) + timeLow;
    return now;
}


