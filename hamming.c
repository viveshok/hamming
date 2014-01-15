
#include "hamming.h"

unsigned int weight(unsigned long long input) {

    unsigned long long mask1, mask2, mask3, mask4, mask5, mask6;
    mask1 = 6148914691236517205ULL; // 01010101...
    mask2 = 3689348814741910323ULL; // 00110011...
    mask3 = 1085102592571150095ULL; // 00001111...
    mask4 = 71777214294589695ULL; // 8 zeroes, 8 ones, etc...
    mask5 = 70367670468607ULL; // 16 zeroes, 16 ones, etc...
    mask6 = 4294967295ULL; // 32 zeroes, 32 ones

    input = (input & mask1) + ((input>>1) & mask1);
    input = (input & mask2) + ((input>>2) & mask2);
    input = (input & mask3) + ((input>>4) & mask3);
    input = (input & mask4) + ((input>>8) & mask4);
    input = (input & mask5) + ((input>>16) & mask5);
    input = (input & mask6) + ((input>>32) & mask6);

    return (unsigned int) input;
}

