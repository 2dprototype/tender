## Stdlib `crypto`

The `crypto` module provides comprehensive cryptographic functionalities including hashing, encryption/decryption, digital signatures, key generation, and secure random number generation.

## Hash Functions

### `md5(input)`
Generates an MD5 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: MD5 hash as a hexadecimal string
- **Example**: `md5_hash := crypto.md5("hello")`

### `sha1(input)`
Generates a SHA-1 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA-1 hash as a hexadecimal string

### `sha224(input)`
Generates a SHA-224 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA-224 hash as a hexadecimal string

### `sha256(input)`
Generates a SHA-256 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA-256 hash as a hexadecimal string

### `sha384(input)`
Generates a SHA-384 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA-384 hash as a hexadecimal string

### `sha512(input)`
Generates a SHA-512 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA-512 hash as a hexadecimal string

### `sha3_224(input)`
Generates a SHA3-224 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA3-224 hash as a hexadecimal string

### `sha3_256(input)`
Generates a SHA3-256 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA3-256 hash as a hexadecimal string

### `sha3_384(input)`
Generates a SHA3-384 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA3-384 hash as a hexadecimal string

### `sha3_512(input)`
Generates a SHA3-512 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: SHA3-512 hash as a hexadecimal string

### `blake2b_256(input)`
Generates a BLAKE2b-256 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: BLAKE2b-256 hash as a hexadecimal string

### `blake2b_512(input)`
Generates a BLAKE2b-512 hash for the given input.
- **Parameters**: `input` - The data to hash (string or byte array)
- **Returns**: BLAKE2b-512 hash as a hexadecimal string

## HMAC Functions

### `hmac.md5(key, data)`
Generates HMAC using MD5 hash function.
- **Parameters**: `key` - HMAC key (string or byte array), `data` - Data to authenticate (string or byte array)
- **Returns**: HMAC as hexadecimal string

### `hmac.sha1(key, data)`
Generates HMAC using SHA-1 hash function.

### `hmac.sha256(key, data)`
Generates HMAC using SHA-256 hash function.

### `hmac.sha384(key, data)`
Generates HMAC using SHA-384 hash function.

### `hmac.sha512(key, data)`
Generates HMAC using SHA-512 hash function.

### `hmac.sha3_256(key, data)`
Generates HMAC using SHA3-256 hash function.

### `hmac.sha3_512(key, data)`
Generates HMAC using SHA3-512 hash function.

## AES Encryption (crypto.aes)

### `aes.encrypt(plaintext, key)`
Encrypts data using AES encryption.
- **Parameters**: `plaintext` - Data to encrypt (string or byte array), `key` - Encryption key (string or byte array, 16/24/32 bytes for AES-128/192/256)
- **Returns**: Encrypted data as byte array

### `aes.decrypt(ciphertext, key)`
Decrypts AES-encrypted data.
- **Parameters**: `ciphertext` - Encrypted data (byte array), `key` - Decryption key (string or byte array)
- **Returns**: Decrypted data as byte array

### `aes.block_size`
AES block size constant (16 bytes).

## RSA Encryption (crypto.rsa)

### `rsa.generate_key(bits)`
Generates RSA key pair.
- **Parameters**: `bits` - Key size in bits (default: 2048)
- **Returns**: Map with `private` and `public` keys in DER format

### `rsa.encrypt(data, public_key)`
Encrypts data using RSA public key.
- **Parameters**: `data` - Data to encrypt (string or byte array), `public_key` - Public key (DER bytes or PEM string)
- **Returns**: Encrypted data as byte array

### `rsa.decrypt(ciphertext, private_key)`
Decrypts RSA-encrypted data.
- **Parameters**: `ciphertext` - Encrypted data (byte array), `private_key` - Private key (DER bytes or PEM string)
- **Returns**: Decrypted data as byte array

### `rsa.sign(data, private_key)`
Signs data using RSA private key.
- **Parameters**: `data` - Data to sign (string or byte array), `private_key` - Private key (DER bytes or PEM string)
- **Returns**: Signature as byte array

### `rsa.verify(data, signature, public_key)`
Verifies RSA signature.
- **Parameters**: `data` - Original data (string or byte array), `signature` - Signature to verify (byte array), `public_key` - Public key (DER bytes or PEM string)
- **Returns**: Boolean indicating if signature is valid

### `rsa.export_key(key, key_type)`
Exports key to PEM format.
- **Parameters**: `key` - Key in DER format (byte array), `key_type` - "private" or "public"
- **Returns**: PEM-encoded key as string

### `rsa.import_key(pem_data, key_type)`
Imports key from PEM format.
- **Parameters**: `pem_data` - PEM-encoded key (string or byte array), `key_type` - "private" or "public"
- **Returns**: Key in DER format as byte array

## ECDSA Digital Signatures (crypto.ecdsa)

### `ecdsa.generate_key(curve)`
Generates ECDSA key pair.
- **Parameters**: `curve` - Elliptic curve name: "p224", "p256", "p384", or "p521" (default: "p256")
- **Returns**: Map with `private` and `public` keys in DER format

### `ecdsa.sign(data, private_key)`
Signs data using ECDSA private key.
- **Parameters**: `data` - Data to sign (string or byte array), `private_key` - Private key (DER bytes or PEM string)
- **Returns**: 64-byte signature (concatenated r and s values)

