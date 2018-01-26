
package main

import (
	"path/filepath"
	"os"
	"io"
	"io/ioutil"
	"archive/zip"
	"log"
	"archive/tar"
	"fmt"
)

// 解凍を行う関数
func Unzip(src, dest string) error {
	// ZIPファイルの中身を取得する
	r, err := zip.OpenReader(src)
	if err != nil { // 取得できなかったとき
		return err
	}
	defer r.Close() // ファイルクローズ

	// ファイルの中身を１つずつ処理
	for _, f := range r.File {
		_, err := openZipFile(*f)
		if err != nil {
			return err
		}
		if f.FileInfo().IsDir() {
			createDirectory(dest, f.Name, f.Mode())
		} else {
			writeFile(dest, f.Name, int64(f.UncompressedSize64), f.Mode())
		}
	}
	return nil
}

func openZipFile(f zip.File) (io.ReadCloser, error){
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return rc, nil
}

func Untar(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(file)
	for {
		f, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if f.FileInfo().IsDir() {
			createDirectory(dest, f.Name, os.FileMode(f.Mode))
		} else {
			writeFile(dest, f.Name, f.Size, os.FileMode(f.Mode))
		}
	}
	return nil
}

func createDirectory(dest, name string, mode os.FileMode) {
	path := filepath.Join(dest, name)
	os.MkdirAll(path, mode)
}

func writeFile(dest, name string, size int64, mode os.FileMode) error {
	buf := make([]byte, size)
	path := filepath.Join(dest, name)
	err := ioutil.WriteFile(path, buf, mode)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	var err error
	switch os.Args[2] {
	case "zip":
		err = Unzip(os.Args[1], "./")
	case "tar":
		err = Untar(os.Args[1], "./")
	default:
		usage()
	}
	if err != nil { // エラー
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println("Usage: [fileName] [zip|tar]")
	os.Exit(1)
}