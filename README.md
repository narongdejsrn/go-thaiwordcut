# go-thaiwordcut - Thai word segmentation in Golang

A simple Thai word segmentation written in Golang, based on Maximum Matching algorithm by [S. Manabu](http://www.aclweb.org/anthology/E14-4016)
. Uses Lexitron (by [NECTEC](http://www.sansarn.com/lexto/license-lexitron.php)) dictionary as default

## Installation

```
go get github.com/narongdejsrn/go-thaiwordcut
```

## Usage
```
import "github.com/narongdejsrn/go-thaiwordcut"

segmenter := gothaiwordcut.Wordcut()
segmenter.LoadDefaultDict()
result := segmenter.Segment("ทดสอบการตัดคำภาษาไทย")
```