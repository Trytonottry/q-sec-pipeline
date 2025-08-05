package resign

import (
	"context"
	"fmt"
	"os"

	"github.com/qsec-pipeline/qsec-pipeline/internal/crypto"
	"github.com/qsec-pipeline/qsec-pipeline/internal/registry"
)

func Run() error {
	image := os.Getenv("IMAGE")
	if image == "" {
		return fmt.Errorf("IMAGE env required")
	}

	// 1. Pull
	layers, err := registry.Pull(context.Background(), image)
	if err != nil { return err }

	// 2. Sign
	sig, err := crypto.SignDilithium([]byte(image))
	if err != nil { return err }

	// 3. Encrypt layer (Kyber)
	encLayers, err := crypto.EncryptKyber(layers)
	if err != nil { return err }

	// 4. Push
	return registry.Push(context.Background(), image, encLayers, sig)
}