/* btckeygenie v1.0.0
 * https://github.com/vsergeev/btckeygenie
 * License: MIT
 */

package btckey

import (
	"bytes"

	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

/******************************************************************************/
/* ECDSA Keypair Generation */
/******************************************************************************/

var curve = elliptic.P256()

// PublicKey represents a Bitcoin public key.
type PublicKey struct {
	X, Y *big.Int
}

// PrivateKey represents a Bitcoin private key.
type PrivateKey struct {
	PublicKey
	D *big.Int
}

func NewPrivateKey(d *big.Int) *PrivateKey {
	key := &PrivateKey{D: d}
	key.derive()
	return key
}

// derive derives a Bitcoin public key from a Bitcoin private key.
func (priv *PrivateKey) derive() (pub *PublicKey) {
	/* See Certicom's SEC1 3.2.1, pg.23 */

	/* Derive public key from Q = d*G */
	x, y := curve.ScalarBaseMult(priv.D.Bytes())

	/* Check that Q is on the curve */
	if !curve.IsOnCurve(x, y) {
		panic("Catastrophic math logic failure in public key derivation.")
	}

	priv.X = x
	priv.Y = y

	return &priv.PublicKey
}

// GenerateKey generates a public and private key pair using random source rand.
func GenerateKey(rand io.Reader) (priv PrivateKey, err error) {
	/* See Certicom's SEC1 3.2.1, pg.23 */
	/* See NSA's Suite B Implementer’s Guide to FIPS 186-3 (ECDSA) A.1.1, pg.18 */

	/* Select private key d randomly from [1, n) */

	/* Read N bit length random bytes + 64 extra bits  */
	b := make([]byte, curve.Params().N.BitLen()/8+8)
	_, err = io.ReadFull(rand, b)
	if err != nil {
		return priv, fmt.Errorf("Reading random reader: %v", err)
	}

	d := new(big.Int).SetBytes(b)

	/* Mod n-1 to shift d into [0, n-1) range */
	d.Mod(d, new(big.Int).Sub(curve.Params().N, big.NewInt(1)))
	/* Add one to shift d to [1, n) range */
	d.Add(d, big.NewInt(1))

	priv.D = d

	/* Derive public key from private key */
	priv.derive()

	return priv, nil
}

/******************************************************************************/
/* Base-58 Check Encode/Decode */
/******************************************************************************/

// b58checkencode encodes version ver and byte slice b into a base-58 check encoded string.
func b58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	bcpy := append([]byte{ver}, b...)

	/* Create a new SHA256 context */
	sha256_h := sha256.New()

	/* SHA256 Hash #1 */
	sha256_h.Reset()
	sha256_h.Write(bcpy)
	hash1 := sha256_h.Sum(nil)

	/* SHA256 Hash #2 */
	sha256_h.Reset()
	sha256_h.Write(hash1)
	hash2 := sha256_h.Sum(nil)

	/* Append first four bytes of hash */
	bcpy = append(bcpy, hash2[0:4]...)

	/* Encode base58 string */
	s = base58.Encode(bcpy)

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}

	return s
}

// b58checkdecode decodes base-58 check encoded string s into a version ver and byte slice b.
func b58checkdecode(s string) (ver uint8, b []byte, err error) {
	/* Decode base58 string */
	b = base58.Decode(s)

	/* Add leading zero bytes */
	for i := 0; i < len(s); i++ {
		if s[i] != '1' {
			break
		}
		b = append([]byte{0x00}, b...)
	}

	/* Verify checksum */
	if len(b) < 5 {
		return 0, nil, fmt.Errorf("Invalid base-58 check string: missing checksum.")
	}

	/* Create a new SHA256 context */
	sha256_h := sha256.New()

	/* SHA256 Hash #1 */
	sha256_h.Reset()
	sha256_h.Write(b[:len(b)-4])
	hash1 := sha256_h.Sum(nil)

	/* SHA256 Hash #2 */
	sha256_h.Reset()
	sha256_h.Write(hash1)
	hash2 := sha256_h.Sum(nil)

	/* Compare checksum */
	if bytes.Compare(hash2[0:4], b[len(b)-4:]) != 0 {
		return 0, nil, fmt.Errorf("Invalid base-58 check string: invalid checksum.")
	}

	/* Strip checksum bytes */
	b = b[:len(b)-4]

	/* Extract and strip version */
	ver = b[0]
	b = b[1:]

	return ver, b, nil
}

/******************************************************************************/
/* Bitcoin Private Key Import/Export */
/******************************************************************************/

