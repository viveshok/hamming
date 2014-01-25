### hamming

sandbox to play with and benchmark bit operations and interop in C, Erlang and Go

#### how to run (on Linux)

For Golang:

    $ go run hamming.go

For C:

    $ gcc -Wall -W -g hamming.c main.c -lrt -O3 -o hamming.out
    $ ./hamming.out

For Erlang:

    $ gcc -o hamming_nif.so -fpic -shared hamming.c hamming_nif.c -I/usr/lib/erlang/erts-5.10.4/include
    $ erlc hamming.erl benchmark.erl
    $ erl
    > benchmark:benchmark().

