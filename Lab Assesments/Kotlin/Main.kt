data class Course(
    val id: Int,
    val name: String,
    val credits: Int,
    val fee: Double,
    val instructor: String?
)

data class Student(
    val name: String,
    val email: String?,
    val courses: MutableList<Course> = mutableListOf()
)

fun Course.displayInfo() {
    println("$name - $credits credits - Rs.$fee")
}

fun Student.totalFees(): Double {
    return courses.sumOf { it.fee }
}

class CourseBuilder {
    val courses = mutableListOf<Course>()

    fun course(id: Int, name: String, credits: Int, fee: Double, instructor: String?) {
        courses.add(Course(id, name, credits, fee, instructor))
    }
}

fun courseCatalog(block: CourseBuilder.() -> Unit): List<Course> {
    val builder = CourseBuilder()
    builder.block()
    return builder.courses
}

fun main() {

    val catalog = courseCatalog {
        course(1, "Kotlin", 4, 2500.0, "Rahul")
        course(2, "Java", 3, 2000.0, "Amit")
        course(3, "Python", 4, 3000.0, null)
        course(4, "Database", 2, 1500.0, "Neha")
    }

    val student = Student(
        "Yugesh",
        null,
        mutableListOf(catalog[0], catalog[2])
    )

    println("Student Details")
    println("Name: ${student.name}")
    println("Email: ${student.email ?: "Not provided"}")

    println("\nAvailable Courses:")
    catalog.forEach { it.displayInfo() }

    val expensiveCourses = catalog.filter { it.fee > 2000 }

    println("\nCourses with fees above 2000:")
    expensiveCourses.forEach {
        println(it.name)
    }

    val courseNames = catalog.map { it.name }

    println("\nCourse Names:")
    println(courseNames)

    val totalCredits = catalog.map { it.credits }.reduce { a, b -> a + b }

    println("\nTotal Credits: $totalCredits")

    val groupedCourses = catalog.groupBy { it.credits }

    println("\nCourses Grouped by Credits:")
    groupedCourses.forEach { (credits, courses) ->
        println("$credits Credits: ${courses.map { it.name }}")
    }

    println("\nRegistered Courses:")
    student.courses.forEach { it.displayInfo() }

    println("Total Fees: Rs.${student.totalFees()}")
}
