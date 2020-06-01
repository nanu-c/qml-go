#include "textflag.h"

TEXT ·Addrs(SB),NOSPLIT,$0-16
	MOVD	$runtime·main(SB), R0
	MOVD	R0, ret+0(FP)
	MOVD	$runtime·main_main(SB), R0
	MOVD	R0, ret+8(FP)
	RET
