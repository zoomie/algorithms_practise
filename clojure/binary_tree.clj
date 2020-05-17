(defrecord Node [value left right])

(defn add-value [node value]
  (if (nil? node)
    (Node. value nil nil)
    (if (< (:value node) value)
      (assoc node :right (add-value (:right node) value))
      (assoc node :left (add-value (:left node) value)))))

(-> (add-value nil 5)
    (add-value 3)
    (add-value 10)
    (add-value 12))

