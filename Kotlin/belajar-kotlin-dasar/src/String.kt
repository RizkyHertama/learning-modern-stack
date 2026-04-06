fun main(){

    /*  Di Kotlin, val digunakan untuk variabel immutable (tidak dapat diubah/read-only) yang hanya bisa diinisialisasi sekali,
    setara dengan final di Java. Sebaliknya, var digunakan untuk variabel mutable (dapat diubah) yang nilainya bisa ditetapkan ulang berkali-kali.
    Disarankan menggunakan val secara default untuk keamanan dan prediktabilitas kode */

//    var firstName: String = "Rizky"
//    var lastName: String = "Hertama"

    val firstName: String = "Rizky"
    val lastName: String = "Hertama"
    val address: String = """
       >Jl kedongdong, RT 008 RW 021 no 29, 
       >Kecamatan Epstein,
       >Kelurahan solo,
       >DKI Jakarta 
    """.trimMargin(">") //Trim margin berfungsi untuk merapihkan output berdasarkan tanda baca tertentu

    println(firstName)
    println(lastName)
    println(address)

    val fullName: String = firstName + " " + lastName
    println(fullName)

    val fullName2: String = "$firstName $lastName"
    println(fullName2)

    val desc: String = "$fullName2 length = ${fullName2.length}"
    println(desc)
}