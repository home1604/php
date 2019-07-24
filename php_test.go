package php

import (
	"net/url"
	"testing"
)

func TestTime(t *testing.T) {
	tt := Time()
	t.Log(tt)
}

func TestDate(t *testing.T) {
	tt := Time()
	dd := Date("2006#01#02", tt)
	t.Log(dd)
}

func TestStrtotime(t *testing.T) {
	t.Log(Strtotime("hello world"))
	t.Log(Strtotime("2019-07-22 23:23:23"))
}

func TestSleep(t *testing.T) {
	t.Log(Time())
	Sleep(5)
	t.Log(Time())
}

func TestUSleep(t *testing.T) {
	t.Log(Time())
	Usleep(500000)
	t.Log(Time())
}

//string ----

func TestStrlen(t *testing.T) {
	t.Log(Strlen("hello world"))
	t.Log(Strlen("您好世界！hello world"))
}

func TestMb_strlen(t *testing.T) {
	t.Log(Mb_strlen("您好世界！hello world"))
}

func TestStr_replace(t *testing.T) {
	t.Log(Str_replace("hello", "HELLO", "hello world hello!", 1))
	t.Log(Str_replace("hello", "HELLO", "hello world hello!"))
}

func TestStrtolower(t *testing.T) {
	t.Log(Strtolower("HELLO Wold!"))
}

func TestStrtoupper(t *testing.T) {
	t.Log(Strtoupper("hello world!"))
}

func TestUcwords(t *testing.T) {
	t.Log(Ucwords("hell world, how are you?"))
}

func TestSubstr(t *testing.T) {
	t.Log(Substr("hello world!", 0, 2))
}

func TestStrrev(t *testing.T) {
	t.Log(Strrev("hello world!啊a"))
}

func TestTrim(t *testing.T) {
	t.Log(Trim(" 世界你好 	 " +
		"	\n  "))
	t.Log(Trim("世界\\r\\n\\t"))
}

func TestLTrim(t *testing.T) {
	t.Log(Ltrim(" 世界你好 	 " +
		"	\n  "))
	t.Log(Ltrim("	aa aa 世界\\r\\n\\t"))
}

func TestRTrim(t *testing.T) {
	t.Log(Rtrim(" 世界你好 	 " +
		"	\n  "))
	t.Log(Rtrim("	hello world 世界\r\\n\\t"))
}

func TestExplode(t *testing.T) {
	t.Log(Explode(",", "a,b,c,d,e,f"))
}

func TestImplode(t *testing.T) {
	t.Log(Implode(",", []string{"hello", "world", "!"}))
}

func TestChr(t *testing.T) {
	t.Log(Chr(52))
	t.Log(Chr(88))
	t.Log(Chr(67))
}

func TestOrd(t *testing.T) {
	t.Log(Ord("A"))
}

func TestMd5(t *testing.T) {
	t.Log(Md5("123456"))
}

func TestMd5_file(t *testing.T) {
	t.Log(Md5_file("php_test.go"))
}

func TestSha1(t *testing.T) {
	t.Log(Sha1("123456"))
}

func TestSha1_file(t *testing.T) {
	t.Log(Sha1_file("php_test.go"))
}

func TestBase64_encode(t *testing.T) {
	t.Log(Base64_encode("啊啊嗯嗯版本哦哦"))
}

func TestBase64_decode(t *testing.T) {
	en_str := Base64_encode("啊啊嗯嗯版本哦哦a  e")
	t.Log(en_str)
	t.Log(Base64_decode(en_str))
}

func TestAddslashes(t *testing.T) {
	t.Log(Addslashes("啊啊\"bbb'cccc"))
}

func TestHtmlentities(t *testing.T) {
	t.Log(Htmlentities("<b>aaa</b>"))
}

func TestHTMLEntityDecode(t *testing.T) {
	str := Htmlentities("<b>aaa</b>")
	t.Log(Html_entity_decode(str))
}

func TestCrc32(t *testing.T) {
	t.Log(Crc32("hello world"))
}

func TestUrl_encode(t *testing.T) {
	t.Log(Url_encode("https://www.baidu.com/s?wd=php&rsv_spt=1&rsv_iqid=0x9c452cfc000a21bf&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=4&rsv_sug1=4&rsv_sug7=101&rsv_sug2=0&inputT=1199&rsv_sug4=2130"))
}

func TestUrl_decode(t *testing.T) {
	en_str := Url_encode("https://www.baidu.com/s?wd=php&rsv_spt=1&rsv_iqid=0x9c452cfc000a21bf&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=4&rsv_sug1=4&rsv_sug7=101&rsv_sug2=0&inputT=1199&rsv_sug4=2130")
	t.Log(en_str)
	t.Log(Url_decode(en_str))
}

func TestHttp_build_query(t *testing.T) {
	val := url.Values{}

	val.Add("query", "baidu")
	val.Add("page", "2")

	t.Log(Http_build_query(val))
}

func TestUniqid(t *testing.T) {
	t.Log(Uniqid())
	t.Log(Uniqid("uid_"))
}

func TestFile_exists(t *testing.T) {
	t.Log(File_exists("php.go"))
}

func TestIs_file(t *testing.T) {
	t.Error(Is_file("php.go"))
	t.Error(Is_file("/Users/wuzhijian"))
}

func TestIs_dir(t *testing.T) {
	t.Error(Is_dir("php.go"))
	t.Error(Is_dir("/Users/wuzhijian"))
}
