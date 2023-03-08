package main

func main() {

}

/*
package main

import (
    "fmt"
    "image"
    "image/jpeg"
    "os"

    "github.com/dsoprea/go-exif/v3"
    "github.com/dsoprea/go-exif/v3/common"
)

func main() {
    // EXIFデータを取得したい画像ファイルを開く
    file, err := os.Open("sample.jpg")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // EXIFデータをパースする
    exifData, err := exif.SearchAndExtractExif(file)
    if err != nil {
        panic(err)
    }

    // サムネイル画像のバイト列を取得する
    thumbnail, err := exifData.Thumbnail()
    if err != nil {
        panic(err)
    }

    // サムネイル画像のバイト列から、image.Imageを生成する
    thumbnailImage, _, err := image.Decode(bytes.NewReader(thumbnail))
    if err != nil {
        panic(err)
    }

    // サムネイル画像を保存する
    out, err := os.Create("thumbnail.jpg")
    if err != nil {
        panic(err)
    }
    defer out.Close()

    // JPEGエンコーダーを使って、サムネイル画像を保存する
    jpeg.Encode(out, thumbnailImage, &jpeg.Options{Quality: 100})
    fmt.Println("Thumbnail saved!")
}

*/
