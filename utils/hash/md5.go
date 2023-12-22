package hash

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

// Md5 计算 md5 值
func Md5(byteData []byte) string {
	hash := md5.New()
	hash.Write(byteData)
	hashByteData := hash.Sum(nil)
	return fmt.Sprintf("%x", hashByteData)
}

// FileMd5 计算上传文件的 md5 值
func FileMd5(file multipart.File) string {
	hash := md5.New()
	// 这里用到了 copy 方法，千万不要直接读取原 file 对象
	io.Copy(hash, file)
	hashByteData := hash.Sum(nil)
	return fmt.Sprintf("%x", hashByteData)
}
