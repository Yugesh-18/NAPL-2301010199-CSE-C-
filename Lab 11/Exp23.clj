(defn check-age [age]
  (if (< age 18)
    (throw (ex-info "Age must be 18 or above"
                    {:given-age age :required-age 18}))
    "Age is valid"))

(println "1. Arithmetic Exception")

(try
  (println (/ 10 0))
  (catch ArithmeticException e
    (println "Error:" (.getMessage e)))
  (finally
    (println "Arithmetic check completed")))

(println "\n2. Null Pointer Exception")

(try
  (let [name nil]
    (.length name))
  (catch NullPointerException e
    (println "Error: Null value found"))
  (finally
    (println "Null check completed")))

(println "\n3. Custom Exception")

(try
  (println (check-age 16))
  (catch clojure.lang.ExceptionInfo e
    (println "Error:" (ex-message e))
    (println "Details:" (ex-data e)))
  (finally
    (println "Age check completed")))

(println "\nProgram finished successfully")
