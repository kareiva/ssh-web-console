package utils

import (
	"errors"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

// HostKeyError is returned (wrapped) by Connect when the remote host key is not trusted.
// The handshake is aborted before user authentication, so no password is sent.
type HostKeyError struct {
	Changed     bool   `json:"changed"` // a different key of the same type was trusted before
	KeyType     string `json:"type"`
	Fingerprint string `json:"fingerprint"`
}

func (e *HostKeyError) Error() string {
	if e.Changed {
		return "host key changed: " + e.KeyType + " " + e.Fingerprint
	}
	return "unknown host key: " + e.KeyType + " " + e.Fingerprint
}

// VerifyHostKey trusts a key listed in the optional known_hosts file (ssh.known_hosts in config),
// or the SHA256 fingerprint confirmed by the user. The accepted key is stored in *accepted.
func VerifyHostKey(confirmed string, accepted *ssh.PublicKey) ssh.HostKeyCallback {
	return func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		fp := ssh.FingerprintSHA256(key)
		if Config.SSH.KnownHosts != "" {
			check, err := knownhosts.New(Config.SSH.KnownHosts)
			if err != nil {
				log.Println("Error: cannot load known_hosts:", err)
			} else {
				var ke *knownhosts.KeyError
				err := check(hostname, remote, key)
				if err == nil {
					*accepted = key
					return nil
				}
				if errors.As(err, &ke) {
					// only a known key of the same type means the key changed; another type is just unknown
					for _, w := range ke.Want {
						if w.Key.Type() == key.Type() {
							return &HostKeyError{Changed: true, KeyType: key.Type(), Fingerprint: fp}
						}
					}
				}
			}
		}
		if confirmed == fp {
			*accepted = key
			return nil
		}
		return &HostKeyError{Changed: confirmed != "", KeyType: key.Type(), Fingerprint: fp}
	}
}

// PinnedHostKey accepts only the host key accepted at sign-in (ssh wire format).
func PinnedHostKey(marshaled []byte) ssh.HostKeyCallback {
	key, err := ssh.ParsePublicKey(marshaled)
	if err != nil {
		return func(string, net.Addr, ssh.PublicKey) error {
			return errors.New("no host key pinned for this session")
		}
	}
	return ssh.FixedHostKey(key)
}
