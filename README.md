# Go Basic Cryptography Algorithms

This repository contains implementations of **basic cryptographic algorithms** in Go, focusing on **mathematical foundations** and simple substitution or transformation techniques.  
The goal is educational: understand how these algorithms work and practice implementing them from scratch.

## Implemented Algorithms

The project includes the following basic cryptography algorithms:

### 1. **Caesar Cipher**
- Simple substitution cipher.  
- Each letter (or byte) is shifted by a fixed key.  
- Formula (byte-wise): `C = P + k` (wrap-around automatic in Go `byte`)  
- Example: `"HELLO"` + key 3 → `"KHOOR"`  

### 2. **ROT-n / ROT13**
- Special case of Caesar Cipher.  
- Rotates letters by n positions (ROT13 rotates by 13).  
- Can be applied to all letters or ASCII bytes.  

### 3. **Vigenère Cipher**
- Polyalphabetic substitution cipher.  
- Each character is shifted by the corresponding character in a repeated key.  
- Byte-wise formula: `C_i = message[i] + key[i % len(key)]`  
- Works on letters or all ASCII bytes.  

### 4. **Affine Cipher**
- A mathematical cipher using a linear function over modulo 26 (or 256 for bytes).  
- Encryption: `C = (a*P + b) mod m`  
- Decryption: `P = a_inv * (C - b) mod m`  
- Introduces multiplication for more complex substitutions.  

### 5. **Substitution Cipher (Random / Fixed Table)**
- Each symbol is replaced with another according to a **mapping table**.  
- Can be implemented with letters, numbers, or full ASCII.  

### 6. **XOR Cipher**
- Simple symmetric cipher using bitwise XOR.  
- `C = P ⊕ K`  
- Very fast and can work on bytes, files, and strings.  
- Can be extended to a repeating key (like Vigenère with XOR).  

### 7. **ROT by Bytes / Extended Caesar**
- Applies the Caesar logic to **all bytes** of a message or file.  
- Demonstrates wrap-around naturally in Go using `byte`.  

---

## Project Goals

- **Educational**: understand basic cryptographic transformations.  
- **Mathematical foundation**: show formulas and byte-level operations.  
- **Hands-on implementation**: all algorithms coded in Go from scratch.  
- **Extendable**: easy to add new simple ciphers in the future.  

---

