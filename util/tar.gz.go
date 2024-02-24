package util

import (
	"archive/tar"
	"compress/gzip"
	"goZipper/log"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func TarGz() error {
	// 源文件夹路径
	srcFolder := "src"

	// 创建目标压缩文件
	targetFile, err := os.Create("src.tar.gz")
	if err != nil {
		log.Println("err when create file:", err)
		return err
	}
	defer targetFile.Close()

	// 创建 gzip 写入器
	gzipWriter := gzip.NewWriter(targetFile)
	defer gzipWriter.Close()

	// 创建 tar 写入器
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// 遍历源文件夹下的所有文件和子文件夹
	err = filepath.Walk(srcFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("err when walk file:", err)
			return err
		}

		// 忽略目录本身
		if path == srcFolder {
			log.Println("ignore src dir")
			return nil
		}

		// 创建 tar 记录
		tarHeader, err := tar.FileInfoHeader(info, "")
		if err != nil {
			log.Println("err when create tar header:", err)
			return err
		}

		// 更新 tar 记录的名称
		relPath, err := filepath.Rel(srcFolder, path)
		if err != nil {
			log.Println("err when get relative path:", err)
			return err
		}
		// 设置 tar 记录的名称
		tarHeader.Name = strings.Replace(relPath, "src\\", "", 1)
		tarHeader.Name = strings.Replace(tarHeader.Name, "\\", "/", -1)

		// 写入 tar 记录
		if err := tarWriter.WriteHeader(tarHeader); err != nil {
			log.Println("err when write tar header:", err)
			return err
		}

		// 如果是文件，则将文件内容写入 tar 文件
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				log.Println("err when open file:", err)
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tarWriter, file); err != nil {
				log.Println("err when copy file:", err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Println("err when walk file:", err)
		return err
	}
	return err
}
