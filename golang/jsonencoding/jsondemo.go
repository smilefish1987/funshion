package main

import (
	"encoding/json"
	"log"
	"os"
)

/*
 * Go内建的encoding/json包还提供Decoder和Encoder两个类型，用于支持JSON数据的流式读写，并提供NewDecoder()
 * 和NewEncoder()两个函数来便于具体实现：
 *  func NewDecoder(r io.Reader) *Decoder
 *  func NewEncoder(w io.Writer) *Encoder
 */

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}

		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		for k := range v {
			if k != "Title" {
				v[k] = nil
			}
		}

		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
