# PHP functions for Golang
   golang使用的一些函数，这些函数很向php的函数
   
# Go func => PHP function

* Time => time
* Date => date
* Strtotime => strtotime
* Sleep => sleep
* Usleep => usleep
* Strlen => strlen
* Mb_strlen => mb_strlen
* Str_replace => str_replace
* Strtoupper => strtoupper
* Strtolower => strtolower
* Ucwords => ucwords
* Substr => substr
* Strrev => strrev
* Trim => trim
* Ltrim => ltrim
* Rtrim => rtimr
* Explode => explode
* Implode => implode
* Chr => char
* Ord => ord
* Md5 => md5
* Md5_file => md5_file
* Sha1 => sha1
* Sha1_file => sha1_file
* Base64_encode => base64_encode
* Base64_decode => base64_decode
* Addslashes => addslashes
* Htmlentities => htmlentities
* HtmlEntityDecode => html_entity_decode()
* Crc32 => crc32
* Url_encode => url_encode
* Url_decode => url_decode
* Http_build_query => http_build_query
* Uniqid => uniqid

# 使用方法
1. Go module

    `require github.com/zhijianwu/php latest`

    上面内容加入go.mod文件内

    `go mod vendor`
2. Go package

    `go get -v github.com/zhijianwu/php`
    
3. Example

```go
package main

import (
	"fmt"
	"github.com/zhijianwu/php"
)

func main()  {
	fmt.Println(php.Time())
	fmt.Println(php.Date("2006~~~01#02",php.Time()))
}

//console
1563257169
2019~~~07#16
```