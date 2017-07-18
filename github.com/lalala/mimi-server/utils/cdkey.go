package utils

import (
	"crypto/rc4"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var rc4key = []byte("Cdkey0509$^&@")
var _codes = map[rune]rune{
	'o': '@',
	'O': '#',
	'/': '$',
}

var _reverseCodes map[rune]rune

func init() {
	_reverseCodes = make(map[rune]rune)
	for k, v := range _codes {
		_reverseCodes[v] = k
	}
}

func beautifyBase64(s string) string {
	return strings.Map(func(r rune) rune {
		v, ok := _codes[r]
		if ok {
			return v
		}
		return r
	}, s)
}

func deBeautifyBase64(s string) string {
	return strings.Map(func(r rune) rune {
		v, ok := _reverseCodes[r]
		if ok {
			return v
		}
		return r
	}, s)
}

// ________________________________________________________
// | 16 bits    |  16 bits   |    8 bits   |    8 bits |
// ________________________________________________________
// |  batch     |     no     |   reserved  |  verify_bits |
//    -------  rc4 -------
// batch: 批次
// no: 序号
// verify_bits: 前40位对2^8取余
func GenerateCdkey(batch int, num int) []string {
	maxBatch := int(math.Pow(2, 16))
	if batch < 1 || batch >= maxBatch {
		log.Panic("invalid batch")
	}
	maxNum := int(math.Pow(2, 16))
	if num < 1 || num >= maxNum {
		log.Panic("invalid num")
	}
	keys := make([]string, 0)
	reserved := uint64(0xff)
	for i := 0; i < num; i++ {
		no := i + 1
		encrypt := uint64(uint64(batch)<<16 | uint64(no))
		verify := encrypt % 256
		// log.Printf("encrypt: %d --> %b\n", encrypt, encrypt)
		// log.Printf("verify: %d --> %b\n", verify, verify)
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, encrypt)
		rc, err := rc4.NewCipher(rc4key)
		if err != nil {
			log.Panic(err)
		}
		rc.XORKeyStream(buf, buf)
		v := ((binary.LittleEndian.Uint64(buf) << 16) | reserved<<8 | verify)
		// log.Printf("v: %d --> %b\n", v, v)
		keybuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(keybuf, v)
		key := base64.RawStdEncoding.EncodeToString(keybuf)
		key = string(key[:8])
		key = beautifyBase64(key)
		keys = append(keys, key)
		// log.Println("key:", key)
	}
	return keys
}

func VerifyCdkey(key string) (batch int, no int, valid bool) {
	valid = false
	key = deBeautifyBase64(key)
	base64Key := fmt.Sprintf("%sAAA", key)
	data, err := base64.RawStdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Println(err)
		return
	}
	v := binary.LittleEndian.Uint64(data)
	verify := uint64(v & 0xff)
	reserved := uint64((v >> 8) & 0xff)
	// log.Printf("v: %d --> %b\n", v, v)
	// log.Printf("verify: %d --> %b\n", verify, verify)
	// log.Printf("reserved: %d --> %b\n", reserved, reserved)
	if reserved != 0xff {
		return
	}
	rc, err := rc4.NewCipher(rc4key)
	if err != nil {
		log.Println(err)
		return
	}

	v0 := (v >> 16 & 0xffffffff)
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, v0)
	rc.XORKeyStream(buf, buf)
	encrypt := binary.LittleEndian.Uint64(buf)
	if encrypt%256 != verify {
		log.Println(verify, encrypt, encrypt%256)
		return
	}
	no = int(encrypt & 0xffff)
	batch = int((encrypt >> 16) & 0xffff)
	valid = true
	return
}

func GenerateInvitationCode(accountId int64) string {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(accountId))
	rc, err := rc4.NewCipher(rc4key)
	if err != nil {
		log.Panic(err)
	}
	rc.XORKeyStream(buf, buf)
	code := base64.RawStdEncoding.EncodeToString(buf)
	return beautifyBase64(string(code[:8]))
}
