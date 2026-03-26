package main

import "fmt"

func EncrypteMessage(key byte, message string) string {
	data := []byte(message)
	for i, b := range data {
		data[i] = (b + key)
	}
	return string(data)
}

func DerypteMessage(key byte, encrypted string)  string {
	data := []byte(encrypted)
	for i, b := range data {
		data[i] = (b - key)
	}
	return string(data)
}


func main() {
	message := "ctf{ju5t3_un_t35t}"
	key := byte(2)

	encrypted := EncrypteMessage(key, message)
	fmt.Println("Encrypted :", string(encrypted))

	decrypted := DerypteMessage(key, encrypted)
	fmt.Println("Decrypted text:", decrypted)
}