package crypto

import "crypto/md5"
import "crypto/hmac"
import "crypto/sha256"
import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "code.google.com/p/go.crypto/ssh"

import "encoding/hex"
import "encoding/pem"
import "encoding/base64"

import "code.google.com/p/go.crypto/bcrypt"

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func Crypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComputeMD5(message string) string {
	hasher := md5.New()
	hasher.Write([]byte(message))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GenerateSSHKey() (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		return "", ""
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		    Type:    "RSA PRIVATE KEY",
			    Headers: nil,
				    Bytes:   privateKeyDer,

	}
	privateKeyPem := string(pem.EncodeToMemory(&privateKeyBlock))

	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", ""
	}
	publickeyPem := ssh.MarshalAuthorizedKey(publicKey)

	return privateKeyPem, string(publickeyPem)
}
