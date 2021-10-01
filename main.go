package main

import (
	"fmt"
	"main.go/Controllers"
)

var ConnectionMongoDB string = Controllers.GetConfig().ConnectionMongoDB[0] //"connection string into your application code"
var DataBaseBlockchain string = Controllers.GetConfig().DataBase[0]         //blockchain

var CollectionAdresses string = Controllers.GetConfig().Collection[0]          // "Adresses"
var CollectionBlockchain string = Controllers.GetConfig().Collection[1]        // "blockchain"
var CollectionEnderecos string = Controllers.GetConfig().Collection[2]         // "enderecos"
var CollectionTeste string = Controllers.GetConfig().Collection[3]             // "teste"
var CollectionTesteMultiAdress string = Controllers.GetConfig().Collection[4]  // testeMultiAdress
var CollectionMapeandoEnderecos string = Controllers.GetConfig().Collection[5] //MapeandoEnderecos
var CollectionTesteMapeando string = Controllers.GetConfig().Collection[6]     // TesteMapeando

var UrlAPI string = Controllers.GetConfig().UrlAPI[0] // "https://blockchain.info"

var RawAddr string = Controllers.GetConfig().RawAddr
var MultiAddr string = Controllers.GetConfig().MultiAddr

var LogBlockchain string = Controllers.GetConfig().FileLog[0]             // "LogBlockchain.txt"
var LogIndiceEndereco string = Controllers.GetConfig().FileLog[1]         // "LogIndiceEndereco.txt"
var LogEndereco string = Controllers.GetConfig().FileLog[2]               // "LogEndereco.txt"
var LogEnderecosSemDados string = Controllers.GetConfig().FileLog[3]      // "LogEnderecosSemDados.txt"
var LogIndiceMultiEndereco string = Controllers.GetConfig().FileLog[4]    // LogIndiceMultiEndereco.txt
var LogMultiEnderecosSemDados string = Controllers.GetConfig().FileLog[5] //LogMultiEnderecosSemDados.txt

func main() {

	// Buscar um unico endereco
	y := Controllers.GetMapAdressId(ConnectionMongoDB, DataBaseBlockchain, CollectionTesteMapeando, "_id", "6153a58d3700e70e40f8177a")
	fmt.Println(y)

	//Controllers.SalvaListaEnderecos(ConnectionMongoDB, DataBaseBlockchain, CollectionEnderecos, UrlAPI, RawAddr, CollectionAdresses, LogEnderecosSemDados,
	//	LogIndiceEndereco)

	// Pega os valores da Collection
	// retorna uma lista mapeada
	//x := Controllers.MapeandoEndereco(ConnectionMongoDB, DataBaseBlockchain, CollectionTeste)
	//// salva no mongoDB
	//Controllers.SalvarMapeamentoTransacaoMongoDB(x, ConnectionMongoDB, DataBaseBlockchain, CollectionTesteMapeando)

}
