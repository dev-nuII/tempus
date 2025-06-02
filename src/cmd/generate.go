package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var letterBytes = []byte(letters)
var lettersLen = len(letterBytes)

const THRESHOLD = 250_000

func GenerateTokens() {
	if TokenLength != 59 && TokenLength != 70 {
		fmt.Println("Invalid token length. Must be 59 or 70.")
		os.Exit(1)
	}

	var filename string
	if GeneratePath != "" {
		filename = GeneratePath
	} else {
		now := time.Now().Format("2006-01-02_15-04-05")
		filename = filepath.Join("tokens", now+".txt")
	}

	dir := filepath.Dir(filename)
	os.MkdirAll(dir, os.ModePerm)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	start := time.Now()

	if Count < THRESHOLD {
		generateTokensSmallCount(file)
	} else {
		generateTokensLargeCount(file)
	}

	elapsed := time.Since(start)
	fmt.Printf("Generated %d tokens in %s in %s\n", Count, filename, elapsed)
}

func generateTokensSmallCount(file *os.File) {
	writer := bufio.NewWriterSize(file, 512*1024)
	defer writer.Flush()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tokenLen := int(TokenLength)
	tokenSize := tokenLen + 1
	batchSize := 10_000
	if int(Count) < batchSize {
		batchSize = int(Count)
	}

	batchBuffer := make([]byte, batchSize*tokenSize)

	tokensLeft := int(Count)
	for tokensLeft > 0 {
		n := batchSize
		if tokensLeft < batchSize {
			n = tokensLeft
		}

		pos := 0
		for i := 0; i < n; i++ {
			fillRandomBytesFast(batchBuffer[pos:pos+tokenLen], r)
			pos += tokenLen
			batchBuffer[pos] = '\n'
			pos++
		}

		_, err := writer.Write(batchBuffer[:pos])
		if err != nil {
			fmt.Println("Error writing tokens:", err)
			os.Exit(1)
		}

		tokensLeft -= n
	}
}

func generateTokensLargeCount(file *os.File) {
	totalTokens := int(Count)
	numWorkers := runtime.GOMAXPROCS(0) * 2
	chunkSize := 100_000

	type job struct{ count int }
	jobs := make(chan job, numWorkers*2)
	chunks := make(chan []byte, numWorkers*2)

	var wg sync.WaitGroup

	var bufPool = sync.Pool{
		New: func() any {
			size := chunkSize * (int(TokenLength) + 1)
			return make([]byte, size)
		},
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(workerID)*1_000_000))

			for job := range jobs {
				buf := bufPool.Get().([]byte)

				if len(buf) < job.count*(int(TokenLength)+1) {
					buf = make([]byte, job.count*(int(TokenLength)+1))
				}

				pos := 0
				for i := 0; i < job.count; i++ {
					fillRandomBytes(buf[pos:pos+int(TokenLength)], r)
					pos += int(TokenLength)
					buf[pos] = '\n'
					pos++
				}

				chunks <- buf[:pos]
			}
		}(i)
	}

	go func() {
		remaining := totalTokens
		for remaining > 0 {
			thisChunk := chunkSize
			if remaining < chunkSize {
				thisChunk = remaining
			}
			jobs <- job{count: thisChunk}
			remaining -= thisChunk
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(chunks)
	}()

	writer := bufio.NewWriterSize(file, 4*1024*1024)
	defer writer.Flush()

	for chunk := range chunks {
		_, err := writer.Write(chunk)
		if err != nil {
			fmt.Println("Error writing tokens:", err)
			os.Exit(1)
		}

		bufPool.Put(chunk[:cap(chunk)])
	}
}

func fillRandomBytes(buf []byte, r *rand.Rand) {
	length := len(buf)
	i := 0
	for ; i < length-3; i += 4 {
		randomBits := r.Uint32()
		buf[i] = letterBytes[randomBits%uint32(lettersLen)]
		buf[i+1] = letterBytes[(randomBits>>8)%uint32(lettersLen)]
		buf[i+2] = letterBytes[(randomBits>>16)%uint32(lettersLen)]
		buf[i+3] = letterBytes[(randomBits>>24)%uint32(lettersLen)]
	}
	for ; i < length; i++ {
		buf[i] = letterBytes[r.Intn(lettersLen)]
	}
}

func fillRandomBytesFast(buf []byte, r *rand.Rand) {
	length := len(buf)

	i := 0
	for ; i < length-7; i += 8 {
		randomBits := r.Uint64()

		buf[i] = letterBytes[randomBits%uint64(lettersLen)]
		buf[i+1] = letterBytes[(randomBits>>8)%uint64(lettersLen)]
		buf[i+2] = letterBytes[(randomBits>>16)%uint64(lettersLen)]
		buf[i+3] = letterBytes[(randomBits>>24)%uint64(lettersLen)]
		buf[i+4] = letterBytes[(randomBits>>32)%uint64(lettersLen)]
		buf[i+5] = letterBytes[(randomBits>>40)%uint64(lettersLen)]
		buf[i+6] = letterBytes[(randomBits>>48)%uint64(lettersLen)]
		buf[i+7] = letterBytes[(randomBits>>56)%uint64(lettersLen)]
	}

	for ; i < length; i++ {
		buf[i] = letterBytes[r.Intn(lettersLen)]
	}
}

