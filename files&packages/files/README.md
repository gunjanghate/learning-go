# Files

This folder demonstrates how to work with files and directories in Go using the `os` and `bufio` packages.

## What's Inside

- `main.go`: Examples of
  - Opening an existing file and reading its metadata
  - Reading file contents using a fixed-size buffer and `os.ReadFile`
  - Creating and writing to new files
  - Copying data from one file to another using buffered streaming
  - Deleting files
- Focus: Learn the basics of file I/O operations and safe patterns around opening, reading, writing, copying, and deleting files.

## Key Concepts

### Opening and Inspecting a File

```go
f, err := os.Open("./ex.txt")
if err != nil {
    panic(err)
}
fileInfo, err := f.Stat()
```

- `os.Open` opens a file for reading.
- `Stat()` returns metadata such as name, size, modification time, and whether it is a directory.
- Always remember to `defer f.Close()` to release resources.

### Reading File Contents

```go
buf := make([]byte, fileInfo.Size())
_, err = f.Read(buf)
fmt.Println(string(buf))

data, err := os.ReadFile("./ex.txt")
fmt.Println(string(data))
```

- Manual read using a byte slice sized to `fileInfo.Size()`.
- `os.ReadFile` is a convenience function that reads the entire file into memory (good for small/medium files).

### Creating and Writing to Files

```go
file, err := os.Create("ex2.txt")
bytes := []byte("Hello gggg bye bye")
_, err = file.Write(bytes)
```

- `os.Create` creates (or truncates) a file.
- `Write` writes raw bytes. Alternatively, `os.WriteFile` can write a full buffer in one call.

### Buffered Streaming Copy (Reader/Writer)

```go
src, err := os.Open("ex.txt")
dest, err := os.Create("ex3.txt")

reader := bufio.NewReader(src)
writer := bufio.NewWriter(dest)

for {
    b, err := reader.ReadByte()
    if err != nil {
        if err.Error() == "EOF" {
            fmt.Println("End of file reached")
            break
        }
        panic(err)
    }

    if err := writer.WriteByte(b); err != nil {
        panic(err)
    }
}
writer.Flush()
```

- Demonstrates copying a file byte by byte using buffered I/O.
- `Flush()` ensures all buffered data is written to disk.

### Deleting Files

```go
err = os.Remove("ee2.txt")
if err != nil {
    panic(err)
}
```

- `os.Remove` deletes a file.

## Run The Example

From PowerShell:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\files&packages\files"
go run main.go
```

Make sure `ex.txt` exists in this folder, since the example expects that file.

## Notes

- Always close files you open (use `defer` right after opening).
- `os.ReadFile` and `os.WriteFile` are convenient but not ideal for very large files; streaming (with `bufio`) is more memory-efficient.
- Use clear file names like `ex.txt`, `ex2.txt`, `ex3.txt` for experimentation, but use meaningful names in real projects.
