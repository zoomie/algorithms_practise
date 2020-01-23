(defn make-dict [dict x]
        (if (contains? dict x)
         (assoc dict x (inc (get dict x)))
         (assoc dict x 1)))

(defn count-dict-pairs [num [_ value]]
    (let [new_pairs (int (Math/floor (/ value 2)))]
        (+ num new_pairs)))

(defn count-total-full-pairs [n ar]
   (let [dict (reduce make-dict {} ar)]
        (reduce count-dict-pairs 0 dict)))

(count-total-full-pairs 9 '(10 20 20 10 10 30 50 10 20))
