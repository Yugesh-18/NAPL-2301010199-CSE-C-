(def my-list '(10 20 30 40))

(println "Original List:" my-list)

(println "First Element:" (first my-list))
(println "Rest of List:" (rest my-list))

(def new-list (conj my-list 5))

(println "New List:" new-list)
(println "Original List:" my-list)


(def my-vector [10 20 30 40])

(println "Original Vector:" my-vector)

(println "Element at index 1:" (get my-vector 1))
(println "Vector Count:" (count my-vector))


(def new-vector (conj my-vector 50))

(println "New Vector:" new-vector)
(println "Original Vector:" my-vector)

(def my-map {:name "Sagar" :age 21})

(println "Original Map:" my-map)

(def updated-map (assoc my-map :city "Delhi"))

(println "After Assoc:" updated-map)

(def removed-map (dissoc my-map :age))

(println "After Dissoc:" removed-map)

(println "Name:" (get my-map :name))
(println "Original Map:" my-map)


(def my-set #{10 20 30})

(println "Original Set:" my-set)

(def new-set (conj my-set 40))

(println "After Conj:" new-set)

(def removed-set (disj my-set 20))

(println "After Disj:" removed-set)

(println "Contains 10:" (contains? my-set 10))
(println "Original Set:" my-set)