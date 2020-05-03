; (defn swap-once [values]
;   (loop [a (first values)
;          b (second values)
;          new_list '()
;          old_list]
;     (println a b)))
  
(defn- bubble [ys x]
  (if-let [y (peek ys)]
    (if (> y x)
      (conj (pop ys) x y)
      (conj ys x))
    [x]))

(defn bubble-sort [xs]
  (let [ys (reduce bubble [] xs)]
    (if (= xs ys)
      xs
      (recur ys))))
