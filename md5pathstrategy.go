package md5_path_strategy

import (
	"crypto/md5"
	"fmt"
	"path"
	"strings"
)

type Md5PathStrategy struct{}

func (s Md5PathStrategy) GeneratePath(data []byte, ext string) (filePath string, fileNameWithExt string) {
	hash := fmt.Sprintf("%x", md5.Sum(data))
	filePath = path.Join(hash[:2], hash[2:4], hash[4:6])
	fileNameWithExt = strings.Join([]string{hash, ext}, ".")
	return filePath, fileNameWithExt
}
