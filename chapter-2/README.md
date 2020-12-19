<details>
<summary>Fonksiyonlar :</summary>

### Fonksiyonlar içlerine parametre girilebilen ve işlemler yapabilen birimlerdir. Matematikteki fonksiyonlar ile aynı mantıkta çalışan bu birimlerden bir örneği inceleyelim.  

```go
package main

import "fmt"

func topla(a int, b int) int {
	return a + b //a ve b’nin toplamını döndürür.
}

func main() {
	fmt.Println(topla(2, 5)) //2+5 sonucunu ekrana bastır
}
```
### Main fonksiyonu içerisinde topla(2,5) fonksiyonu ile 2 ve 5 sayısının toplamını ekrana bastırmış olduk. Yani ekrana 7 sayısı verildi.

### örnek.2

```go
package main

import "fmt"

func yazdir() {
	fmt.Println("yazı yazdırdık")
}

func main() {
	yazdir()
}
```
</details>

<details>
<summary>Fonksiyonlar Çeşitleri :</summary>

### Golang’ta genel olarak 3 çeşit fonksiyon yapısı bulunmaktadır. Hemen bu çeşitleri görelim.  
### Variadic Fonksiyonlar :   
Variadic fonksiyon tipi ile fonksiyonumuza kaç tane değer girişi olduğunu belirtmeden istediğiniz kadar değer girebilirsiniz.

```go
package main

import "fmt"

func toplama(sayilar ...int) int {
    toplam := 0
    for _, n := range sayilar {
        toplam += n
    }
    return toplam
}

func main() {
    fmt.Println(toplama(3, 4, 5, 6)) //18
}
```

range'in burdaki kullanım amacı bunu for döngüsü ile yaptığımız öğe uzunluğuna göre işlemi sürdürürüz. yani ne kadar sayı eklersek ona göre şekillenir.

### Closure Fonksiyonlar :   
Closure fonksiyonlar ile değişkenlerimizi fonksiyon olarak tanımlayabiliriz.

```go
package main

import "fmt"

func main() {
    toplam := func(x, y int) int {
        return x + y
    }
    fmt.Println(toplam(2, 3))
}
```
### Recursive (İç-içe) Fonksiyonlar : 


Recursive fonksiyonlar yazdığımız fonksiyonun içinde aynı fonksiyonu kullanmamız demektir. Fonksiyonumun tüm işlemler bittiğinde return olur.   
```go
package main

import "fmt"

func main() {
    fmt.Println(faktoriyel(4))
}

func faktoriyel(a uint) uint {
    if a == 0 {
        return 1
    }
    return a * faktoriyel(a-1)
}
```

### Anonim Fonksiyonlar :

Adından belli olduğu kadarıyla isimleri yoktur yazıldığı yerde çalışırlar. İsimleri olmadığı için diğer funclar gibi parametre verilmediği için parametresi sonuna eklenir çalışması için. 

```go
package main

import "fmt"

func main() {
	metin := "Merhaba Dünya"

	func(a string) {
		fmt.Println(a)
	}(metin)
}
```

</details>

<details>
<summary>Boş Tanımlayıcılar</summary>

# Boş Tanımlayıcılar

Golang kodlarımızda bazen 2 adet değer döndüren fonksiyonlar kullanırız. Bu değerlerden hangisini kullanmak istemiyorsak, değişken adı yerine **\_ \(alt tire\)** kullanırız.

Örneğimizi görelim:

```go
package main

import "fmt"

func fonksiyonumuz(girdi int) (int, int) {
	işlem1 := girdi / 2
	işlem2 := girdi / 4
	return işlem1, işlem2
}

func main() {
	ikiyeböl, dördeböl := fonksiyonumuz(16)
	fmt.Println(ikiyeböl, dördeböl)
}
```

Gördüğünüz gibi fonksiyonumuzdan dönen iki değeri de değişkenlere atadık. Eğer birini atamak istemeseydik şöyle yapardık:

