# go-thaiwordcut - Thai word segmentation in Golang

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/12dbca6bf1c6471aacd46e4d829cc14d)](https://www.codacy.com/app/narongdejsrn/go-thaiwordcut?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=narongdejsrn/go-thaiwordcut&amp;utm_campaign=Badge_Grade)
---
A simple Thai word segmentation written in Golang, based on Maximum Matching algorithm by [S. Manabu](http://www.aclweb.org/anthology/E14-4016)
. Uses Lexitron (by [NECTEC](http://www.sansarn.com/lexto/license-lexitron.php)) dictionary as default

## Installation

```bash
go get github.com/narongdejsrn/go-thaiwordcut
```

## Usage
```golang
import "github.com/narongdejsrn/go-thaiwordcut"

segmenter := gothaiwordcut.Wordcut()
segmenter.LoadDefaultDict()
result := segmenter.Segment("ทดสอบการตัดคำภาษาไทย")
```