; First attempt
(defn see-not-equal [sublist superlist]
  (if (= (first sublist) (first superlist))
    (if (<= (second sublist) (second superlist))
      true false)
     false))

(defn compare-maps [subset superset]
  (if (<= (count subset) (count superset))
    (let [ordered-sub (into (sorted-map) subset)
          ordered-super (into (sorted-map) superset)]
      (every? identity (map see-not-equal ordered-sub ordered-super)))
    false))

(defn add-to-map [m char]   
  (if (contains? m char)
    (assoc m char (inc (get m char)))
    (assoc m char 1)))

(defn is-substring? [subset superset]
  (let [subset-map (reduce add-to-map {} subset)
        superset-map (reduce add-to-map {} superset)]
    (compare-maps subset-map superset-map)))

(def substring "clojur")
(def superstring"clojure")
(println (is-substring? substring superstring))

; Second attempt
(defn create-dict [input-string]
  (loop [string input-string
        dict {}]
    (if (empty? string)
      dict
      (let [char (first string)]
        (if (contains? dict char)
          (recur (rest string) (assoc dict char (inc (dict char))))
          (recur (rest string) (assoc dict char 1)))))))

(defn compare-strings [outer inner]
  (let [dict-outer (create-dict outer)
        dict-inner (create-dict inner)]
    (for [[key value] dict-inner]
      (if (contains? dict-outer key)
        (if (>= (dict-outer key) value)
          true 
          false)
        false))))

(compare-strings "string" "str")