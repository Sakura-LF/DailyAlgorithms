package StringsBenchMark

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestPlusConcat(t *testing.T) {
	s1 := "FF14"
	s2 := "World"
	s3 := "FF15"
	// go 底层会分配一个内存空间，然后把 s1 和 s2 的内容拷贝进去
	newString := s1 + s2
	// s1 拷贝了两次, s2 也是拷贝了两次, s3 拷贝了一次
	s4 := newString + s3
	fmt.Println(s3)
	fmt.Println(s4)
}

func TestSprintfConcat(t *testing.T) {
	userName := "Sakura"
	website := "www.sakurasss.top"
	email := "1808479176@qq.com"
	newString := fmt.Sprintf("用户名:%s\n网站:%s\n邮箱:%s", userName, website, email)
	fmt.Println(newString)
}

func TestBuilderConcat(t *testing.T) {
	var builder strings.Builder
	// builder也可以采用 Grow 的方式提前分配内存,以减少内存的分配次数
	//builder.Grow(1024)
	builder.WriteString("FF14")
	builder.WriteString("Sakura")

}

func TestBufferConcat(t *testing.T) {
	var buf bytes.Buffer
	// 向 buffer 中写入数据
	buf.WriteString("FF14")
	buf.WriteString("World")
	// buf.Reset() 用于清空 buffer,之后就读不到了

	// 从 buffer 中读取数据
	readDate := make([]byte, 4)
	n, err := buf.Read(readDate)
	if err != nil {
		log.Println(err)
		return
	}
	// 输出读取的数据
	fmt.Println(string(readDate[:n]))

	s := buf.String() // 获取 buffer 中的数据
	fmt.Println(s)    // 因为 FF14 已经被读取, 所以这里只会打印 World
}

//func BenchmarkPlusConcat2(b *testing.B) {
//	baseSlice := []string{"FF14", "World", "FF15"}
//	for i := 0; i < b.N; i++ {
//		strings.Join(baseSlice, "Sakura")
//	}
//}

func TestPlusConcat2(t *testing.T) {
	testSlice := []string{"FF14", "Sakura"}
	join := strings.Join(testSlice, " ")
	fmt.Println(join)
}
