package main

import (
	"fmt"
	"math/rand"
)

func main() {
	harita := make(map[int]string)

	harita[0] = "adana"
	harita[1] = "adıyaman"
	harita[2] = "afyon"
	harita[3] = "ağrı"
	/*harita[4] = "amasya"
	harita[5] = "ankara"
	harita[6] = "antalya"
	harita[7] = "artvin"
	harita[8] = "aydın"
	harita[9] = "balıkesir"
	harita[10] = "bilecik"
	harita[11] = "bingöl"
	harita[12] = "bitlis"
	harita[13] = "bolu"
	harita[14] = "burdur"
	harita[15] = "bursa"
	harita[16] = "çanakkale"
	harita[17] = "çankırı"
	harita[18] = "çorum"
	harita[19] = "denizli"
	harita[20] = "diyarbakır"
	harita[21] = "edirne"
	harita[22] = "elazığ"
	harita[23] = "erzincan"
	harita[24] = "erzurum"
	harita[25] = "eskişehir"
	harita[26] = "gaziantep"
	harita[27] = "giresun"
	harita[28] = "gümüşhane"
	harita[29] = "hakkâri"
	harita[30] = "hatay"
	harita[31] = "ısparta"
	harita[32] = "mersin"
	harita[33] = "istanbul"
	harita[34] = "izmir"
	harita[35] = "kars"
	harita[36] = "kastamonu"
	harita[37] = "kayseri"
	harita[38] = "kırklareli"
	harita[39] = "kırşehir"
	harita[40] = "kocaeli"
	harita[41] = "konya"
	harita[42] = "kütahya"
	harita[43] = "malatya"
	harita[44] = "manisa"
	harita[45] = "kahramanmaraş"
	harita[46] = "mardin"
	harita[47] = "muğla"
	harita[48] = "muş"
	harita[49] = "nevşehir"
	harita[50] = "niğde"
	harita[51] = "ordu"
	harita[52] = "rize"
	harita[53] = "sakarya"
	harita[54] = "samsun"
	harita[55] = "siirt"
	harita[56] = "sinop"
	harita[57] = "sivas"
	harita[58] = "tekirdağ"
	harita[59] = "tokat"
	harita[60] = "trabzon"
	harita[61] = "tunceli"
	harita[62] = "şanlıurfa"
	harita[63] = "uşak"
	harita[64] = "van"
	harita[65] = "yozgat"
	harita[66] = "zonguldak"
	harita[67] = "aksaray"
	harita[68] = "bayburt"
	harita[69] = "karaman"
	harita[70] = "kırıkkale"
	harita[71] = "batman"
	harita[72] = "şırnak"
	harita[73] = "bartın"
	harita[74] = "ardahan"
	harita[75] = "ığdır"
	harita[76] = "yalova"
	harita[77] = "karabük"
	harita[78] = "kilis"
	harita[79] = "osmaniye"
	harita[80] = "düzce"

	*/

	fmt.Println("**********Plaka öğreten programa hoşgeldiniz**********\n\n")

	for oyun := true; oyun; {

		fmt.Println("1- Plakalar verilir şehir tahmini yaparsınız\n2- Şehirler verilir, plaka tahmini yaparsınız\n3- Çıkış\nTercihinize göre 1,2 veya 3 giriniz\n\n")
		secim := ""
		fmt.Scan(&secim)
		for secim != "1" && secim != "2" && secim != "3" {
			fmt.Println("Yanlış giriş yaptınız. Lütfen 1,2 veya 3 giriniz\n")
			fmt.Scan(&secim)
		}

		if secim == "1" {
			plakaOgrenme(harita)
		} else if secim == "2" {
			SehirOgrenme(harita)
		} else {
			oyun = false
		}

	}

	fmt.Println("Program sonlandırıldı")
}

func plakaOgrenme(harita map[int]string) {
	bilinenler := []int{}

	for {
		indeks := rand.Intn(len(harita))

		for icerir(bilinenler, indeks) {
			indeks = rand.Intn(len(harita))
		}

		fmt.Println("Programı sonlandırmak için bitir yaziniz\n", indeks+1, " numarali plaka hangi ile aittir?\nCevabınız tamamen küçük harflerden oluşmalıdır\n")
		cevap := ""
		fmt.Scan(&cevap)

		if cevap == "bitir" {
			fmt.Println("*****Plakaya göre şehir öğrenme programı sonlandırıldı*****\n\n")
			return
		}

		if cevap == harita[indeks] {
			fmt.Println("Tebrikler, dogru cevap\n\n")
			bilinenler = append(bilinenler, indeks)
			if len(bilinenler) == len(harita) {
				fmt.Println("*****Tebrikler, tum sehirleri dogru bildiniz*****\n\n")
				return
			}
		} else {
			fmt.Println("Yanlis cevap\n doğru cevap şu idi : " + harita[indeks] + "\n\n")
		}

	}
}

func SehirOgrenme(harita map[int]string) {
	bilinenler := make([]int, 0)
	for {
		indeks := rand.Intn(len(harita))
		for icerir(bilinenler, indeks) {
			indeks = rand.Intn(len(harita))
		}

		fmt.Println("Programı sonlandırmak -1 giriniz\n", harita[indeks], " ilinin plaka numarası nedir?\nCevabınız verilen ilin plaka numarası olmalıdır\n")
		var cevap int
		fmt.Scan(&cevap)

		if cevap == -1 {
			fmt.Println("şehire göre plaka öğrenme programı sonlandırıldı")
			return
		}

		if cevap == indeks+1 {
			fmt.Println("Tebrikler, dogru cevap\n\n")
			bilinenler = append(bilinenler, indeks)
			if len(bilinenler) == len(harita) {
				fmt.Println("*****Tebrikler, tum plakaları dogru bildiniz*****\n\n")
				return
			}
		} else {
			fmt.Println("Yanlis cevap\n doğru cevap şu idi : ", indeks+1, "\n\n")
		}
	}
}

func icerir(slice []int, element int) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}
