package types

import (
	"crypto"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"hash"
	"io"
)

type TrustChain struct {
	CommonName                    string
	CertificateAuthorityDirectory []string
	CertificateAuthorityFiles     []string
}

type Path struct {
	File   string
	Buffer int
}

type Reader struct {
	Interface io.Reader
	Buffer    int
}

type Data struct {
	Raw    *[]byte
	Path   Path
	Reader Reader
}

type Manifest struct {
	CertificateChain []*x509.Certificate
	SigningAlgorithm x509.SignatureAlgorithm
	Signature        *[]byte
	Hash             *[]byte
	Data             Data
}

var SignatureAlgorithm = map[x509.SignatureAlgorithm]func() (hash.Hash, crypto.Hash){
	x509.ECDSAWithSHA256: func() (hash.Hash, crypto.Hash) {
		return sha256.New(), crypto.SHA256
	},
	x509.ECDSAWithSHA384: func() (hash.Hash, crypto.Hash) {
		return sha512.New384(), crypto.SHA384
	},
	x509.ECDSAWithSHA512: func() (hash.Hash, crypto.Hash) {
		return sha512.New(), crypto.SHA512
	},
	x509.SHA256WithRSA: func() (hash.Hash, crypto.Hash) {
		return sha256.New(), crypto.SHA256
	},
	x509.SHA384WithRSA: func() (hash.Hash, crypto.Hash) {
		return sha512.New384(), crypto.SHA384
	},
	x509.SHA512WithRSA: func() (hash.Hash, crypto.Hash) {
		return sha512.New(), crypto.SHA512
	},
	x509.SHA256WithRSAPSS: func() (hash.Hash, crypto.Hash) {
		return sha256.New(), crypto.SHA256
	},
	x509.SHA384WithRSAPSS: func() (hash.Hash, crypto.Hash) {
		return sha512.New384(), crypto.SHA384
	},
	x509.SHA512WithRSAPSS: func() (hash.Hash, crypto.Hash) {
		return sha512.New(), crypto.SHA512
	},
}
