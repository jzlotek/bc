package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jzlotek/bc/block"
)

type Response struct {
	Data interface{} `json:"data"`
}

func getHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{"Hello world!"})
}

type hashNoncePair struct {
	Hash  []byte
	Nonce uint
}

func doHashRange(
	block block.Block, // will contain all data (nonce gets changed in this functions)
	startingString string,
	maxNonce uint,
	offset uint,
	workers uint,
	hashChan chan hashNoncePair, // return channel
	lock chan bool, // locks the hash channel
	kill chan bool, // if retrieves a kill signal, return immediately
) {
	fmt.Printf("Starting worker: %d\n", offset)
	var hash []byte
	var err error
	nonce := uint(0) + offset

	for maxNonce == 0 || nonce < maxNonce {
		block.Nonce = nonce
		hash, err = block.Hash()
		select {
		case <-kill:
			fmt.Printf("Killing worker: %d\n", offset)
			return
		default:
			if err != nil {
				continue
			}
			if strings.HasPrefix(string(hash), startingString) {
				select {
				case <-lock:
					fmt.Printf("Worker: %d found: %s\n", offset, string(hash))
					hashChan <- hashNoncePair{hash, nonce}
					return
				default:
					fmt.Printf("Killing worker: %d\n", offset)
					return
				}
			}
			nonce += workers
		}
	}
}

func doHash(block block.Block, startingString string, maxNonce uint) (string, uint, time.Duration, error) {
	workers := uint(16)
	start := time.Now().UTC()

	hashChan := make(chan hashNoncePair)
	lock := make(chan bool, 1)
	lock <- true
	kill := make(chan bool, workers-1)

	wg := &sync.WaitGroup{}
	wg.Add(int(workers))

	// dispatch 16 workers
	for i := uint(0); i < workers; i++ {
		go func(i uint) {
			defer wg.Done()
			doHashRange(block, startingString, maxNonce, i, workers, hashChan, lock, kill)
		}(i)
	}

	hash := <-hashChan

	// send kill signals to workers
	for i := 0; uint(i) < workers-1; i++ {
		kill <- true
	}

	duration := time.Now().Sub(start)
	wg.Wait()
	return string(hash.Hash), hash.Nonce, duration, nil
}

func doSingleHash(ctx *gin.Context) {
	type req struct {
		Block       block.Block `json:"block" binding:"required"`
		StartString string      `json:"start" binding:"required"`
		MaxNonce    uint        `json:"max"`
	}

	request := &req{}
	if err := ctx.ShouldBindBodyWith(request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{err.Error()})
		return
	}

	hash, nonce, time, err := doHash(request.Block, request.StartString, request.MaxNonce)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{map[string]interface{}{
		"hash":  hash,
		"nonce": nonce,
		"time":  time,
	}})
}

func main() {
	var addr *string
	addr = flag.String("bind", ":8000", "address to bind to")
	flag.Parse()

	router := gin.Default()
	router.GET("/", getHome)
	router.POST("/single", doSingleHash)

	srv := &http.Server{
		Addr:    *addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
