// Package passwordhash implements Argon2id password hashing with a
// self-describing PHC-like encoded string, so params can change over time
// without breaking verification of older hashes.
package passwordhash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Params controls Argon2id cost. Tune Memory/Iterations for your hardware;
// these are reasonable defaults for an interactive login endpoint.
type Params struct {
	Memory      uint32 // KiB
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var DefaultParams = Params{
	Memory:      64 * 1024, // 64 MB
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

var (
	ErrInvalidHash         = errors.New("passwordhash: invalid encoded hash format")
	ErrIncompatibleVersion = errors.New("passwordhash: incompatible argon2 version")
)

// Hash generates a salted Argon2id hash encoded as:
// $argon2id$v=19$m=65536,t=3,p=2$<salt>$<hash>
func Hash(password string) (string, error) {
	return HashWithParams(password, DefaultParams)
}

func HashWithParams(password string, p Params) (string, error) {
	salt := make([]byte, p.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Iterations, p.Parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
	return encoded, nil
}

// Verify checks a plaintext password against a previously encoded hash,
// using constant-time comparison to avoid timing side-channels.
func Verify(password, encodedHash string) (bool, error) {
	p, salt, hash, err := decode(encodedHash)
	if err != nil {
		return false, err
	}

	candidate := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	if subtle.ConstantTimeCompare(hash, candidate) == 1 {
		return true, nil
	}
	return false, nil
}

func decode(encodedHash string) (Params, []byte, []byte, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return Params{}, nil, nil, ErrInvalidHash
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return Params{}, nil, nil, err
	}
	if version != argon2.Version {
		return Params{}, nil, nil, ErrIncompatibleVersion
	}

	p := Params{}
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism); err != nil {
		return Params{}, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return Params{}, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return Params{}, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
