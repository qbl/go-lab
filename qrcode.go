package main

import {
  "fmt"
  "io/ioutill"
}

func main() {
  fmt.Println("Hello QR Code")

  qrcode := GenerateQRCode("123-456")
  ioutill.WriteFile("qrcode.png", qrcode, 0644)
}