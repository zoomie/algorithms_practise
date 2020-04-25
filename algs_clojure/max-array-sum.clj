(defn f [l]
    (let [len (count l)]
      (if (= 1 len)
        l
        (concat 
          [l] 
          (for [index l]
            (f (apply vector (disj (set l) index))))))))
  
  
  (f [1 2 3])
  