// CheckWIF checks that string wif is a valid Wallet Import Format or Wallet Import Format Compressed string. If it is not, err is populated with the reason.
func CheckWIF(wif string) (valid bool, err error) {
	/* See https://en.bitcoin.it/wiki/Wallet_import_format */

	/* Base58 Check Decode the WIF string */
	ver, priv_bytes, err := b58checkdecode(wif)
	if err != nil {
		return false, err
	}

	/* Check that the version byte is 0x80 */
	if ver != 0x80 {
		return false, fmt.Errorf("Invalid WIF version 0x%02x, expected 0x80.", ver)
	}

	/* Check that private key bytes length is 32 or 33 */
	if len(priv_bytes) != 32 && len(priv_bytes) != 33 {
		return false, fmt.Errorf("Invalid private key bytes length %d, expected 32 or 33.", len(priv_bytes))
	}

	/* If the private key bytes length is 33, check that suffix byte is 0x01 (for compression) */
	if len(priv_bytes) == 33 && priv_bytes[len(priv_bytes)-1] != 0x01 {
		return false, fmt.Errorf("Invalid private key bytes, unknown suffix byte 0x%02x.", priv_bytes[len(priv_bytes)-1])
	}

	return true, nil
}

// ToBytes converts a Bitcoin private key to a 32-byte byte slice.
func (priv *PrivateKey) ToBytes() (b []byte) {
	d := priv.D.Bytes()

	/* Pad D to 32 bytes */
	padded_d := append(bytes.Repeat([]byte{0x00}, 32-len(d)), d...)

	return padded_d
}

// FromBytes converts a 32-byte byte slice to a Bitcoin private key and derives the corresponding Bitcoin public key.
func (priv *PrivateKey) FromBytes(b []byte) (err error) {
	if len(b) != 32 {
		return fmt.Errorf("Invalid private key bytes length %d, expected 32.", len(b))
	}

	priv.D = new(big.Int).SetBytes(b)

	/* Derive public key from private key */
	priv.derive()

	return nil
}

// ToWIF converts a Bitcoin private key to a Wallet Import Format string.
func (priv *PrivateKey) ToWIF() (wif string) {
	/* See https://en.bitcoin.it/wiki/Wallet_import_format */

	/* Convert the private key to bytes */
	priv_bytes := priv.ToBytes()

	/* Convert bytes to base-58 check encoded string with version 0x80 */
	wif = b58checkencode(0x80, priv_bytes)

	return wif
}

// ToWIFC converts a Bitcoin private key to a Wallet Import Format string with the public key compressed flag.
func (priv *PrivateKey) ToWIFC() (wifc string) {
	/* See https://en.bitcoin.it/wiki/Wallet_import_format */

	/* Convert the private key to bytes */
	priv_bytes := priv.ToBytes()

	/* Append 0x01 to tell Bitcoin wallet to use compressed public keys */
	priv_bytes = append(priv_bytes, []byte{0x01}...)

	/* Convert bytes to base-58 check encoded string with version 0x80 */
	wifc = b58checkencode(0x80, priv_bytes)

	return wifc
}

// FromWIF converts a Wallet Import Format string to a Bitcoin private key and derives the corresponding Bitcoin public key.
func (priv *PrivateKey) FromWIF(wif string) (err error) {
	/* See https://en.bitcoin.it/wiki/Wallet_import_format */

	/* Base58 Check Decode the WIF string */
	ver, priv_bytes, err := b58checkdecode(wif)
	if err != nil {
		return err
	}

	/* Check that the version byte is 0x80 */
	if ver != 0x80 {
		return fmt.Errorf("Invalid WIF version 0x%02x, expected 0x80.", ver)
	}

	/* If the private key bytes length is 33, check that suffix byte is 0x01 (for compression) and strip it off */
	if len(priv_bytes) == 33 {
		if priv_bytes[len(priv_bytes)-1] != 0x01 {
			return fmt.Errorf("Invalid private key, unknown suffix byte 0x%02x.", priv_bytes[len(priv_bytes)-1])
		}
		priv_bytes = priv_bytes[0:32]
	}

	/* Convert from bytes to a private key */
	err = priv.FromBytes(priv_bytes)
	if err != nil {
		return err
	}

	/* Derive public key from private key */
	priv.derive()

	return nil
}

/******************************************************************************/
/* Bitcoin Public Key Import/Export */
/******************************************************************************/

