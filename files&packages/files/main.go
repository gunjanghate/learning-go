package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./ex.txt")
	if err != nil {
		panic(err)
	}

	fileInfo, errr := f.Stat()
	if errr != nil {
		panic(errr)
	}

	// fmt.Println(fileInfo)
	fmt.Println("File Name:", fileInfo.Name())
	// read fileInfo
	// fmt.Println("Size in bytes:", fileInfo.Size())
	// fmt.Println("Last modified:", fileInfo.ModTime())
	// fmt.Println("Is Directory:", fileInfo.IsDir())

	defer f.Close()
	// read file content

	buf := make([]byte, fileInfo.Size())

	fmt.Println(buf) // [0 0 0 0 0 0 0 0]
	d, errr := f.Read(buf)
	if errr != nil {
		panic(errr)
	}
	fmt.Println("Bytes read:", d)
	fmt.Println("File content:")
	fmt.Println(buf)         // [72 101 108 108 111 32 103 103]
	fmt.Println(string(buf)) // Hello gg

	data, errrrrr := os.ReadFile("./ex.txt") // ReadFile is a utility function that reads the entire file content
	// not recommended for large files
	if errrrrr != nil {
		panic(errrrrr)
	}
	fmt.Println("File content using ReadFile:")
	fmt.Println(string(data))

	// read folders
	// dirEntries, err := os.ReadDir("./")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Directory entries:")
	// for _, entry := range dirEntries {
	// 	fmt.Println("Name:", entry.Name(), "IsDir:", entry.IsDir())
	// }

	// create a new file
	file, err := os.Create("ex2.txt")
	if err != nil {
		panic(err)
	}

	// i, err :=file.WriteString("Hi gg")
	// if err != nil {
	// 	panic(err)
	// } else{
	// 	fmt.Println("Bytes written:", i)
	// }
	// // replace file content
	bytes := []byte("Hello gggg bye bye")
	// err = os.WriteFile("ex2.txt", bytes, 0644)
	_, err = file.Write(bytes)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("File content replaced successfully")
	}
	defer file.Close()

	// read & write to another file usinf streaming fashion

	// reader
	src, err := os.Open("ex.txt")
	if err != nil {
		panic(err)
	}

	// writer dest
	dest, err := os.Create("ex3.txt")
	if err != nil {
		panic(err)
	}
	defer dest.Close()
	// fileInfoo, err := src.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// defer src.Close()

	// buff := make([] byte, fileInfoo.Size())
	// _, err = src.Read(buff)
	// if err != nil {
	// 	panic(err)
	// }

	// s := string(buff)
	// fmt.Println("Source file content:", s)

	
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dest)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("End of file reached")
				break
			} else {

				panic(err)
			}
		}
		fmt.Println("Read byte:", b)

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}

	}
	writer.Flush()
	fmt.Println("File copied successfully using streaming")	

	// delete
	err = os.Remove("ee2.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("File deleted successfully")
	}
}
