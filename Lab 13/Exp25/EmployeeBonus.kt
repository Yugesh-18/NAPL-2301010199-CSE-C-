fun main() {

    val deptCode = readln().uppercase()
    val experience = readln().toInt()
    val salary = readln().toDouble()
    val bonusPercent = when {
        deptCode == "IT" && experience > 5 -> 20
        deptCode == "HR" && experience > 3 -> 15
        else -> 10
    }

    val bonusAmount = salary * bonusPercent / 100

    println("Department: $deptCode")
    println("Bonus Percentage: $bonusPercent%")
    println("Bonus Amount: Rs. $bonusAmount")

}