// ToBytes converts a Bitcoin public key to a 33-byte byte slice with point compression.
func (pub *PublicKey) ToBytes() (b []byte) {
	/* See Certicom SEC1 2.3.3, pg. 10 */

	x := pub.X.Bytes()

	/* Pad X to 32-bytes */
	padded_x := append(bytes.Repeat([]byte{0x00}, 32-len(x)), x...)

	/* Add prefix 0x02 or 0x03 depending on ylsb */
	if pub.Y.Bit(0) == 0 {
		return append([]byte{0x02}, padded_x...)
	}

	return append([]byte{0x03}, padded_x...)
}

// ToBytesUncompressed converts a Bitcoin public key to a 65-byte byte slice without point compression.
func (pub *PublicKey) ToBytesUncompressed() (b []byte) {
	/* See Certicom SEC1 2.3.3, pg. 10 */

	x := pub.X.Bytes()
	y := pub.Y.Bytes()

	/* Pad X and Y coordinate bytes to 32-bytes */
	padded_x := append(bytes.Repeat([]byte{0x00}, 32-len(x)), x...)
	padded_y := append(bytes.Repeat([]byte{0x00}, 32-len(y)), y...)

	/* Add prefix 0x04 for uncompressed coordinates */
	return append([]byte{0x04}, append(padded_x, padded_y...)...)
}

// FromBytes converts a byte slice (either with or without point compression) to a Bitcoin public key.
func (pub *PublicKey) FromBytes(b []byte) (err error) {
	/* See Certicom SEC1 2.3.4, pg. 11 */

	if len(b) < 33 {
		return fmt.Errorf("Invalid public key bytes length %d, expected at least 33.", len(b))
	}

	if b[0] == 0x02 || b[0] == 0x03 {
		/* Compressed public key */

		if len(b) != 33 {
			return fmt.Errorf("Invalid public key bytes length %d, expected 33.", len(b))
		}

		P, err := secp256k1.Decompress(new(big.Int).SetBytes(b[1:33]), uint(b[0]&0x1))
		if err != nil {
			return fmt.Errorf("Invalid compressed public key bytes, decompression error: %v", err)
		}

		pub.X = P.X
		pub.Y = P.Y

	} else if b[0] == 0x04 {
		/* Uncompressed public key */

		if len(b) != 65 {
			return fmt.Errorf("Invalid public key bytes length %d, expected 65.", len(b))
		}

		pub.X = new(big.Int).SetBytes(b[1:33])
		pub.Y = new(big.Int).SetBytes(b[33:65])

		/* Check that the point is on the curve */
		if !secp256k1.IsOnCurve(pub.Point) {
			return fmt.Errorf("Invalid public key bytes: point not on curve.")
		}

	} else {
		return fmt.Errorf("Invalid public key prefix byte 0x%02x, expected 0x02, 0x03, or 0x04.", b[0])
	}

	return nil
}

// ToAddress converts a Bitcoin public key to a compressed Bitcoin address string.
func (pub *PublicKey) ToAddress() (address string) {
	/* See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses */

	/* Convert the public key to bytes */
	pub_bytes := pub.ToBytes()

	/* SHA256 Hash */
	sha256_h := sha256.New()
	sha256_h.Reset()
	sha256_h.Write(pub_bytes)
	pub_hash_1 := sha256_h.Sum(nil)

	/* RIPEMD-160 Hash */
	ripemd160_h := ripemd160.New()
	ripemd160_h.Reset()
	ripemd160_h.Write(pub_hash_1)
	pub_hash_2 := ripemd160_h.Sum(nil)

	/* Convert hash bytes to base58 check encoded sequence */
	address = b58checkencode(0x00, pub_hash_2)

	return address
}

// ToAddressUncompressed converts a Bitcoin public key to an uncompressed Bitcoin address string.
func (pub *PublicKey) ToAddressUncompressed() (address string) {
	/* See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses */

	/* Convert the public key to bytes */
	pub_bytes := pub.ToBytesUncompressed()

	/* SHA256 Hash */
	sha256_h := sha256.New()
	sha256_h.Reset()
	sha256_h.Write(pub_bytes)
	pub_hash_1 := sha256_h.Sum(nil)

	/* RIPEMD-160 Hash */
	ripemd160_h := ripemd160.New()
	ripemd160_h.Reset()
	ripemd160_h.Write(pub_hash_1)
	pub_hash_2 := ripemd160_h.Sum(nil)

	/* Convert hash bytes to base58 check encoded sequence */
	address = b58checkencode(0x00, pub_hash_2)

	return address
}
