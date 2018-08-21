/*
SUMMARY
this is an example of simple encryption and
decryption with Go. taken from
https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
*/

package main

import (

	//https://golang.org/pkg/crypto/aes/
	/*implements AES encryption (formerly Rijndael), 
	as defined in U.S. Federal Information Processing Standards Publication 197.*/
	"crypto/aes"		

	//https://golang.org/pkg/crypto/cipher/
	/*implements standard block cipher modes that can be wrapped around low-level 
	block cipher implementations.*/
	"crypto/cipher"		

	//https://golang.org/pkg/crypto/md5/
	/* implements the MD5 hash algorithm as defined in RFC 1321. */
	"crypto/md5"		

	//https://golang.org/pkg/crypto/rand/
	/* implements a cryptographically secure pseudorandom number generator.*/
	"crypto/rand"	

	//https://golang.org/pkg/encoding/hex/
	/*implements hexadecimal encoding and decoding.*/	
	"encoding/hex"

	//https://golang.org/pkg/fmt/
	/*implements formatted I/O with functions analogous to C's printf and scanf.*/	
	"fmt"	

	//https://golang.org/pkg/io/
	/*provides basic interfaces to I/O primitives.*/		
	"io"			

	//https://golang.org/pkg/io/ioutil/
	/*implements some I/O utility functions.*/
	"io/ioutil"		

	//https://golang.org/pkg/os/
	/*provides a platform-independent interface to operating system functionality.*/
	"os"			
)

func createHash(key string) string {
	hasher := md5.New()				//crypto/md5
	hasher.Write([]byte(key))			//os
	return hex.EncodeToString(hasher.Sum(nil))	//encoding/hex
}

//initialize the cypher block
//choose a block cypher mode
//generate a randomized nonce
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase))) 	//crypto/aes
	gcm, err := cipher.NewGCM(block)				//crypto/cipher 
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())				//crypto/cipher
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil { 	 //io
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil) 		//crypto/cipher
	return ciphertext
}
//initialize the block cypher
//choose a block cypher mode
/*NOTE: A nonce is a number or string used only once. 
This is useful for generating a unique token for login pages to 
prevent duplicate or unauthorized submissions*/
//we stored the nonce at the beginning of the encrypted data
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))	//crypto/aes
	block, err := aes.NewCipher(key) 	//crypto/cipher
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)	//crypto/cipher
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()	//crypto/cipher
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {
	fmt.Println("Starting the application...")
	ciphertext := encrypt([]byte("Hello World"), "password")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := decrypt(ciphertext, "password")
	fmt.Printf("Decrypted: %s\n", plaintext)
	encryptFile("sample.txt", []byte("Hello World"), "password1")
	fmt.Println(string(decryptFile("sample.txt", "password1")))
}