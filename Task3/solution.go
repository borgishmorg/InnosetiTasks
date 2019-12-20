package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const maxKeyCount = 916132832
const keyLen = 5
const alp = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const radix = len(alp)

type key struct {
	key    string
	closed bool
}

var keys map[string]*key = make(map[string]*key)
var mux sync.Mutex

func intToKey(n int) *key {
	keyRep := ""
	for i := 0; i < keyLen; i++ {
		keyRep = string(alp[n%radix]) + keyRep
		n /= radix
	}
	k := new(key)
	k.key = keyRep
	k.closed = false
	return k
}

func open(w http.ResponseWriter, r *http.Request) {
	var msg string
	mux.Lock()
	defer mux.Unlock()
	defer func() {
		io.WriteString(w, msg)
		log.Println(msg)
	}()

	nextKey := len(keys)
	if nextKey >= maxKeyCount {
		msg = "Не выданных ключей не осталось :("
		return
	}
	k := intToKey(nextKey)
	msg = k.key
	keys[k.key] = k
}

func close(w http.ResponseWriter, r *http.Request) {
	var msg string
	mux.Lock()
	defer mux.Unlock()
	defer func() {
		io.WriteString(w, msg)
		log.Println(msg)
	}()
	kReq := string(r.URL.EscapedPath()[len("/close/"):])

	if len(kReq) == 0 {
		msg = "Ошибка. Необходимо указать ключ"
		return
	}

	k, ok := keys[kReq]
	if !ok {
		msg = fmt.Sprintf("Ключ \"%s\" еще не выдан или не возможен", kReq)
		return
	}

	if !k.closed {
		k.closed = true
		msg = fmt.Sprintf("Ключ \"%s\" успешно закрыт", k.key)
	} else {
		msg = fmt.Sprintf("Ключ \"%s\" уже закрыт", k.key)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	var msg string
	mux.Lock()
	defer mux.Unlock()
	defer func() {
		io.WriteString(w, msg)
		log.Println(msg)
	}()
	kReq := string(r.URL.EscapedPath()[len("/info/"):])

	if len(kReq) == 0 {
		msg = fmt.Sprintf("Число использованных ключей: %d\n", len(keys))
		msg += fmt.Sprintf("Число оставшихся ключей: %d", maxKeyCount-len(keys))
		return
	}

	k, ok := keys[kReq]
	if !ok {
		msg = fmt.Sprintf("Ключ \"%s\" не выдан или не возможен", kReq)
		return
	}

	if !k.closed {
		msg = fmt.Sprintf("Ключ \"%s\" выдан", kReq)
	} else {
		msg = fmt.Sprintf("Ключ \"%s\" закрыт", kReq)
	}
}

func main() {
	http.HandleFunc("/open/", open)
	http.HandleFunc("/close/", close)
	http.HandleFunc("/info/", info)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
