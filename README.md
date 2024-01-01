# labGob

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Pinablink/tmdbGoTutorial?style=plastic)

Imagine você trabalhando em um sistema que precise obter alguns dados sensíveis em um certo momento do processamento. E que esses dados só serão úteis em um processo posterior, digamos que seja um processo temporal. Seja em uma fila ou em um cache, esses dados precisam permanecer ali até que seja solicitado a sua disponiblidade.

## A idéia
Esse laboratório irá demonstrar como trabalhar com serialização, compressão e output de um map de dados no Golang. Nessa solução proposta para estudo, teremos um map que será composto por duas structs. O valor serializado em bytes desse map na memória, ficou em 415 Bytes em nosso experimento. Após a execução do processo, foi reduzido o valor para 300 Bytes.

## Material de consulta desse laboratório

https://medium.com/@weberasantos/serializando-um-map-em-golang-aplicando-compress%C3%A3o-gzip-e-persistindo-em-um-arquivo-bin%C3%A1rio-9a8ac33e25da


