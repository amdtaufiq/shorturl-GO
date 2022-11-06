package service

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/amdtaufiq/sorturl-GO/app/dto"
	"github.com/itchyny/base58-go"
)

type IGeneratorService interface {
	GenerateSortUrl(request dto.CreateGeratorRequest) (response dto.CreateGeratorResponse)
}

type generatorService struct {
}

func GeneratorService() *generatorService {
	return &generatorService{}
}

func (s *generatorService) GenerateSortUrl(request dto.CreateGeratorRequest) (response dto.CreateGeratorResponse) {
	urlHashBytes := sha256Of(request.LongURL)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	response.SortURL = finalString[:6]
	return
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
