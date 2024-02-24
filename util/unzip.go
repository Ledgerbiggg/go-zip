package util

import (
	"archive/zip"
	"goZipper/log"
	"io"
	"os"
)

// Unzip 解压
func Unzip(src string, dest string) error {
	// 打开 ZIP 文件
	zipFile, err := zip.OpenReader(src)
	if err != nil {
		log.Println("cant open zip:", err)
		return err
	}
	defer zipFile.Close()

	err = os.Mkdir(dest, os.ModePerm)
	if err != nil {
		//log.Println(dest+"dir is exist", err)
	}

	// 遍历 ZIP 文件中的文件和文件夹
	for _, file := range zipFile.File {
		// 打开 ZIP 文件中的文件
		fileReader, err := file.Open()
		if err != nil {
			log.Println("unzipping error:", err)
			return err
		}
		defer fileReader.Close()

		if file.FileInfo().IsDir() {
			err := os.Mkdir(dest+"\\"+file.Name, os.ModePerm)
			if err != nil {
				//log.Println("cant create dir:", err)
			}
			continue
		}

		// 创建目标文件
		targetFile, err := os.Create(dest + "\\" + file.Name)
		if err != nil {
			log.Println("cant create file:", err)
			return err
		}
		defer targetFile.Close()

		// 将 ZIP 文件中的文件内容复制到目标文件中
		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			log.Println("cant copy:", err)
			return err
		}

		log.Println("success unzip file:", file.Name)
	}

	log.Println("success unzip file:", src)
	return nil
}
