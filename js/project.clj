(defproject bridge.client "0.1"
  :dependencies [[domina "1.0.0"]]
  :cljsbuild
    {:builds 
      [{:source-path "src"
        :compiler {:output-to "out/bridge.js"
                   :output-dir "out"
                   :optimizations nil
                   :pretty-print true}}]})