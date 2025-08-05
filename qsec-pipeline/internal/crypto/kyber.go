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

func EncryptKyber(plain []byte) ([]byte, error) {
	kem := C.OQS_KEM_new(C.CString("Kyber768"))
	if kem == nil {
		return nil, errors.New("cannot init Kyber")
	}
	defer C.OQS_KEM_free(kem)

	pk := make([]byte, kem.length_public_key)
	sk := make([]byte, kem.length_secret_key)
	ciphertext := make([]byte, kem.length_ciphertext)
	sharedSecret := make([]byte, kem.length_shared_secret)

	if C.OQS_KEM_keypair(kem,
		(*C.uint8_t)(unsafe.Pointer(&pk[0])),
		(*C.uint8_t)(unsafe.Pointer(&sk[0]))) != C.OQS_SUCCESS {
		return nil, errors.New("kem keypair")
	}
	if C.OQS_KEM_encaps(kem,
		(*C.uint8_t)(unsafe.Pointer(&ciphertext[0])),
		(*C.uint8_t)(unsafe.Pointer(&sharedSecret[0])),
		(*C.uint8_t)(unsafe.Pointer(&pk[0]))) != C.OQS_SUCCESS {
		return nil, errors.New("encaps")
	}
	// На проде реальный шифр потока через sharedSecret
	return ciphertext, nil
}