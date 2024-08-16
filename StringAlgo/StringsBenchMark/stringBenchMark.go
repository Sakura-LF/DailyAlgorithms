package StringsBenchMark

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

const letterBytes = "SakurasssFF14246810GoFlutter"

// 生成随机字符串
func randomString(n int) string {
	// n 表示生成多长的字符串
	b := make([]byte, n)
	for i := range b {
		// 每次循环随机中间随机一个字符,构成字符串
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

func joinConcat(n int, str []string) string {
	newStr := ""
	for i := 0; i < n; i++ {
		newStr = strings.Join(str, "")
	}
	return newStr
}

func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
