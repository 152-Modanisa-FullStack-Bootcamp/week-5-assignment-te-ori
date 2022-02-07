package assignment

import (
	"math"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y

	// 0 veya 0'dan büyük iki sayının toplamı sayıların
	// kendisinden küçük olamaz. Eğer toplam bu sayılardan
	// birinden küçükse overflow gerçekleşmiş demektir.
	return sum, sum < x || sum < y
}

func CeilNumber(f float64) float64 {
	integer, fraction := math.Modf(f)

	if fraction == 0 {
		return integer
	} else if fraction <= 0.25 {
		return integer + 0.25
	} else if fraction <= 0.5 {
		return integer + 0.5
	} else if fraction <= 0.75 {
		return integer + 0.75
	} else {
		return integer + 1
	}

}

func AlphabetSoup(s string) string {

	// harfler ayrıca `byte` tipinde. 'a' -> 65... şeklinde gidiyor.
	// alfabe uzunluğunda dizi sanımlıyoruz. Dizi alfebetik sıraya göre
	// sıralamış sayıyoruz.
	var letters [26]byte

	// kelimdeki her harfi teker teker gezerek dizimide o harfe karşılık
	// gelen yerdeki sayıyı bir arttırıyoruz. Böylece her harfin kaç
	// defa geçtiğini sayıyoruz
	for _, letter := range s {
		letters[letter-'a']++
	}

	sb := strings.Builder{}

	var i byte

	// alfabedeki her harf için sırayla geziyoruz
	for i = 0; i < 26; i++ {
		var j byte

		// harf sırasına göre ka. defa geçtiyse o kadar harfi tekrarlıyoruz
		// hiç yoksa 0 geçicek yani boş
		for j = 0; j < letters[i]; j++ {
			sb.WriteByte('a' + i)
		}
	}

	return sb.String()
}

func StringMask(s string, n uint) string {
	lenOfWord := len(s)
	if lenOfWord == 0 {
		return "*"
	}

	if n == 0 || n >= uint(lenOfWord) {
		return strings.Repeat("*", lenOfWord)
	}
	arr := []string{
		s[0:n],
		strings.Repeat("*", lenOfWord-int(n)),
	}
	return strings.Join(arr, "")
}

func WordSplit(arr [2]string) string {
	// aramayı kolaylaştırmak için kelimelerimizi bir map a atacağız
	splited := strings.Split(arr[1], ",")
	dictionary := make(map[string]struct{})

	for _, s := range splited {
		dictionary[s] = struct{}{}
	}

	// sonra bize verilen kelimeleri sırayla tarayacağız
	for i := 0; i < len(arr); i++ {

		// kelimenin ilk harfiden başlayarak sözlüğümüzde var mı
		// kontrol edeceğiz
		for j := 0; j <= len(arr[i]); j++ {
			_, ok := dictionary[arr[i][:j]]

			// eğer aradığımız kelime sözlükte yoksa o kelimeyle
			// sözlükteki kelimerele oluşturamıypruz demektir.
			if !ok {
				continue
			}

			// 0..'dan i'ye kadar olan harflerle oluşturduğumuz
			// kelime sözlükte varsa, geriye kalanların da tamamının
			// olması gerekir.
			_, ok = dictionary[arr[i][j:]]

			if !ok {
				continue
			}

			// eğer şartlar sağlanırsa daha fazla aramaya gerek yok
			return arr[i][:j] + "," + arr[i][j:]
		}
	}
	return "not possible"
}

func VariadicSet(i ...interface{}) []interface{} {
	m := make(map[interface{}]struct{})

	for _, element := range i {
		m[element] = struct{}{}
	}
	var result []interface{}

	for k := range m {
		result = append(result, k)
	}

	return result
}
