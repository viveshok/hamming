
-module(benchmark).

-export([benchmark/0]).

benchmark() ->
    NumInts = 2000000,
    Payload = list_to_binary([random:uniform(100) || _ <- lists:seq(1, NumInts)]),

    T0_Erlang = timestamp(),
    PopCountErlang = hamming:weight(Payload),
    T1_Erlang = timestamp(),

    T0_C = timestamp(),
    PopCountC = hamming:cweight(Payload),
    T1_C = timestamp(),

    io:format("~nTest size: ~p words (64 bits per words)~n", [NumInts]),
    io:format("Erlang and C produce same Hamming weigth? ~p ~n", [PopCountErlang=:=PopCountC]),
    io:format("Erlang took ~p microseconds ~n", [T1_Erlang-T0_Erlang]),
    io:format("C (NIF) took ~p microseconds ~n~n", [T1_C-T0_C]),
    ok.

timestamp() ->
    {Mega, Sec, Micro} = now(),
    Mega * 1000000 * 1000000 + Sec * 1000000 + Micro.

