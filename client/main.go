package main

import (
	"flag"
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/DesmondANIMUS/chainchali/chali"

	"github.com/DesmondANIMUS/chainchali/chalipackages/chalimodel"
	"google.golang.org/grpc"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	conn, err := grpc.Dial(chalimodel.Port, grpc.WithInsecure())
	handleErr(err)
	defer conn.Close()

	client := chali.NewBlockchainClient(conn)

	if *addFlag {
		addBlock(client)
	}

	if *listFlag {
		listBlockchain(client)
	}
}

func addBlock(client chali.BlockchainClient) {
	block, err := client.AddBlock(context.Background(), &chali.AddBlockRequest{
		Data: "New Block created at: " + time.Now().String(),
	})
	handleErr(err)
	fmt.Println("New block added: ", block)
}

func listBlockchain(client chali.BlockchainClient) {
	bc, err := client.GetBlockchain(context.Background(), &chali.GetBlockchainRequest{})
	handleErr(err)
	fmt.Println("Entire Blockchain: ")
	for _, b := range bc.Blocks {
		fmt.Println(b)
	}
}
