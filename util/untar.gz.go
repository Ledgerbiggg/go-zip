package util

import (
	"archive/tar"
	"compress/gzip"
	"goZipper/log"
	"io"
	"os"
	"path/filepath"
)

func UnTarGz(src string, dest string) error {
	// 打开要解压缩的文件
	file, err := os.Open(src)
	if err != nil {
		log.Println("open file error:", err)
		return err
	}
	defer file.Close()

	err = os.Mkdir(dest, os.ModePerm)
	if err != nil {
		//log.Println(dest+"dir is exist", err)
	}

	// 创建 gzip 读取器
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		log.Println("read gzip error:", err)
		return err
	}
	defer gzipReader.Close()

	// 创建 tar 读取器
	tarReader := tar.NewReader(gzipReader)

	// 解压缩文件
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("read tar error:", err)
			return err
		}

		targetPath := filepath.Join("."+"\\"+dest, header.Name)

		// 根据文件类型执行不同的操作
		switch header.Typeflag {
		case tar.TypeDir:
			// 如果是目录，创建目录
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				log.Println("mkdir error:", err)
				return err
			}
		case tar.TypeReg:
			// 如果是文件，创建文件并写入内容
			fileToWrite, err := os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				log.Println("open file error:", err)
				return err
			}
			defer fileToWrite.Close()

			if _, err := io.Copy(fileToWrite, tarReader); err != nil {
				log.Println("copy error:", err)
				return err
			} else {
				log.Println("success unzip file:", header.Name)
			}
		}
	}
	return nil
}
