package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	matris = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	sira   = "X"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for !Bitti() {
		Goster()
		metin := ""
		isaretli := true

		for isaretli {
			fmt.Printf("\nOyuncu %s, sırasıyla satır ve sütun endeksini giriniz. Örnek: 0-2\n", RenklendirSira(sira))
			metin, _ = reader.ReadString('\n')
			metin = strings.TrimSpace(metin)
			for uygunDegil(metin) {
				Goster()
				fmt.Println("Girdiğiniz değer uygun değil. Sırasıyla satır ve sütun endeksi - ile ayrılarak girilmelidir. Örnek: 0-2\nSatır ve sütun endeksleri 0, 1 veya 2 olmalıdır.")
				fmt.Printf("\nOyuncu %s, sırasıyla satır ve sütun endeksini giriniz. Örnek: 0-2\n", RenklendirSira(sira))
				metin, _ = reader.ReadString('\n')
				metin = strings.TrimSpace(metin)
			}
			diziIndeks := strings.Split(metin, "-")
			satir, _ := strconv.Atoi(diziIndeks[0])
			sutun, _ := strconv.Atoi(diziIndeks[1])

			isaretli = matris[satir][sutun] != " "

			if isaretli {
				Goster()
				fmt.Println("Bu alan daha önceden işaretlenmiş. Sadece boş bir yere işaretleme yapabilirsiniz. Geçerli bir değer giriniz.")
			}
		}

		diziIndeks := strings.Split(metin, "-")
		satir, _ := strconv.Atoi(diziIndeks[0])
		sutun, _ := strconv.Atoi(diziIndeks[1])
		matris[satir][sutun] = sira
		SiraDegis()
	}

	Goster()
	vaziyet := Sonuc()

	switch vaziyet {
	case "beraber":
		mesajYaz("Oyun berabere bitti.", "sari")
	case "X", "O":
		mesajYaz(fmt.Sprintf("Oyunu kazanan: %s", Renklendir(vaziyet)), "yesil")
	default:
		mesajYaz("Oyun bilinmeyen bir nedenle sona erdi.", "kirmizi")
	}
}

func uygunDegil(metin string) bool {
	return len(metin) != 3 || (metin[0] != '0' && metin[0] != '1' && metin[0] != '2') || (metin[2] != '0' && metin[2] != '1' && metin[2] != '2') || metin[1] != '-'
}

func Goster() {
	Temizle()
	fmt.Println("   0   1   2")
	fmt.Println("  +---+---+---+")

	for i, row := range matris {
		fmt.Print(i, " ")
		for _, harf := range row {
			fmt.Print("| ", Renklendir(harf), " ")
		}
		fmt.Println("|")
		if i < len(matris)-1 {
			fmt.Println("  +---+---+---+")
		}
	}

	if !Bitti() {
		fmt.Printf("\nSıra: %s\n", RenklendirSira(sira))
	}
}

func Temizle() {
	fmt.Print("\033[H\033[2J")
}

func Bitti() bool {
	for _, xo := range []string{"X", "O"} {
		for i := 0; i < 3; i++ {
			if matris[i][0] == xo && matris[i][1] == xo && matris[i][2] == xo {
				return true
			}
			if matris[0][i] == xo && matris[1][i] == xo && matris[2][i] == xo {
				return true
			}
		}
		if matris[0][0] == xo && matris[1][1] == xo && matris[2][2] == xo {
			return true
		}
		if matris[0][2] == xo && matris[1][1] == xo && matris[2][0] == xo {
			return true
		}
	}

	return HepsiDolu()
}

func HepsiDolu() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matris[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func SiraDegis() {
	if sira == "X" {
		sira = "O"
	} else {
		sira = "X"
	}
}

func Sonuc() string {
	for _, xo := range []string{"X", "O"} {
		for i := 0; i < 3; i++ {
			if matris[i][0] == xo && matris[i][1] == xo && matris[i][2] == xo {
				return xo
			}
			if matris[0][i] == xo && matris[1][i] == xo && matris[2][i] == xo {
				return xo
			}
		}
		if matris[0][0] == xo && matris[1][1] == xo && matris[2][2] == xo {
			return xo
		}
		if matris[0][2] == xo && matris[1][1] == xo && matris[2][0] == xo {
			return xo
		}
	}

	return "beraber"
}

func Renklendir(harf string) string {
	switch harf {
	case "X":
		return "\x1b[31;1mX\x1b[0m"
	case "O":
		return "\x1b[34;1mO\x1b[0m"
	default:
		return " "
	}
}

func RenklendirSira(sira string) string {
	return fmt.Sprintf("\x1b[1;33m%s\x1b[0m", sira)
}

func rengiSec(color string) string {
	switch color {
	case "kirmizi":
		return "31"
	case "yesil":
		return "32"
	case "sari":
		return "33"
	default:
		return "0"
	}
}

func mesajYaz(message string, color string) {
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("\x1b[1;%sm%s\x1b[0m\n", rengiSec(color), message)
	fmt.Println(strings.Repeat("-", 40))
}
