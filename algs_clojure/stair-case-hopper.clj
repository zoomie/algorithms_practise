(defn f[n]
    (if (< n 0)
      0
      (if (<= n 1)
        1
        (+ (f (- n 1)) (f (- n 2)) (f (- n 3))))))
  (f 10)