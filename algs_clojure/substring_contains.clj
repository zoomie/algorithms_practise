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