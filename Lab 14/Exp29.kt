data class Student(val name: String, val percentage: Double)

fun main() {

    println("LIST OPERATIONS")

    val list = listOf(10, 20, 30, 40)
    println("Read-only list: $list")
    println("Contains 20: ${list.contains(20)}")
    println("Element at index 2: ${list[2]}")

    val numbers = mutableListOf(10, 20, 30, 40)
    numbers.add(50)
    numbers.remove(20)
    println("Mutable list: $numbers")

    println("\nSET OPERATIONS")

    val set = setOf("Apple", "Mango", "Banana")
    println("Read-only set: $set")
    println("Contains Mango: ${set.contains("Mango")}")

    val fruits = mutableSetOf("Apple", "Mango", "Banana")
    fruits.add("Orange")
    fruits.remove("Apple")
    println("Mutable set: $fruits")

    println("\nMAP OPERATIONS")

    val map = mapOf(1 to "Rahul", 2 to "Aman", 3 to "Rohit")
    println("Read-only map: $map")
    println("Value of key 2: ${map[2]}")

    val students = mutableMapOf(1 to "Rahul", 2 to "Aman")
    students[3] = "Rohit"
    students.remove(1)
    println("Mutable map: $students")

    println("\nFUNCTIONAL TRANSFORMATIONS")

    val numbersList = listOf(1, 2, 3, 4, 5, 6)

    val doubled = numbersList.map { it * 2 }
    println("After map: $doubled")

    val evenNumbers = numbersList.filter { it % 2 == 0 }
    println("After filter: $evenNumbers")

    val sum = numbersList.reduce { a, b -> a + b }
    println("After reduce: $sum")

    val names = listOf("Rahul", "Aman", "Rohit", "Priya")

    val capitalNames = names.map { it.uppercase() }
    println("Uppercase names: $capitalNames")

    val filteredNames = names.filter { it.length > 4 }
    println("Filtered names: $filteredNames")

    val allNames = names.reduce { a, b -> "$a $b" }
    println("Combined names: $allNames")

    println("\nSORTING AND GROUPING")

    val studentList = listOf(
        Student("Rahul", 85.0),
        Student("Aman", 72.0),
        Student("Rohit", 35.0),
        Student("Priya", 91.0),
        Student("Neha", 42.0)
    )

    val sortedList = studentList.sortedByDescending { it.percentage }

    println("Students sorted by percentage:")
    for (student in sortedList) {
        println("${student.name} - ${student.percentage}%")
    }

    val groupedList = studentList.groupBy {
        if (it.percentage >= 40) "Pass" else "Fail"
    }

    println("\nStudents grouped by result:")

    for ((result, students) in groupedList) {
        println("$result: ${students.map { it.name }}")
    }
}
