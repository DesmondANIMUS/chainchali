package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"

	"github.com/DesmondANIMUS/chainchali/chalipackages/blockchain"

	"github.com/DesmondANIMUS/chainchali/chali"
	"github.com/DesmondANIMUS/chainchali/chalipackages/chalimodel"
	"google.golang.org/grpc"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

type server struct {
	BlockChain *blockchain.Blockchain
}

func main() {
	lis, err := net.Listen("tcp", chalimodel.Port)
	handleErr(err)

	s := grpc.NewServer()
	registerServices(s)

	fmt.Println("Server listening at ", chalimodel.Port)
	handleErr(s.Serve(lis))
}

func registerServices(s *grpc.Server) {
	chali.RegisterBlockchainServer(s, server{
		BlockChain: blockchain.NewBlockChain(),
	})
}

func (s server) AddBlock(ctx context.Context, in *chali.AddBlockRequest) (*chali.AddBlockResponse, error) {
	block := s.BlockChain.AddBlock(in.Data)
	return &chali.AddBlockResponse{Hash: block.Hash}, nil
}

func (s server) GetBlockchain(ctx context.Context, in *chali.GetBlockchainRequest) (*chali.GetBlockchainResponse, error) {
	resp := &chali.GetBlockchainResponse{}
	for _, b := range s.BlockChain.Blocks {
		resp.Blocks = append(resp.Blocks, &chali.Block{
			Data:          b.Data,
			Hash:          b.Hash,
			PrevBlockHash: b.PrevBlockHash,
		})
	}
	return resp, nil
}
