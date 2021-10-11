package amcs

import "testing"


func TestSign(t *testing.T) {
	mobileconfigPath := "resource/example.mobileconfig"
	iosCrtPath := "resource/ios.crt"
	sslCrtPath := "resource/ssl.crt"
	sslKeyPath := "resource/ssl.key"
	outPath := "resource/out.mobileconfig"
	err := Sign(mobileconfigPath,outPath,sslKeyPath,sslCrtPath,iosCrtPath)
	if err != nil{
		panic(err)
	}
}