(def numbers [1 2 3 5 7])

(println "Original:" numbers)

(def squared (map #(* % %) numbers))

(println "Squared Numbers:" squared)


(def numbers [2 3 5 7 11 13])

(def even-numbers (filter even? numbers))

(println "Original:" numbers)
(println "Even Numbers:" even-numbers)


(def numbers [15 25 35 45])

(def total (reduce + numbers))

(println "Numbers:" numbers)


(def numbers [13 23 33 43 53 63])

(println "Original:" numbers)

(println "Take First 3:" (take 3 numbers))

(println "Drop First 3:" (drop 3 numbers))


(def numbers [1 2 3 4 5 6 7 8])

(def result
  (->> numbers
       (filter even?)
       (map #(* % %))
       (take 3)))

(println "Original:" numbers)
(println "Result:" result)


(def list-data '(1 2 3 4 5))
(def vector-data [10 20 30 40 50])

(println "List Result:"
         (->> list-data
              (filter odd?)
              (map inc)))

(println "Vector Result:"
         (->> vector-data
              (map #(* % 2))
              (take 3)))

