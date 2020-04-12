package main

import (
	"fmt"
)

func generator(vparam int) *int {
	return &vparam
}

type Pessoa struct {
	nome string
	idade string
}

// uma copia da struct variavel é enviada para a funcao.
// a função altera o valor do field nome do struct copiado
func changeStructPessoa(pessoa Pessoa, new_name string){
	pessoa.nome = new_name
	fmt.Printf(" %v\n " , pessoa)
}
func changeStructPessoaByPointer(pessoa *Pessoa, new_name string){
	pessoa.nome = new_name
}

/*
Exemplificando o uso de ponteiros de forma simples na linguagem de programação GO.
 -> Problema: Realizar alterações em uma variavel a partir de uma função.
			* As alterações são realizadas, porem em um copia da variavel.
	Resolução:
			Passamos a referencia da variavel ou seja o seu endereço de memoria atraves
			de um apontador e a partir dai fazemos as alterações permanentes atraves
			do endereço de memoria da mesma.

*/
// Apesar de muito util os ponteiros em Go, não possuem operações aritimeticas.
// Logo é preciso desreferenciar o ponteiro para trabalhar com suas variaveis reais.
func main(){

	var pessoa Pessoa = Pessoa{nome:"osvaldo", idade: "23"}

	fmt.Printf(" %v \n", pessoa)

	// a alteração da copia ocorre na funcao.
	changeStructPessoa(pessoa,"airon")

	// depois da função aplicada o valor original da variavel permanece imutavel.
	fmt.Printf(" %v \n", pessoa)

	// acessando o endereço de memoria da variavel struct e realizando as alterações
	// permanentes.
	changeStructPessoaByPointer(&pessoa,"cavalcanti")

	fmt.Printf(" %v \n", pessoa)

}