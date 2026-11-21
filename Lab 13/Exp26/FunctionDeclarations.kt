fun displayDetails(name: String = "Student", age: Int = 20) {
    println("Name: $name")
    println("Age: $age")
}

fun printNumbers(vararg numbers: Int) {
    println("Numbers are:")

    for (num in numbers) {
        println(num)
    }
}

fun calculate(numbers: List<Int>, operation: (Int) -> Int) {
    for (num in numbers) {
        println(operation(num))
    }
}

fun main() {

    println("Default arguments:")
    displayDetails()

    println("\nCustom arguments:")
    displayDetails("Yugesh", 22)

    println("\nNamed arguments:")
    displayDetails(age = 21, name = "Rahul")

    println("\nVararg function:")
    printNumbers(10, 20, 30, 40, 50)

    println("\nHigher-order function:")
    val numbers = listOf(1, 2, 3, 4, 5)

    calculate(numbers) { it * it }
}
