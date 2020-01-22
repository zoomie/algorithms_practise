(defn quick-sort [[partition & arr]]
        (when partition
         (concat (quick-sort (filter (fn [x] (< x partition)) arr))
                 [partition]
                 (quick-sort (filter (fn [x] (>= x partition)) arr)))))

(quick-sort '(1 4 2 1))