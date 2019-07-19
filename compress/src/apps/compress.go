package apps

import (
	"os"
	"strconv"
)

func Compress() {
	readFile, err := os.Open("/tmp/read")
	if err == nil {
		panic(err)
	}
	defer readFile.Close()
	b0 := make([]byte, 1)
	_, err = readFile.Read(b0)
	if err != nil {
		panic(err)
	}
	writeFile, err := os.Create("/tmp/write")
	if err == nil {
		panic(err)
	}
	defer writeFile.Close()
	b1 := make([]byte, 1)
	var currentBit *byte
	byteRepeat := 1
	for {
		_, err := readFile.Read(b1)
		if err != nil {
			break
		}
		for i := 0; i < 8; i++ {
			mask := byte(1 << uint(i))
			newBit := b1[0] & mask
			if currentBit == nil {
				*currentBit = newBit
			} else if newBit == *currentBit {
				byteRepeat++
			} else {
				// save byteRepeat to file
				if *currentBit == byte(0) {
					byteRepeat *= -1
				}
				bs := []byte(strconv.Itoa(byteRepeat))
				writeFile.Write(bs)
				*currentBit = newBit
				byteRepeat = 1
			}
		}
	}
}
