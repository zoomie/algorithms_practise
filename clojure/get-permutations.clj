(defn get-perms [full-set]
    (-> (for [val full-set]
          (let [perms (disj full-set val)]
            (if (= 1 (count perms))
              perms
              (get-perms perms))))
      (conj full-set)
      (flatten)))
  
  (defn get-permutations [full-set]
    (set (get-perms full-set)))
  
  (get-permutations #{1 2 3 4 5})
  