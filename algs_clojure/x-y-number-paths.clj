(defn find-ways [[x y]]
    (if (or (= x 0) (= y 0))
      1
      (+ (find-ways [(- x 1) y])
         (find-ways [x (- y 1)]))))
  
(find-ways [12 4])