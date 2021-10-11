# amcs(apple mobile config signature)
sign Apple’s mobileconfig file to solve the ‘unsigned’ problem

# the project rely openssl
> https://github.com/openssl/openssl

# first step

* to prepare a self-signed certificate

```
openssl genrsa -out ios.key 2048
openssl req -new -sha256 -key ios.key -out ios.csr
```

# SSL certificate

* You need to hold the crt and key files issued by the operator

# how to use?

```
go get github.com/19byte/amcs

err := amcs.Sign(mobileconfigPath,outPath,sslKeyPath,sslCrtPath,iosCrtPath)
if err != nil {
    panic(err)
}
```

## before signing
![before signing](https://s3.bmp.ovh/imgs/2021/10/b337ce84d8081225.png)

## after signing

![after signing](https://s3.bmp.ovh/imgs/2021/10/922cab47d0be5a6a.png)
