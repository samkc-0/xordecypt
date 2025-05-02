sí, aquí va — terse, clear, hacker-friendly, no bs:

⸻

📜 README.md

# xordecrypt

bruteforce xor decryptor for ciphertexts with known plaintext prefixes.

useful for ctf challenges, key recovery, and xor puzzles.

---

## 🔧 usage

```bash
xordecrypt -c HEXSTRING -p PREFIX -keylen N

	•	-c — hex-encoded ciphertext string (e.g. "0d3c28...")
	•	-p — known plaintext prefix (e.g. "THM{")
	•	-l — total length of repeating xor key

🚨 all flags are required. no defaults.

⸻

🧪 example

xordecrypt -c "$(cat key.txt)" -p "THM{" -l 5

or with stdin fallback:

cat key.txt | xordecrypt -p "THM{" -l 5



⸻

💡 what it does
	1.	xor-decodes known plaintext against ciphertext to recover part of the key
	2.	bruteforces remaining key bytes using alphanumeric keyspace
	3.	prints any valid-looking plaintexts

⸻

⚙️ build

go build -o xordecrypt



⸻

🔬 tests

go test



⸻

🚀 roadmap (?)
	•	support multibyte bruteforce
	•	smarter scoring / autoclassify results
	•	tui / curses mode for interactive decoding

---

