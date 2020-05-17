(defn number-elements [k array]
    (loop [i 0
           total (first array)
           larray (rest array)]
        (if (> total k)
            i
            (recur (inc i) (+ total (first larray)) (rest larray)))))

(number-elements 10 '(1 3 4 10 110))