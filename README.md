## 📁 Go File Chunk Upload Server

This project is a simple file upload server built using [Gin](https://github.com/gin-gonic/gin) in Go.

---

### 🏗️ Features

* Upload files in chunks (resumable uploads)
* Handles duplicate filenames by renaming (`file.txt`, `file(1).txt`, etc.)
* Merges uploaded chunks into a complete file
* Simple HTML interface for uploading

---

### 📂 Folder Structure

```
.
├── main.go           # Core server code
├── templates        
    ├── index.html        # HTML file upload interface
├── uploads/          # Final merged files stored here
├── handlers/      
    ├── upload.go        # Handle uploading files in chunks
├── initializers/       
    ├── LoadEnvFile.go        # Load .env file
```

---

### ⚙️ How It Works

1. **Frontend (`index.html`)**

   * Splits selected file into chunks
   * Uploads each chunk via `/upload-chunk`

2. **Backend (`main.go`)**

   * Stores each chunk in a folder (named after FileID)
   * Once all chunks are uploaded, merges them into one final file
   * If the filename already exists, automatically renames the new file

---

### 🚀 Getting Started

#### 1. Clone the Repo

```bash
git clone https://github.com/Amirreza-Zeraati/go-file-service.git
cd go-file-service
```

#### 2. Install Dependencies

Make sure you have Go installed: [https://golang.org/dl/](https://golang.org/dl/)

```bash
go mod init file-service
go get github.com/gin-gonic/gin
```

#### 3. Run the Server

```bash
go run main.go
```

#### 4. Create .env file (if you want)

```bash
PORT=3000
```

#### 5. Open in Browser

Visit: [http://localhost:3000/](http://localhost:3000/)

---

### 📤 Upload Endpoint

**POST** `/upload-chunk`

**Form Data Fields:**

| Field         | Type   | Description                      |
| ------------- | ------ | -------------------------------- |
| `fileName`    | string | Original file name               |
| `chunkIndex`  | number | Current chunk number (0-based)   |
| `totalChunks` | number | Total number of chunks to expect |
| `chunk`       | file   | File chunk data (binary)         |

---

### 📌 Notes

* Final files are saved in the `uploads/` directory.
* Temporary chunk folders are also stored in `uploads/` but will remove after complete upload.
* File name conflicts are handled by appending a number to the new file.

---

### 🧪 TODO / Improvements

* [ ] Add progress bar to the frontend
* [ ] Add download route
* [ ] Handle interrupted uploads and resume from last chunk
* [ ] Add authentication and rate-limiting

---
