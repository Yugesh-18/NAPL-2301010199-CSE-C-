(import '(java.util ArrayList HashMap))

(println "Absolute value:" (Math/abs -45))
(println "Power value:" (Math/pow 3 2))
(println "Current time:" (System/currentTimeMillis))

(def fruits (ArrayList.))

(.add fruits "Orange")
(.add fruits "Grapes")
(.add fruits "Pineapple")

(println "\nFruit List:" fruits)
(println "Total fruits:" (.size fruits))
(println "First fruit:" (.get fruits 0))

(def student (HashMap.))

(.put student "Name" "Aman")
(.put student "RollNo" 180)
(.put student "Course" "BTECH")

(println "\nStudent Details:" student)
(println "Student Name:" (.get student "Name"))

(def clojure-fruits (vec fruits))

(println "\nClojure Vector:" clojure-fruits)

(def java-fruits (ArrayList. clojure-fruits))

(println "Java ArrayList:" java-fruits)

(def student-data (into {} student))

(println "Clojure Map:" student-data)
