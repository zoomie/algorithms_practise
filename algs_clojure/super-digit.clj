
(defn str-summer [sum char]
    (let [num (Character/digit char 10)]
      (+ sum num)))
  
  (defn super-digit [num]
    (let [chars (str num)
          len (count chars)]
      (if (>= 1 len)
        num
        (let [sum (reduce str-summer 0 chars)]
          (super-digit sum)))))
  
(super-digit 9875)
  