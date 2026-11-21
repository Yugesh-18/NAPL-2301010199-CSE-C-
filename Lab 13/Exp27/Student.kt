class Student(
    private val rollNo: Int,
    private var name: String,
    private var course: String,
    private var percentage: Double
) {

    fun updateDetails() {

        print("Enter student name: ")
        name = readln()

        print("Enter course: ")
        course = readln()

        print("Enter percentage: ")
        percentage = readln().toDouble()
    }

    fun displayDetails() {
        println("\nStudent Details")
        println("Roll Number: $rollNo")
        println("Name: $name")
        println("Course: $course")
        println("Percentage: $percentage%")
    }
}

fun main() {

    val student1 = Student(101, "Rahul", "BCA", 85.5)
    val student2 = Student(102, "Aman", "BBA", 90.0)

    println("Enter details for Student 1")
    student1.updateDetails()
    student1.displayDetails()

    println("\nEnter details for Student 2")
    student2.updateDetails()
    student2.displayDetails()
}
