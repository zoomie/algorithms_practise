
(defn perms [l]
  (let [len (count l)]
    (if (= 1 len)
      (set l)
        (concat 
          [(set l)] 
          (for [index l]
            (perms (apply vector (disj (set l) index))))))))

(defn get-all-perms [l]
  (let [out-list (map 
                  (fn [s] (sort (apply list s)))
                  (apply list (set (flatten (perms l)))))]
    (filter #(< 1 (count %)) out-list)))


(defn make-bools [sub-list]
  (loop [l (rest sub-list)
         last-value (first sub-list)
         list-contains-adj? '()]
    (if (empty? l)
      list-contains-adj?
      (let [current-value (first l)
            diff (- current-value last-value)]
        (recur (rest l)
               current-value
                (conj list-contains-adj?
                      (if (> 2 diff)
                        false
                        true)))))))

(defn contains-adj? [l]
  (let [bools (make-bools l)]
    (every? #(= true %) bools)))
(contains-adj? [1 3])



(defn get-vals-sum [values-list indexs]
  (reduce + (for [i indexs]
              (values-list i))))
(get-vals-sum [1 2 3 4] [0 2])


(defn get-max-pairs [in-value]
  (let [indexs (filter contains-adj? (get-all-perms (range (count in-value))))]
   (loop [index-set indexs
          max-so-far 0
          max-index []]
    (if (empty? index-set)
      {:max-so-far max-so-far :max-index max-index}
      (let [test-max (get-vals-sum in-value (first index-set))]
        test-max
        (if (> test-max max-so-far)
          (recur (rest index-set) test-max (first index-set))
          (recur (rest index-set) max-so-far max-index)))))))

(get-max-pairs [1 5 6 9 30])
