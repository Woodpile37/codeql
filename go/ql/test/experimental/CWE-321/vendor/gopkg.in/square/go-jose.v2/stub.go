// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for gopkg.in/square/go-jose.v2, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: gopkg.in/square/go-jose.v2 (exports: Recipient; functions: NewEncrypter,NewSigner)

// Package go_pkg is a stub of gopkg.in/square/go-jose.v2, generated by depstubber.
package go_pkg

import (
	crypto "crypto"
	x509 "crypto/x509"
	url "net/url"
)

type CompressionAlgorithm string

type ContentEncryption string

type ContentType string

type Encrypter interface {
	Encrypt(_ []byte) (*JSONWebEncryption, error)
	EncryptWithAuthData(_ []byte, _ []byte) (*JSONWebEncryption, error)
	Options() EncrypterOptions
}

type EncrypterOptions struct {
	Compression  CompressionAlgorithm
	ExtraHeaders map[HeaderKey]interface{}
}

func (_ *EncrypterOptions) WithContentType(_ ContentType) *EncrypterOptions {
	return nil
}

func (_ *EncrypterOptions) WithHeader(_ HeaderKey, _ interface{}) *EncrypterOptions {
	return nil
}

func (_ *EncrypterOptions) WithType(_ ContentType) *EncrypterOptions {
	return nil
}

type Header struct {
	KeyID        string
	JSONWebKey   *JSONWebKey
	Algorithm    string
	Nonce        string
	ExtraHeaders map[HeaderKey]interface{}
}

func (_ Header) Certificates(_ x509.VerifyOptions) ([][]*x509.Certificate, error) {
	return nil, nil
}

type HeaderKey string

type JSONWebEncryption struct {
	Header Header
}

func (_ JSONWebEncryption) CompactSerialize() (string, error) {
	return "", nil
}

func (_ JSONWebEncryption) Decrypt(_ interface{}) ([]byte, error) {
	return nil, nil
}

func (_ JSONWebEncryption) DecryptMulti(_ interface{}) (int, Header, []byte, error) {
	return 0, Header{}, nil, nil
}

func (_ JSONWebEncryption) FullSerialize() string {
	return ""
}

func (_ JSONWebEncryption) GetAuthData() []byte {
	return nil
}

type JSONWebKey struct {
	Key                         interface{}
	KeyID                       string
	Algorithm                   string
	Use                         string
	Certificates                []*x509.Certificate
	CertificatesURL             *url.URL
	CertificateThumbprintSHA1   []byte
	CertificateThumbprintSHA256 []byte
}

func (_ JSONWebKey) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (_ *JSONWebKey) IsPublic() bool {
	return false
}

func (_ *JSONWebKey) Public() JSONWebKey {
	return JSONWebKey{}
}

func (_ *JSONWebKey) Thumbprint(_ crypto.Hash) ([]byte, error) {
	return nil, nil
}

func (_ *JSONWebKey) UnmarshalJSON(_ []byte) error {
	return nil
}

func (_ *JSONWebKey) Valid() bool {
	return false
}

type JSONWebSignature struct {
	Signatures []Signature
}

func (_ JSONWebSignature) CompactSerialize() (string, error) {
	return "", nil
}

func (_ JSONWebSignature) DetachedCompactSerialize() (string, error) {
	return "", nil
}

func (_ JSONWebSignature) DetachedVerify(_ []byte, _ interface{}) error {
	return nil
}

func (_ JSONWebSignature) DetachedVerifyMulti(_ []byte, _ interface{}) (int, Signature, error) {
	return 0, Signature{}, nil
}

func (_ JSONWebSignature) FullSerialize() string {
	return ""
}

func (_ JSONWebSignature) UnsafePayloadWithoutVerification() []byte {
	return nil
}

func (_ JSONWebSignature) Verify(_ interface{}) ([]byte, error) {
	return nil, nil
}

func (_ JSONWebSignature) VerifyMulti(_ interface{}) (int, Signature, []byte, error) {
	return 0, Signature{}, nil, nil
}

type KeyAlgorithm string

func NewEncrypter(_ ContentEncryption, _ Recipient, _ *EncrypterOptions) (Encrypter, error) {
	return nil, nil
}

func NewSigner(_ SigningKey, _ *SignerOptions) (Signer, error) {
	return nil, nil
}

type NonceSource interface {
	Nonce() (string, error)
}

type Recipient struct {
	Algorithm  KeyAlgorithm
	Key        interface{}
	KeyID      string
	PBES2Count int
	PBES2Salt  []byte
}

type Signature struct {
	Header      Header
	Protected   Header
	Unprotected Header
	Signature   []byte
}

type SignatureAlgorithm string

type Signer interface {
	Options() SignerOptions
	Sign(_ []byte) (*JSONWebSignature, error)
}

type SignerOptions struct {
	NonceSource  NonceSource
	EmbedJWK     bool
	ExtraHeaders map[HeaderKey]interface{}
}

func (_ *SignerOptions) WithBase64(_ bool) *SignerOptions {
	return nil
}

func (_ *SignerOptions) WithContentType(_ ContentType) *SignerOptions {
	return nil
}

func (_ *SignerOptions) WithCritical(_ ...string) *SignerOptions {
	return nil
}

func (_ *SignerOptions) WithHeader(_ HeaderKey, _ interface{}) *SignerOptions {
	return nil
}

func (_ *SignerOptions) WithType(_ ContentType) *SignerOptions {
	return nil
}

type SigningKey struct {
	Algorithm SignatureAlgorithm
	Key       interface{}
}