### `ecdsa.verify(data, signature, public_key)`
Verifies ECDSA signature.
- **Parameters**: `data` - Original data (string or byte array), `signature` - 64-byte signature (byte array), `public_key` - Public key (DER bytes or PEM string)
- **Returns**: Boolean indicating if signature is valid

### `ecdsa.export_key(key, key_type)`
Exports ECDSA key to PEM format.

### `ecdsa.import_key(pem_data, key_type)`
Imports ECDSA key from PEM format.

## Ed25519 Digital Signatures (crypto.ed25519)

### `ed25519.generate_key()`
Generates Ed25519 key pair.
- **Returns**: Map with `private` and `public` keys as byte arrays

### `ed25519.sign(data, private_key)`
Signs data using Ed25519 private key.
- **Parameters**: `data` - Data to sign (string or byte array), `private_key` - Private key (byte array)
- **Returns**: Signature as byte array

### `ed25519.verify(data, signature, public_key)`
Verifies Ed25519 signature.
- **Parameters**: `data` - Original data (string or byte array), `signature` - Signature (byte array), `public_key` - Public key (byte array)
- **Returns**: Boolean indicating if signature is valid

## Key Derivation Functions

### `pbkdf2(password, salt, iterations, key_len, hash_func)`
Derives key using PBKDF2.
- **Parameters**: `password` - Password (string or byte array), `salt` - Salt (string or byte array), `iterations` - Number of iterations, `key_len` - Desired key length, `hash_func` - Hash function name ("sha1", "sha256", "sha512", "sha3_256", "sha3_512")
- **Returns**: Derived key as byte array

### `bcrypt(password, cost)`
Hashes password using bcrypt.
- **Parameters**: `password` - Password (string or byte array), `cost` - Computational cost (default: 10)
- **Returns**: bcrypt hash as byte array

### `scrypt(password, salt, key_len, N, r, p)`
Derives key using scrypt.
- **Parameters**: `password` - Password (string or byte array), `salt` - Salt (string or byte array), `key_len` - Desired key length, `N` - CPU/memory cost, `r` - Block size, `p` - Parallelization
- **Returns**: Derived key as byte array

## Argon2 Password Hashing (crypto.argon2)

### `argon2.id(password, salt, time_cost, memory_cost, threads, key_len)`
Derives key using Argon2id.
- **Parameters**: `password` - Password (string or byte array), `salt` - Salt (string or byte array), `time_cost` - Time cost (iterations), `memory_cost` - Memory cost in KiB, `threads` - Parallelism, `key_len` - Desired key length
- **Returns**: Derived key as byte array

### `argon2.i(password, salt, time_cost, memory_cost, threads, key_len)`
Derives key using Argon2i.

## Random Generation (crypto.random)

### `random.bytes(size)`
Generates cryptographically secure random bytes.
- **Parameters**: `size` - Number of bytes to generate
- **Returns**: Random bytes as byte array

### `random.int(min, max)`
Generates cryptographically secure random integer.
- **Parameters**: `min` - Minimum value, `max` - Maximum value
- **Returns**: Random integer

### `random.float()`
Generates cryptographically secure random float.
- **Returns**: Random float between 0.0 and 1.0

### `random.uuid()`
Generates random UUID v4.
- **Returns**: UUID as string

## Security Utilities

### `constant_time_compare(a, b)`
Compares two values in constant time to prevent timing attacks.
- **Parameters**: `a`, `b` - Values to compare (strings or byte arrays)
- **Returns**: Boolean indicating if values are equal

## Example Usage

```javascript
import "crypto"

// Hashing examples
data := "Hello, World!"
println("MD5:", crypto.md5(data))
println("SHA256:", crypto.sha256(data))

// HMAC examples
key := "secret"
println("HMAC-SHA256:", crypto.hmac.sha256(key, data))

// AES encryption
aes_key := "0123456789abcdef0123456789abcdef" // 32 bytes
encrypted := crypto.aes.encrypt("secret message", aes_key)
decrypted := crypto.aes.decrypt(encrypted, aes_key)
println("Decrypted:", string(decrypted))

// RSA encryption
rsa_keys := crypto.rsa.generate_key(2048)
encrypted_rsa := crypto.rsa.encrypt("secret", rsa_keys.public)
decrypted_rsa := crypto.rsa.decrypt(encrypted_rsa, rsa_keys.private)
println("RSA Decrypted:", string(decrypted_rsa))

// Digital signatures
ec_keys := crypto.ecdsa.generate_key("p256")
signature := crypto.ecdsa.sign(data, ec_keys.private)
is_valid := crypto.ecdsa.verify(data, signature, ec_keys.public)
println("ECDSA Signature valid:", is_valid)

// Password hashing
password := "my_password"
salt := crypto.random.bytes(16)
argon_hash := crypto.argon2.id(password, salt, 3, 64*1024, 4, 32)
println("Argon2 hash:", argon_hash)

// Random generation
println("Random bytes:", crypto.random.bytes(32))
println("Random int:", crypto.random.int(1, 100))
println("UUID:", crypto.random.uuid())
```
