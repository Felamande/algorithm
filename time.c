#include "time.h"


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
