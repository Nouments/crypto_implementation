# Caesar Cipher (Byte / ASCII Version)

The Caesar cipher is a classical substitution cipher where each byte in the plaintext is shifted by a fixed number of positions. While traditionally it only shifts letters, in real-life applications you can apply it to **all ASCII characters** (0–255).

## How it works

- Each byte in the plaintext is replaced by another byte a fixed number of positions away.  
- The fixed number of positions is called the **key** or **shift**.  
- After reaching the maximum byte value (255), the count wraps around back to 0.

### Mathematical Explanation

Let each byte be represented by a number from 0 to 255.  

**Encryption formula (byte-wise):**  

$$
C = (P + k) \mod 256
$$`

- `P` = numeric value of the plaintext byte  
- `k` = key (shift value, 0–255)  
- `C` = numeric value of the ciphertext byte  

**Decryption formula (byte-wise):**  

$$
P = (C - k + 256) \mod 256
$$

- This ensures wrap-around works both for positive overflow and negative subtraction.

### Example

- Plaintext: `"test"` → bytes `[116, 101, 115, 116]`  
- Key: `2`  
- Ciphertext bytes: `[118, 103, 117, 118]`  
- Decrypted back: `[116, 101, 115, 116]` → `"test"`

Notes:  

- All bytes (letters, numbers, punctuation, special ASCII characters) are shifted.  
- Wrap-around ensures no overflow or negative values.  
- ROT13 can be implemented on ASCII bytes as well, not just letters.  
- This is the version suitable for **real-world text and binary data**, not just alphabetic characters. 