package main

import (
	"fmt"
	"os"
	"os/exec"

	//"bufio"
	"golang.org/x/term"
)

func main() {
	var A int
poczatek_gry:
	//definicja planszy do gry
	fmt.Println("\n autor:Witold Surówka\n Zasady gry:\n -zdefiniuj wielkość planszy,\n -startujesz w lewym dolnym rogu,\n -celem jest dostać się w górny prawy róg,\n -używaj jako strzałek WASD,\n -GL&HF.\n\n")
	fmt.Println("zdefinuj rozmiar planszy AxA (A_min=1, A_max=30)\n podaj A= ")
podaj:
	_, err := fmt.Scanf("%d", &A)
	if err != nil {
		fmt.Println(err)
		return
	}
	if A < 1 || A > 30 {
		fmt.Println("podano nieodpowiedni rozmiar pola,\n Podaj A=")
		goto podaj
	}

	//czyszcze konsole---------------------------------------------------
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	//czyszcze konsole---------------------------------------------------
	fmt.Printf("plansza %v x %v \n", A, A)

	//pierwsze inicjowanie matrixa - do startu
	matrix := make([][]int, A)
	for g := 0; g < A; g++ {
		matrix[g] = make([]int, A)
	}

	matrix[0][0] = 1
	//pierwsze drukowanie matrixa - do startu
	for g := A - 1; g >= 0; g-- {
		fmt.Println(matrix[g])
	}

	//definicje do petli gry
	var ca byte = 97
	var cd byte = 100
	var cw byte = 119
	var cs byte = 115
	var x, y int = 0, 0
	var k, i int = 1, 0
	//pętla gry---------------------------------------------------------------------------
	for k < 2 {
		i = i + 1
		if k == A {
			goto Koniec_gry_wygrana
		}

		for g := 0; g < A; g++ {
			matrix[g] = make([]int, A)
		}
		//łapanie znaku-------------------------------------------------------------------------
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		b := make([]byte, 1)
		_, err = os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		//łapanie znaku-------------------------------------------------------------------------

		//czyszcze konsole---------------------------------------------------
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		//czyszcze konsole---------------------------------------------------
		d := b[0]
		//definiuje dzialanie klawiszy
		switch d {
		case ca:
			{
				x = x - 1
				if x < 0 {
					fmt.Println("niedozwolony ruch!")
					x = 0
				} else {
					fmt.Println("idziesz w złą stronę...")
				}
			}

		case cd:
			{
				x = x + 1
				fmt.Println("dobrze ci idzie...")
			}
			if x > A-1 {
				//czyszcze konsole---------------------------------------------------
				cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
				cmd.Stdout = os.Stdout
				cmd.Run()
				//czyszcze konsole---------------------------------------------------
				fmt.Println("niedozwolony ruch!")
				x = A - 1
			}
		case cs:
			{
				y = y - 1
				if y < 0 {
					fmt.Println("niedozwolony ruch!")
					y = 0
				} else {
					fmt.Println("idziesz w złą stronę...")
				}
			}
		case cw:
			{
				y = y + 1
				fmt.Println("dobrze ci idzie...")
			}
			if y > A-1 {
				//czyszcze konsole---------------------------------------------------
				cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
				cmd.Stdout = os.Stdout
				cmd.Run()
				//czyszcze konsole---------------------------------------------------
				fmt.Println("niedozwolony ruch!")
				y = A - 1
			}
		default:
			{
				fmt.Println("wcisnieto zły klawisz!")
			}
		} //119 97 115 100

		matrix[y][x] = 1
		//drukuję aktualną planszę
		for g := A - 1; g >= 0; g-- {
			fmt.Println(matrix[g])
		}
		//jezeli doszedles do konca to wygrales
		if x == A-1 && y == A-1 {
			goto Koniec_gry_wygrana
		}
	}
	//pętla gry-------------------------------------------------------------------------
Koniec_gry_wygrana:
	fmt.Println("\n wygrana!!! \n")
	var nMin int = 2 * (A - 1)
	if i == 1 {
		fmt.Printf("Gratulacje, wykonałeś najmniejszą możliwą liczbę ruchów (0) - jesteś baaardzo mądry! \n")
	} else {
		if i == nMin {
			fmt.Printf("Gratulacje, wykonałeś najmniejszą możliwą liczbę ruchów (%v) - jesteś mądry! \n", i)
		} else {
			fmt.Printf("Wykonałeś %v ruchów. Najmniejsza liczba ruchów możliwych do ukończenia gry to: %v. \n", i, nMin)
			fmt.Println(" ")
		}
	}
czy_jeszcze_raz:
	fmt.Println("Czy chcesz zagrac jeszcze raz? (t/n)")
	//łapanie znaku-------------------------------------------------------------------------
	b2 := make([]byte, 1)
	_, err = os.Stdin.Read(b2)
	if err != nil {
		fmt.Println(err)
		return
	}
	//łapanie znaku-------------------------------------------------------------------------
	d1 := b2[0]
	var ct byte = 116
	var cn byte = 110
	switch d1 {
	case ct:
		{
			//czyszcze konsole---------------------------------------------------
			cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
			//czyszcze konsole---------------------------------------------------

			goto poczatek_gry
		}
	case cn:
		{
			os.Exit(3)
		}
	default:
		{
			fmt.Println("wcisnieto zły klawisz!")
			goto czy_jeszcze_raz
		}
	}
}
