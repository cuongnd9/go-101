package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func downloadImage(url string, filePath string, done chan bool) {
	// Tải ảnh từ URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		done <- false
		return
	}
	defer response.Body.Close()

	// Mở một file để lưu ảnh
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		done <- false
		return
	}
	defer file.Close()

	// Copy dữ liệu từ response.Body vào file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error saving image:", err)
		done <- false
		return
	}

	done <- true
}

func saveToDatabase(imageURL string, filePath string, db *sql.DB) {
	// Chèn thông tin ảnh vào cơ sở dữ liệu
	_, err := db.Exec("INSERT INTO images (url, file_path) VALUES (?, ?)", imageURL, filePath)
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return
	}
}

func main() {
	// Kết nối đến cơ sở dữ liệu
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// URL của ảnh cần tải về
	imageURL := "https://example.com/image.jpg"

	// Tạo một channel để đợi tất cả các goroutine hoàn thành
	done := make(chan bool)

	// Thực hiện tải ảnh và lưu vào đĩa trong một goroutine
	go downloadImage(imageURL, "image.jpg", done)

	// Chờ goroutine hoàn thành
	select {
	case success := <-done:
		if success {
			fmt.Println("Image downloaded successfully")
			// Sau khi tải xong ảnh, lưu thông tin vào cơ sở dữ liệu
			saveToDatabase(imageURL, "image.jpg", db)
		} else {
			fmt.Println("Failed to download image")
		}
	case <-time.After(30 * time.Second):
		fmt.Println("Timeout: Downloading image took too long")
	}
}
