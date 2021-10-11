package amcs

import (
	"fmt"
)

// Sign
// mobileConfigPath the path of the file that needs to be signed
// outPath file output path after signing
// sslKeyPath ssl key path
// sslCrtPath ssl crt path
// iosCrtPath self-signed csr path
func Sign(mobileConfigPath, outPath, sslKeyPath, sslCrtPath, iosCrtPath string) error {
	err := cmd(fmt.Sprintf(`
	openssl smime -sign -in %s -out %s -signer %s -inkey %s -certfile %s -outform der -nodetach
	`, mobileConfigPath, outPath, sslCrtPath, sslKeyPath, iosCrtPath))
	if err != nil {
		return err
	}
	return nil
}
