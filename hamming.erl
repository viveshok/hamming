
-module(hamming).

-export([weight/1, cweight/1]).

-on_load(init/0).

init() ->
    ok = erlang:load_nif("./hamming_nif", 0).

weight(Bitstring) -> weight(Bitstring, 0).

weight(<<>>, Acc) -> Acc;
weight(<<0:1, Tl/bits>>, Acc) -> weight(Tl, Acc);
weight(<<1:1, Tl/bits>>, Acc) -> weight(Tl, Acc+1).

cweight(Binary) -> cweight(Binary, 0).

cweight(<<>>, Acc) -> Acc;
cweight(<<Char:64, Tl/binary>>, Acc) -> weight(Tl, Acc + char_weight(Char)).

char_weight(_) ->
    exit(nif_library_not_loaded).

