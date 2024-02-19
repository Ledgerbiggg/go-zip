package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src string) error {
	// 打开 ZIP 文件
	zipFile, err := zip.OpenReader(src)
	if err != nil {
		fmt.Println("cant open zip:", err)
		return err
	}
	defer zipFile.Close()

	// 遍历 ZIP 文件中的文件和文件夹
	for _, file := range zipFile.File {
		// 打印当前文件名
		fmt.Println("unzipping:", file.Name)

		// 打开 ZIP 文件中的文件
		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("unzipping error:", err)
			return err
		}
		defer fileReader.Close()

		if file.FileInfo().IsDir() {
			err := os.Mkdir("dest\\"+file.Name, os.ModePerm)
			if err != nil {
				fmt.Println("cant create dir:", err)
			}
			continue
		}

		// 创建目标文件
		targetFile, err := os.Create("dest\\" + file.Name)
		if err != nil {
			fmt.Println("cant create file:", err)
			return err
		}
		defer targetFile.Close()

		// 将 ZIP 文件中的文件内容复制到目标文件中
		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			fmt.Println("cant copy:", err)
			return err
		}

		fmt.Println("success unzip file:", file.Name)
	}

	fmt.Println("success unzip file:", src)
	return nil
}

func Zip(dst string) (err error) {
	var src = "src"
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
			fmt.Println("filepath.Walk:", errBack)
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
			fmt.Println("fh error:", err)
			return
		}

		// 使用 zip.Writer.CreateHeader 方法创建一个新的文件并得到一个 io.Writer 对象
		w, err := zw.CreateHeader(fh)
		if err != nil {
			fmt.Println("w error:", err)
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
			fmt.Println("fr error:", err)
			return
		}

		_, errBack = io.Copy(w, fr)
		if errBack != nil {
			fmt.Println("written error:", errBack)
			return
		}
		return nil
	})
}
