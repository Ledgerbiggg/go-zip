package util

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Zip 压缩
func Zip(dst string, src string) (err error) {
	// 创建准备写入的文件
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer zw.Close()

	// 去除src目录前缀
	var i = true

	// 下面来将文件写入 zw ，因为有可能会有很多个目录及文件，所以递归处理
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		// 处理遍历过程中的错误
		if errBack != nil {
			log.Println("filepath.Walk:", errBack)
			return errBack
		}

		// 使用 zip.FileInfoHeader 方法将文件信息转换为 zip 文件头
		fh, err := zip.FileInfoHeader(fi)
		// 将最外层的src去除掉
		fh.Name = strings.Replace(path, "src\\", "", 1)
		// 将\ 替换为 /
		fh.Name = strings.Replace(fh.Name, "\\", "/", -1)

		// 将src目录去除,但是只能去除一次
		if fi.IsDir() && fh.Name == "src" && i {
			i = false
			return nil
		}

		if err != nil {
			log.Println("fh error:", err)
			return
		}

		// 使用 zip.Writer.CreateHeader 方法创建一个新的文件并得到一个 io.Writer 对象
		w, err := zw.CreateHeader(fh)
		if err != nil {
			log.Println("w error:", err)
			return
		}

		// 文件夹不需要拷贝
		if fi.IsDir() {
			return nil
		}

		// 打开文件
		fr, err := os.Open(path)

		defer fr.Close()
		if err != nil {
			log.Println("fr error:", err)
			return
		}

		_, errBack = io.Copy(w, fr)
		log.Println("success zip file:", fh.Name)
		if errBack != nil {
			log.Println("written error:", errBack)
			return
		}
		return nil
	})
}
