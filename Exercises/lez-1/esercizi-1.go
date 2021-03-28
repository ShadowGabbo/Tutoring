/*1) Scrivere un programma Go main.go che non abbia nessuna istruzione.
2)  Problema: Scrivere un programma Go echo.go che legge un valore di tipo int da tastiera e lo stampa.
3)  Scrivere un programma Go operazioni.go che, dati due numeri float64 da standard input, esegua addizione, sottrazione, moltiplicazione e divisione dei due numeri.
4)  Scrivere un programma Go quoziente_resto.go che legge da standard input un dividendo e un divisore (interi) e calcola il quoziente e il resto.
5)  Scrivere un programma Go angolo_triangolo.go che, date in input le ampiezze (float) di due angoli di un triangolo, determini l’ampiezza del terzo angolo
6)  Scrivere un programma Go scambio.go che, dati in input due valori di tipo int, controlli che i valori letti siano di tipo int, li stampi nell’ordine in cui sono stati forniti, scambi i due valori, e poi esegua le stessa istruzione di stampa (per mostrare che i valori sono stati effettivamente scambiati tra le due var).
7)  Scrivere un programma Go random.go che generi un numero random int tra 0 e 100 e lo stampi.
8)  Scrivere un programma Go conversione_giorni.go che converta un numero specificato di giorni (letto da tastiera) in anni, settimane, giorni senza mai usare l’operazione di sottrazione. Si ignorino gli anni bisestili.*/

/*1)package main

func main()  {

}
__________________________
2)package main

import(
  "fmt"
)

func main()  {
  var x int
  fmt.Scan(&x)
  fmt.Print(x)
}
__________________________
3)package main

import "fmt"

func main()  {
  var x,y float64
  fmt.Scan(&x,&y)
  addizione := x + y
  sottrazione := x - y
  divisone := x/y
  moltiplicazione := x*y
  fmt.Println(addizione)
  fmt.Println(sottrazione)
  fmt.Println(divisione)
  fmt.Println(moltiplicazione)
}

__________________________
4)package main

import "fmt"

func main()  {
  var dividendo, divisore int
  fmt.Scan(&dividendo, &divisore)
  quoziente := dividendo/divisore
  resto := dividendo%divisore
  fmt.Printf("%d\n%d\n",quoziente,resto)
}
__________________________
5)package main

import "fmt"

func main()  {
  const ampiezza = 180
  var x,y float64
  fmt.Scan(&x,&y)
  z := ampiezza - x - y
  fmt.Println(z)
}

6)package main

import "fmt"

func main()  {
  var x,y int
  fmt.Scan(&x,&y)
  fmt.Printf("X è di tipo %T\tY è di tipo %T\t",x,y)
  fmt.Println("before change",x,y)
  x, y = y, x
  fmt.Println("after change",x,y)
}
__________________________

7)package main

import (
  "fmt"
  "math/rand"
)

func main()  {
  rand.seed(unix.Nano()time.Now())
  random := rand.Intn(101)
  fmt.Println(random)
}
 __________________________

 8)package main

 import "fmt"

 func main()  {
   var n int
   const giorniSettimana = 7
   const giorniAnno = 365
   fmt.Println("quanti giorni vuoi convertire in anno?")
   fmt.Scan(&n)
   anni := n / giorniAnno
   resto1 := n % giorniAnno
   settimane := resto1 / giorniSettimana
   giorni := resto1 % giorniSettimana
   fmt.Printf("anni %d\tsettimane %d\t giorni %d\n", anni, settimane, giorni)
 }*/
