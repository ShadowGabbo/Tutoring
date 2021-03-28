package main

func main() {
	//dichiarazione di una variabile
	//var [name] [tipo]
	var Giovanni string
	//assegnamento di una variabile
	// [name] = X
	Giovanni = "Studente"
	//copia di una variabile (INDIRIZZO)
	var Gabriele string
	Gabriele = Giovanni
	//utilizzo
	Gabriele = "Pippoè"

	/*off-topic
	slice_stringhe := string["cane","bianco","ciao"]
	for index,stringa := range slice_stringhe{
		for index2,char := range stringa{
		}
	}
	for (i:=0; i<len();i++){
		for (j:=0;j<len();j++){
			slice_stringhe[i][j]
		}
	}
	*/

	//bool
	var boolean bool = false / true

	//nil string/mappe

	var stringa string //array utf-8

	//int (32 o 64)bit 1/2/4/8 byte
	var integer int

	//float32 float64
	var floating_point float32
}
