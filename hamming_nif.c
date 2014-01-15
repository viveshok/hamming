
/*
gcc -I /home/alexandre/erlang/otp_src_R15B02/erts-5.9.2/emulator/beam/ -I /home/alexandre/erlang/otp_src_R15B02/erts-5.9.2/include/x86_64-unknown-linux-gnu/ -o hamming_nif.so -fpic -shared hamming.c hamming_nif.c
*/

#include "erl_nif.h"

extern unsigned int weight(unsigned long long);

static ERL_NIF_TERM weight_nif(ErlNifEnv* env, int argc, const ERL_NIF_TERM argv[])
{
    unsigned long long input;
    unsigned int ret;

    if (!enif_get_uint64(env, argv[0], &input)) {
	return enif_make_badarg(env);
    }
    ret = weight(input);
    return enif_make_uint(env, ret);
}

static ErlNifFunc nif_funcs[] = {
    {"char_weight", 1, weight_nif}
};

ERL_NIF_INIT(hamming, nif_funcs, NULL, NULL, NULL, NULL)

