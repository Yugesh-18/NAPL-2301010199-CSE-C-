fun main() {

    print("Enter an integer: ")
    val number = readln().toInt()

    var temp = kotlin.math.abs(number)
    var sum = 0
    var count = 0

    while (temp > 0) {

        val digit = temp % 10
        sum = sum + digit
        count++

        temp = temp / 10
    }

    if (count > 0) {
        val average = sum.toDouble() / count
        println("Sum of digits: $sum")
        println("Average of digits: $average")
    } else {
        println("Sum of digits: 0")
        println("Average of digits: 0.0")
    }

}
