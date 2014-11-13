package main

import (
  "net/http"
  "io"
  "os"
)

func main() {
  if len(os.Args) != 3 {
    println("Usage: consul-service <consul-url> <service-name>")
    os.Exit(1)
  }

  consulUrl := "http://" + os.Args[1]
  service := os.Args[2]
  serviceUrl := consulUrl + "/v1/catalog/service/" + service

  println("Requesting: " + consulUrl + "...")

  resp, err := http.Get(serviceUrl)

  if err != nil {
    println(err)
  }

  _, err = io.Copy(os.Stdout, resp.Body)
}
