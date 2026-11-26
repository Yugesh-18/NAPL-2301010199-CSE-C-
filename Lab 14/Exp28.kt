fun main() {

    println("NULL SAFETY IN KOTLIN")

    val name: String = "Rahul"
    var studentName: String? = null

    println("Non-null value: $name")
    println("Nullable value: $studentName")

    studentName = "Aman"
    println("Updated nullable value: $studentName")

    println("\nSAFE CALL OPERATOR")

    var city: String? = "Delhi"
    println("City length: ${city?.length}")

    city = null
    println("City length: ${city?.length}")

    println("\nELVIS OPERATOR")

    var username: String? = null
    val result = username ?: "Guest"
    println("Username: $result")

    username = "Rohit"
    println("Username: ${username ?: "Guest"}")

    println("\nNOT-NULL ASSERTION OPERATOR")

    var message: String? = null

    try {
        println(message!!.length)
    } catch (e: NullPointerException) {
        println("Null value found. Exception handled.")
    }

    println("\nProgram completed.")
}