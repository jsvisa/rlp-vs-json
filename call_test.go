package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

// callLog is the result of LOG opCode
type callLog struct {
	Address  common.Address `json:"address"`
	Topics   []common.Hash  `json:"topics"`
	Data     hexutil.Bytes  `json:"data"`
	Position hexutil.Uint   `json:"position"`
}

// callFrame is the result of a callTracer run.
type callFrame struct {
	Type         string          `json:"type"`
	From         common.Address  `json:"from"`
	Gas          *hexutil.Uint64 `json:"gas"`
	GasUsed      *hexutil.Uint64 `json:"gasUsed"`
	To           *common.Address `json:"to,omitempty" rlp:"optional"`
	Input        hexutil.Bytes   `json:"input" rlp:"optional"`
	Output       hexutil.Bytes   `json:"output,omitempty" rlp:"optional"`
	Error        string          `json:"error,omitempty" rlp:"optional"`
	RevertReason string          `json:"revertReason,omitempty" rlp:"optional"`
	Calls        []callFrame     `json:"calls,omitempty" rlp:"optional"`
	Logs         []callLog       `json:"logs,omitempty" rlp:"optional"`
	Value        *hexutil.Big    `json:"value,omitempty" rlp:"optional"`
}

func loadJSONData() ([]callFrame, error) {
	file, err := os.Open("calls.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []callFrame
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	return data, err
}

func BenchmarkJSONEncoding(b *testing.B) {
	data, err := loadJSONData()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRLPEncoding(b *testing.B) {
	data, err := loadJSONData()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := rlp.EncodeToBytes(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSONDecoding(b *testing.B) {
	data, err := loadJSONData()
	if err != nil {
		b.Fatal(err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decodedData []callFrame
		err := json.Unmarshal(jsonData, &decodedData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRLPDecoding(b *testing.B) {
	data, err := loadJSONData()
	if err != nil {
		b.Fatal(err)
	}
	rlpData, err := rlp.EncodeToBytes(data)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decodedData []callFrame
		err := rlp.DecodeBytes(rlpData, &decodedData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStorageComparison(b *testing.B) {
	data, err := loadJSONData()
	if err != nil {
		b.Fatal(err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		b.Fatal(err)
	}

	rlpData, err := rlp.EncodeToBytes(data)
	if err != nil {
		b.Fatal(err)
	}

	b.Logf("JSON size: %d bytes", len(jsonData))
	b.Logf("RLP size: %d bytes", len(rlpData))
}
