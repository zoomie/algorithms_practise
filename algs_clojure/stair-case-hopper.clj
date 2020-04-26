(defn f [n]
  (if (> 0 n)
    0
    (if (= 0 n)
      1
      (+ (f (- n 3)) (f (- n 2)) (f (- n 1))))))
(f 20)