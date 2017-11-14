package gothaiwordcut

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPureThaiCut(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict()
	result := segmenter.Segment("ทดสอบการตัดคำภาษาไทย")
	assert.Equal(t, []string{"ทดสอบ", "การ", "ตัด", "คำ", "ภาษาไทย"}, result)
}
