(defn perms [n]
    (cond 
      (>= 0 n) 0
      (= 5 n) 1
      (= 10 n) 2
      :else (+ (perms (- n 5))
               (perms (- n 10))
               (perms (- n 25)))))