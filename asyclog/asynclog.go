package asynclog

import (
	"log"
	"os"
	"sync"
)

type Block struct {
	logs 	 []string
	rwlock   *sync.RWMutex
}

type AsyncLogger struct {
	mutex       sync.Mutex
	logBlocks   []Block
	currentSize int
	maxSize     int
	waitGroup   sync.WaitGroup
}

func NewAsyncLogger(blockSize, initialBlocks int) *AsyncLogger {
	logger := &AsyncLogger{
		logBlocks:   make([]Block, initialBlocks),
		currentSize: 0,
		maxSize:     blockSize * initialBlocks * 8 / 10,
	}

	logger.waitGroup.Add(1)
	go logger.logWorker()

	return logger
}

func (al *AsyncLogger) logWorker() {
	defer al.waitGroup.Done()

	file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)

	for {
		for _, block := range al.logBlocks {
			locker := block.rwlock.RLocker()
			locker.Lock()
			for _, msg := range block.logs {
				logger.Println(msg)
			}
		}
	}
}

func (al *AsyncLogger) Log(message string) {
	al.mutex.Lock()
	defer al.mutex.Unlock()

	for i := range al.logBlocks {
		if len(al.logBlocks) < len(al.logBlocks[i].logs) {
			al.logBlocks[i].logs = append(al.logBlocks[i].logs, message)
			al.currentSize++
			return
		}
	}

	if al.currentSize >= al.maxSize {
		al.maxSize++
		go al.expandBlocks()
	}

	for i := range al.logBlocks {
		if len(al.logBlocks) < len(al.logBlocks[i].logs) {
			al.logBlocks[i].logs = append(al.logBlocks[i].logs, message)
			al.currentSize++
			return
		}
	}
}

func (al *AsyncLogger) expandBlocks() {
	// 创建新块
	newBlock := Block{}
	al.logBlocks = append(al.logBlocks, newBlock)
	al.currentSize = 0
}
