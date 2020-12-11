# Basics!
 
## Bu Kısımda Fazla Kod Örneği Göremeyebilirsiniz.
<details>
<summary>Değişkenler :</summary> 
 
### Integer Türler
### Öncelikle tüm integer türleri bir görelim;
  -int, int8, int16, int32, int64  
  -uint, uint8, uint16, uint32, uint64, uintptr

### NOT
- U harfi ile başlayan sayı veritiplerinde ise     sayının değeri pozitif veya negatif işarette   değildir. Sadece bir sayısal değerdir. U’nun anlamı unassigned yani işaretsizdir. Uint8 0-255 arası, uint16 0-65535, uint32 0-42967295 arası, uint64 0-18446744073709551615 arası değerler alabilir. Uintptr ise yazdığınız sayıya göre alanı belirlenir.

### Byte Veri Tipi: uint8 ile aynıdır.
Rune: int32 ile aynıdır. Unicode karakter kodlarını ifade eder.
### Float Türler    
Float türleri integer türlerden farklı olarak küsüratlı sayıları tutar. Örnek: 3.14  
float32: 32bitlik değer alabilir.  
float64: 64 değer alabilir.    
### Complex Türler
Complex türleri içerisinde gerçel küsüratlı (float) ve sanal sayılar barındırabilir. Türkçe’de karmaşık sayılar diye adlandırılır.   
complex64: Gerçel float32 ve sanal sayı değeri barındırır.   
complex128: Gerçel float64 ve sanal sayı değeri barındırır.   
### BOOLEAN VERİ TİPİ      
Boolean yani mantıksal veri tipi bir durumun var olması halinde olumlu (true) değer, var olmaması halinde olumsuz (false) değer alan veri tipidir.
### STRING VERİ TİPİ   
String yani dizgi veri tipi içerisinde metinsel ifadeler barındırır. 
</details>



<details>
<summary>Aritmetik Operatörler :</summary> 

### + Toplar.     
### - Çıkartır.  
### * Çarpar.
### / Böler.  
### % Bölümden Kalanı Verme.  
### ++ 1 Arttırır.  
### -- 1 Eksiltir.  

</details>

<details>
<summary>İlişkisel Operatörler :</summary> 

### == İki verinin eşitliği.     
### != İki verinin eşitsizliği.  
### > 1. verinin 2. veriden büyüklüğü.
### < 1. verinin 2. veriden küçüklüğü.  
### <= 1. verinin 2. veriden küçük veya eşitliği.  
### >= 1. verinin 2. veriden büyük veya eşitliği.  

</details>

<details>
<summary>Mantıksal Operatörler :</summary> 

### && VE operatörüdür. 2 koşulda doğruysa true değer verir.      
### || VEYA operatörüdür. 2 koşuldan ikisi veya biri doğru ise true değer verir.   
### !  DEĞİL operatörüdür. Koşul sonucunun tersini verir.  
</details>

<details>
<summary>Atama Operatörler :</summary> 

### = Atama Operatörüdür.     
### += Kendiyle toplar.  
### -= Kendinden çıkarır.
### *= Kendiyle çarpar.  
### /= Kendine böler.  
### %= Kendine bölümünden kalanı atar.  
</details>

<details>
<summary>Değişkenler ve Atanması :</summary> 

```go
var isim string = “Ali”
var yas int = 20
var ogrenci bool = true
``` 

 #### Değişkenler içerisinde değer barındırarak RAM’e kaydettiğimiz bilgilerdir. Değişkenler programımızın işleyişinde önemli bir role sahiptir. Değişkenleri şu şekillerde atayabiliriz. Değişkenler var ifadesi ile atanır. Tabi ki zorunlu değildir.  

### En basit şekilde değişken ataması yapmak istersek;  

```go
isim:=”Ali”
yas:=20
ogrenci:=true
```

</details>

<details>
<summary>Sabitler :</summary> 

### Sabitler de değişkenler gibi değer alır. Fakat adından da anlaşılabileceği üzere verilen değer daha sonradan değiştirilemez.
  
```go
const isim string = “Ali”
const isim=”Veli”
``` 

</details>

<details>
<summary>Tür Dönüşümü  :  </summary>

### Tür dönüşümü şu şekilde gerçekleştirilir.   
   tür(değer)
### Örnek olarak bakmak gerekir ise;

```go
i := 42
f := float64(i)
u := uint(f)
```
</details>
