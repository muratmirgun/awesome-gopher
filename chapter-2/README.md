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