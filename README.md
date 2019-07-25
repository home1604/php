# About
   Those functions are like PHP's functions.
   
# Mapping
   **Go func => PHP function**

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
* Html_entity_decode => html_entity_decode()
* Crc32 => crc32
* Url_encode => url_encode
* Url_decode => url_decode
* Http_build_query => http_build_query
* Uniqid => uniqid
* File_exists => file_exists
* Is_file => is_file
* Is_dir => is_dir
* Array_keys => array_keys
* Array_values => array_values
* Array_column => array_column

# Usage
1. Go module

    `require github.com/zhijianwu/php latest`

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