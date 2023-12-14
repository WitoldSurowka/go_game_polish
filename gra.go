package main

import (
	"fmt"
	"os"
	"os/exec"

	//"bufio"
	"golang.org/x/term"
)

const (
	ca byte = 97
	cd byte = 100
	cw byte = 119
	cs byte = 115
	ct byte = 116
	cn byte = 110
)

var (
	A        int
	x, y     = 0, 0
	i        = 0
	b, b2    = make([]byte, 1), make([]byte, 1)
	d, d2    byte
	czy_goto bool
	err      error
)

func field_size() (bool, int) {
	_, err := fmt.Scanf("%d", &A)
	if err != nil {
		fmt.Println(err)
	}
	czy_goto = false
	if A < 1 || A > 30 {
		fmt.Println("podano nieodpowiedni rozmiar pola,\n Podaj A=")
		czy_goto = true
	}
	return czy_goto, A
}

func getchar() ([]byte, error) {
	_, err := os.Stdin.Read(b2)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b2, err
}

func clearconsole() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func key_pressed(d byte) (int, int, int) {
	switch d {
	case ca:
		{
			x = x - 1
			if x < 0 {
				fmt.Println("niedozwolony ruch!")
				x = 0
			} else {
				fmt.Println("idziesz w złą stronę...")
				i = i + 1

			}
		}
	case cd:
		{
			x = x + 1
			fmt.Println("dobrze ci idzie...")
			i = i + 1

		}
		if x > A-1 {

			//czyszcze konsole---------------------------------------------------
			clearconsole()
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
				i = i + 1

			}
		}
	case cw:
		{
			y = y + 1
			fmt.Println("dobrze ci idzie...")
			i = i + 1

		}
		if y > A-1 {

			//czyszcze konsole---------------------------------------------------
			clearconsole()
			//czyszcze konsole---------------------------------------------------

			fmt.Println("niedozwolony ruch!")
			y = A - 1
		}
	default:
		if i != 0 {
			fmt.Println("wcisnieto zły klawisz!")
		}
	} //119 97 115 100
	return x, y, i
}

func yes_no_press(d2 byte) bool {

	switch d2 {
	case ct:
		{
			//czyszcze konsole---------------------------------------------------
			clearconsole()
			//czyszcze konsole---------------------------------------------------

			czy_goto = true
		}
	case cn:
		{
			os.Exit(3)
		}
	default:
		{
			fmt.Println("wcisnieto zły klawisz!")
			czy_goto = false
		}
	}
	return czy_goto
}

func game() {

	//definicja planszy do gry-------------------------------------------
poczatek_gry:
	x, y = 0, 0 //jezeli znowu zaczynasz, to wyzeruj swoja pozycje po ostatniej grze
	i = 0

	fmt.Println("\n autor:Witold Surówka\n Zasady gry:\n -zdefiniuj wielkość planszy,\n -startujesz w lewym dolnym rogu,\n -celem jest dostać się w górny prawy róg,\n -używaj jako strzałek WASD,\n -GL&HF.\n\n")
	fmt.Println("zdefinuj rozmiar planszy AxA (A_min=1, A_max=30)\n podaj A= ")
podaj:
	czy_goto, A = field_size()
	if czy_goto == true {
		goto podaj
	}
	//definicja planszy do gry-------------------------------------------

	//czyszcze konsole---------------------------------------------------
	clearconsole()
	//czyszcze konsole---------------------------------------------------

	fmt.Printf("plansza %v x %v \n", A, A)

	//pierwsze inicjowanie matrixa - do startu---------------------------
	matrix := make([][]int, A) //przygotowanie miejsca
	for g := 0; g < A; g++ {   //zdefiniuj wiersze w liczbie A
		matrix[g] = make([]int, A) //kazdy wiersz sklada sie z tablicy zer w liczbie A
	}
	matrix[0][0] = 1 //pozycja początkowa
	//pierwsze inicjowanie matrixa - do startu---------------------------

	//pierwsze drukowanie matrixa - do startu----------------------------
	for g := A - 1; g >= 0; g-- { //drukujemy od gory do dolu
		fmt.Println(matrix[g])
	}
	//pierwsze drukowanie matrixa - do startu----------------------------

	//pętla gry---------------------------------------------------------------------------
	for { //petla w nieskonczonosc

		if A == 1 { //jezeli plansza trywialna
			goto Koniec_gry_wygrana
		}

		for g := 0; g < A; g++ { //flashing matrixa
			matrix[g] = make([]int, A)
		}

		//łapanie znaku-------------------------------------------------------------------------
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		b, err = getchar()
		if err != nil {
			fmt.Println(err)
			return
		}
		//łapanie znaku-------------------------------------------------------------------------

		//czyszcze konsole---------------------------------------------------
		clearconsole()
		//czyszcze konsole---------------------------------------------------

		//definiuje dzialanie klawiszy
		d = b[0]
		x, y, i = key_pressed(d)

		matrix[y][x] = 1 //drukuję aktualną pozycję

		for g := A - 1; g >= 0; g-- { //drukuję od góry do dołu
			fmt.Println(matrix[g])
		}
		if x == A-1 && y == A-1 { //jezeli doszedles do konca to wygrales
			goto Koniec_gry_wygrana
		}
	}
	//pętla gry-------------------------------------------------------------------------

Koniec_gry_wygrana:
	fmt.Println("\n wygrana!!! \n")
	var nMin int = 2 * (A - 1)
	if i == 0 {
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
	b2, err = getchar()
	if err != nil {
		fmt.Println(err)
		return
	}
	//łapanie znaku-------------------------------------------------------------------------

	d2 = b2[0] //przypisz przejety znak do d2 i uruchom funckcje z nim jako argumentem
	czy_goto = yes_no_press(d2)
	if czy_goto == true {
		goto poczatek_gry
	}
	if czy_goto == false {
		goto czy_jeszcze_raz
	}
}

func main() {
	game()
}
