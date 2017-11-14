package gothaiwordcut

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func TestPureThaiCut(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict()
	result := segmenter.Segment("ทดสอบการตัดคำภาษาไทย")
	assert.Equal(t, []string{"ทดสอบ", "การ", "ตัด", "คำ", "ภาษาไทย"}, result)
}

func TestMixEnglishThaiCut(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict()
	result := segmenter.Segment("มาลองตัดคำปนภาษา English กันนะ Alright เพื่อน")
	assert.Equal(t, []string{"มา", "ลอง", "ตัด", "คำ", "ปน", "ภาษา", "English", "กัน", "นะ", "Alright", "เพื่อน"}, result)
}

func BenchmarkWordcut(b *testing.B) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict()
	dat, _ := ioutil.ReadFile("./dict/benchmark_text.txt")
	segmenter.Segment(string(dat))
}