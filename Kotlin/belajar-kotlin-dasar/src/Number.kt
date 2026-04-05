fun main(){

    //Untuk tipe data integer banyak variasinya tergantung banyaknya digit 0 nya
    println("=== INTEGER ===")
    var age: Byte = 30
    var height: Int = 170
    var distance: Short = 2000
    var balance: Long = 100000000L

    println(age)
    println(height)
    println(distance)
    println(balance)
    println()

    //Seringkali yang digunakan dalam real case kerja adalah double. Untuk float harus diakhir F diakhir angka
    println("=== DECIMAL ===")
    var value: Float = 98.87F
    var radius: Double = 2312131312.64324

    println(value)
    println(radius)
    println()

    //Seringkali digunakan untuk color pallete
    println("=== LITERALS ===")
    var decimalLiteral: Int = 100
    var hexadecimalLiteral: Int = 0xFF
    var binaLiteralL: Int = 0b0001

    println(decimalLiteral)
    println(hexadecimalLiteral)
    println(binaLiteralL)
    println()

    //Fitur underscore ga ada efek apapun hanya untuk mempermudah membaca nominal
    println("=== UNDERSCORE ===")
    var age_underscore: Byte = 3_0
    var height_underscore: Int = 1_7_0
    var distance_underscore: Short = 2_000
    var balance_underscore: Long = 100_000_000L

    println(age_underscore)
    println(height_underscore)
    println(distance_underscore)
    println(balance_underscore)
    println()

    //Konversi bisa mengubah tipe data satu dengan lainnya
    println("=== CONVERSION ===")
    var number: Int = 100
    var byte_convert: Byte = number.toByte()
    var short_convert: Short = number.toShort()
    var int_convert: Int = number.toInt()
    var long_convert: Long = number.toLong()
    var float_convert: Float = number.toFloat()
    var double_convert: Double = number.toDouble()

    println(byte_convert)
    println(short_convert)
    println(int_convert)
    println(long_convert)
    println(float_convert)
    println(double_convert)
    println()



}
