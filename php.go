package php

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"html"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"
)

const GO_TIME_LAYOUT = "2006-01-02 15:04:05"

func Time() int64 {
	return time.Now().Unix()
}

func Date(format string, timestamp int64) string {
	if len(format) > 0 {
		return time.Unix(timestamp, 0).Format(format)
	}

	return time.Unix(timestamp, 0).Format(GO_TIME_LAYOUT)
}

func Strtotime(datetime string, format ...string) int64 {
	if len(format) > 0 {
		t, e := time.Parse(format[0], datetime)
		if e != nil {
			return 0
		}

		return t.Unix()
	}

	t, e := time.Parse(GO_TIME_LAYOUT, datetime)
	if e != nil {
		return 0
	}

	return t.Unix()
}

func Sleep(sec int64) {
	time.Sleep(time.Duration(sec) * time.Second)
}

func Usleep(usec int64) {
	time.Sleep(time.Duration(usec) * time.Microsecond)
}

/////////////
// strings //
/////////////

func Strlen(str string) int {
	return len(str)
}

func Mb_strlen(str string) int {
	return utf8.RuneCountInString(str)
}

func Str_replace(search, replace, subject string, count ...int) string {
	if len(count) > 0 {
		return strings.Replace(subject, search, replace, count[0])
	}

	return strings.ReplaceAll(subject, search, replace)
}

func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

func Strtolower(str string) string {
	return strings.ToLower(str)
}

func Ucwords(str string) string {
	return strings.Title(str)
}

func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}

	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}

	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}

	return str[start:end]
}

func Strrev(str string) string {
	new_str := ""
	ll := len(str)

	for ; ll > 0; ll-- {
		new_str += str[ll-1 : ll]
	}

	return new_str
}

func Trim(str string, charlist ...string) string {
	def := "\t\n\r\\0\\x0B "

	if len(charlist) > 0 {
		def = charlist[0]
	}

	return strings.Trim(str, def)
}

func Ltrim(str string, charlist ...string) string {
	def := "\t\n\r\\0\\x0B "

	if len(charlist) > 0 {
		def = charlist[0]
	}

	return strings.TrimLeft(str, def)
}

func Rtrim(str string, charlist ...string) string {
	def := "\t\n\r\\0\\x0B "

	if len(charlist) > 0 {
		def = charlist[0]
	}

	return strings.TrimRight(str, def)
}

func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

func Implode(glue string, arr []string) string {
	var buffer bytes.Buffer
	ll := len(arr)

	for _, str := range arr {
		buffer.WriteString(str)
		ll--

		if ll > 0 {
			buffer.WriteString(glue)
		}
	}
	return buffer.String()
}

func Chr(code int) string {
	return string(code)
}

func Ord(char string) int {
	r, _ := utf8.DecodeRune([]byte(char))
	return int(r)
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}

func Md5_file(filename string) string {
	data, _ := ioutil.ReadFile(filename)

	hash := md5.New()
	hash.Write([]byte(data))

	return hex.EncodeToString(hash.Sum(nil))
}

func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}

func Sha1_file(filename string) string {
	data, _ := ioutil.ReadFile(filename)

	hash := sha1.New()
	hash.Write([]byte(data))

	return hex.EncodeToString(hash.Sum(nil))
}

func Base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64_decode(str string) string {
	if len(str)%4 == 2 {
		str += "=="
	}

	if len(str)%4 == 3 {
		str += "="
	}

	en_byt, _ := base64.StdEncoding.DecodeString(str)
	return string(en_byt)
}

func Addslashes(str string) string {
	var buffer bytes.Buffer

	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buffer.WriteRune('\\')
		}

		buffer.WriteRune(char)
	}

	return buffer.String()
}

func Htmlentities(str string) string {
	return html.EscapeString(str)
}

func Html_entity_decode(str string) string {
	return html.UnescapeString(str)
}

func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func Url_encode(str string) string {
	return url.QueryEscape(str)
}

func Url_decode(str string) string {
	en_str, _ := url.QueryUnescape(str)
	return en_str
}

func Http_build_query(queryData url.Values) string {
	return queryData.Encode()
}

func Uniqid(prefix ...string) string {
	pre := ""
	if len(prefix) > 0 {
		pre = prefix[0]
	}

	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", pre, now.Unix(), now.UnixNano()%0x100000)
}
