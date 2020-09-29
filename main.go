// Perdir el dia y mes de nacimiento y mediante esto destinar el signo zodiacal.

package main

import "fmt"

func main(){
	var dia int64 // son los dias del mes.
	var mes int64 //meses 1,2,3,4,5,6,7,8,9,10,11,12
	var seleccionSigno int64 //Dato para el switch

	var signo string

	//Preguntamos el dia y el mes
	fmt.Scan(&dia)
	if dia == 0 { fmt.Scan(&dia)} //Esto seria por si pone 0 volvemos a preguntarle.
	fmt.Scan(&mes)
	if mes == 0 { fmt.Scan(&mes)} //Esto seria por si pone 0 volvemos a preguntarle.

	if dia >= 20 && dia <= 31 && mes == 1 {
		seleccionSigno = 1
	} else if dia >= 1 && dia <= 18 && mes == 2 {
		seleccionSigno = 1 // 1 = a Acuario
	} else if dia >= 19 && dia <= 29 && mes == 2 {
		seleccionSigno = 2
	} else if dia >= 1 && dia <= 20 && mes == 3 {
		seleccionSigno = 2 // 2 = a Picis
	} else if dia >= 21 && dia <= 31 && mes == 3 {
		seleccionSigno = 3
	} else if dia >= 1 && dia <= 19 && mes == 4 {
		seleccionSigno = 3 // 3 = a Aries
	} else if dia >= 20 && dia <= 30 && mes == 4 {
		seleccionSigno = 4
	} else if dia >= 1 && dia <= 20 && mes == 5 {
		seleccionSigno = 4 // 4 = a Tauro 
	} else if dia >= 21 && dia <= 31 && mes == 5 {
		seleccionSigno = 5
	} else if dia >= 1 && dia <= 20 && mes == 6 {
		seleccionSigno = 5 // 5 = a Geminis
	} else if dia >= 21 && dia <= 30 && mes == 6 {
		seleccionSigno = 6 
	} else if dia >= 1 && dia <= 22 && mes == 7 {
		seleccionSigno = 6 // 6 = a Cancer
	} else if dia >= 23 && dia <= 31 && mes == 7 {
		seleccionSigno = 7
	} else if dia >= 1 && dia <= 22 && mes == 8 {
		seleccionSigno = 7 // 7 = a Leo
	} else if dia >= 23 && dia <= 31 && mes == 8 {
		seleccionSigno = 8
	} else if dia >= 1 && dia <= 22 && mes == 9 {
		seleccionSigno = 8 // 8 = a Virgo
	} else if dia >= 23 && dia <= 30 && mes == 9 {
		seleccionSigno = 9
	} else if dia >= 1 && dia <= 22 && mes == 10 {
		seleccionSigno = 9 // 9 = a Libra
	} else if dia >= 23 && dia <= 31 && mes == 10 {
		seleccionSigno = 10
	} else if dia >= 1 && dia <= 21 && mes == 11 {
		seleccionSigno = 10 // 10 = a Escorpio
	} else if dia >= 22 && dia <= 30 && mes == 11 {
		seleccionSigno = 11
	} else if dia >= 1 && dia <= 21 && mes == 12 {
		seleccionSigno =  11 //11 = a Sagitario
	} else if dia >= 22 && dia <= 31 && mes == 12 {
		seleccionSigno = 12
	} else if dia >= 1 && dia <= 19 && mes == 1 {
		seleccionSigno = 12 // 12 = a Capricornio
	} 

	switch seleccionSigno {
	case 1:
		signo = "acuario"
		fmt.Println(signo)

	case 2:
		signo = "piscis"
		fmt.Println(signo)
	case 3:
		signo = "aries"
		fmt.Println(signo)
	case 4:
		signo = "tauro"
		fmt.Println(signo)
	case 5:
		signo = "geminis"
		fmt.Println(signo)
	case 6:
		signo = "cancer"
		fmt.Println(signo)
	case 7: 
		signo = "leo"
		fmt.Println(signo)
	case 8: 
		signo = "virgo"
		fmt.Println(signo)
	case 9:
		signo = "libra"
		fmt.Println(signo)
	case 10:
		signo = "escorpio"
		fmt.Println(signo)
	case 11:
		signo = "sagitario"
		fmt.Println(signo)
	case 12:
		signo = "capricornio"
		fmt.Println(signo)
	default:

	}
}