```go
package main

import "fmt"

func fonksiyonumuz(girdi int) (int, int) {
	işlem1 := girdi / 2
	işlem2 := girdi / 4
	return işlem1, işlem2
}

func main() {
	ikiyeböl, _ := fonksiyonumuz(16)
	fmt.Println(ikiyeböl)
}
```

Yukarıdaki kodlarımızda fonksiyonumuzun 4’e bölme özelliğini kullanmak istemediğimizden dolayı boş tanımlama işlemi yaptık.

Boş tanımlama işlemleri çoğunlukla Golang’ta programcılar tarafından hata çıktısını kullanmak istenmediğinizde yapılıyor.


</details>

<details>
<summary>Döngüler</summary>

# Döngüler

Programlama dillerinde while, do while ve for döngüleri vardır ama golangda tek for döngüsü vardır.

### Standart For Kullanımı:   
```go 
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

### SADECE KOŞUL BELİRTEREK KULLANMA
bu kullanımda for while ile aynı mantıkta çalışır. parametre olarak koşul belirtilir.

```go
package main

import "fmt"

func main() {
	count := 0
	for count < 10 {
		fmt.Println(count)
		deger++
	}
}
```
</details>

<details>
<summary>If-Else</summary>

### If-Else

Tğrkçe karşılık olarak bakarsan if = eğer, Else = yoksa anlamına gelir ve bizim golangda koşul durumları oluşturmamızı sağlar.   
Basit olarak syntax yapısı böyledir :   
```go
if koşul {
	//Koşul sağlandığında yapılacak işlemler
} else {
	//Koşul sağlanmadığında yapılacak işlemler
}
```
### else if kullanımı:

If-Else akışında birden fazla koşul kontrolü ekleyebiliriz. Bunu else if deyimi ile yapabiliriz. Kısaca bakacak olursak;    
```go
i := 5
if i == 5 {
 fmt.Println("i'nin değeri 5'tir.")
} else if i==3{
 fmt.Println("i'nin değeri 3'tür.")
}else{
 fmt.Println("i'nin değeri belirsiz.")
}
```
</details>

<details>
<summary>Switch Yapısı</summary>

### Switch Yapısı:

Türkçe karşılığı anahtardır. Switch deyimi if-else gibi koşul amaçlı çalışır.Case deyimi durumu ifade eder.

```go
package main
import "fmt"
func main() {
 i := 5
 switch i {
  case 5:
   fmt.Println("i eşittir 5")
  case 10:
   fmt.Println("i eşittir 10")
  case 15:
   fmt.Println("i eşittir 15")
 }
}
```
Switch’te koşulların gerçekleşmediği zaman işlem uygulamak istiyorsak bunu default terimi ile yaparız. Örnek;   
```go
i := 5
switch i {
 case 5:
  fmt.Println("i eşittir 5")
 default:
  fmt.Println("i bilinmiyor")
}
```
### Koşulsuz Switch
Switch’in tanımını daha iyi anlayabilmeniz için koşulsuz switch kullanımına örnek verelim. Bu yöntemde switch deyiminin yanına koşul girmek yerine case deyiminin yanına koşul giriyoruz.   
```go
package main
import "fmt"
func main() {
 i := 5
 switch {
  case i == 5: //i=5 olduğu için diğer case’ler sorgulanmaz
   fmt.Println("i eşittir 5")
  case i < 10:
   fmt.Println("i küçüktür 10")
  case i > 3:
   fmt.Println("i büyüktür 3")
 }
}
```
</details>

<details>
<summary>Defer Fonksiyonu</summary>

### Defer Fonksiyonu:
Defer kelimesinin Türkçe’deki karşılığı ertelemektir. Bu deyimi yapacağımız işlemin başına eklersek o işlemi içerisinde bulunduğu fonksiyonun içindeki işlemlerden sonra çalıştırır.   
```go
package main
import "fmt"
func main() {
 defer fmt.Println("İlk Cümle")
 fmt.Println("İkinci Cümle")
}
```
I/O:
```
İkinci Cümle
İlk Cümle
```
gibi bir basit örnekler açıkladık.
</details>