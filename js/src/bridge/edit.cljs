(ns bridge.edit
  (:require [domina :as d])
  (:require-macros [domina.macros :as dm]))

(defn init []
  (-> (sel "body")
    (d/append! "<h2>Edit and send bridge message</h2>")
    (d/append! "<div><h5>Message name</h5></div>")))

