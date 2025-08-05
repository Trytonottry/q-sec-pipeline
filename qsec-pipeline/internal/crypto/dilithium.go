package crypto

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -loqs
#include <oqs/oqs.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

func SignDilithium(msg []byte) ([]byte, error) {
	sig := C.OQS_SIG_new(C.CString("Dilithium3"))
	if sig == nil {
		return nil, errors.New("cannot init Dilithium3")
	}
	defer C.OQS_SIG_free(sig)

	pk := make([]byte, sig.length_public_key)
	sk := make([]byte, sig.length_secret_key)
	signature := make([]byte, sig.length_signature)

	if C.OQS_SIG_keypair(sig,
		(*C.uint8_t)(unsafe.Pointer(&pk[0])),
		(*C.uint8_t)(unsafe.Pointer(&sk[0]))) != C.OQS_SUCCESS {
		return nil, errors.New("keygen failed")
	}

	var sigLen C.size_t
	if C.OQS_SIG_sign(sig,
		(*C.uint8_t)(unsafe.Pointer(&signature[0])),
		&sigLen,
		(*C.uint8_t)(unsafe.Pointer(&msg[0])),
		C.size_t(len(msg)),
		(*C.uint8_t)(unsafe.Pointer(&sk[0]))) != C.OQS_SUCCESS {
		return nil, errors.New("sign failed")
	}
	return signature[:sigLen], nil
}