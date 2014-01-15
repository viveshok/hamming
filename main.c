
// gcc -Wall -W -g hamming.c main.c -lrt

#include<stdio.h>
#include <stdlib.h>
#include <time.h>

#include "hamming.h"

int main() {

    int i, numInts = 2000000;
    unsigned long long * randoms = malloc(numInts*sizeof(unsigned long long));
    for(i = 0; i < numInts; ++i)
        randoms[i] = (unsigned long long) rand() % 100 + 1;

    struct timespec startTime;
    struct timespec endTime;

    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &startTime);
    for(i = 0; i < numInts; ++i)
        weight(randoms[i]);
    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &endTime);

    free(randoms);
    long double delta = 1000000*(endTime.tv_sec - startTime.tv_sec) + (long double) (endTime.tv_nsec - startTime.tv_nsec)/1000.0;
    printf ("C (RAW) took %lu microseconds\n", (unsigned long int) delta);
    return 0;
}

