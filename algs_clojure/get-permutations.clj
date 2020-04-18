(defn get-permutations [full-set]
    (if (= 1 (count full-set))
      full-set
      (for [subset full-set]
        (disj full-set subset))))
  
  (get-permutations #{1 2 4})
