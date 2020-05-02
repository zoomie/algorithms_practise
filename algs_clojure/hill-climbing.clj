(defn count-hills [dict char]
    (if (= (get dict :last-step) char)
        (assoc dict :number (inc (get dict :number)))
        (assoc dict :last-step char)))

(defn number-changes [topography]
    (let [first-step (first topography)
          path (rest topography)]
        (reduce count-hills {:number 0 :last-step first-step} path)))

(number-changes "UUUDDDUUU")