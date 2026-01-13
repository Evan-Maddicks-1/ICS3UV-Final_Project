# ICS3UV

## GitHub

- ensure you commit your code after every major change, to keep your history and
  not loose anything
  ```console
  git add -A
  git commit -m "𝑐𝑜𝑚𝑚𝑖𝑡 𝑚𝑒𝑠𝑠𝑎𝑔𝑒"
  git push origin main
  ```

## Typescript, with Deno

- to compile your program
  ```console
  deno run main.ts
  ```

- to format your program
  ```console
  deno fmt --check ./*.ts
  ```
- to lint your program
  ```console
  deno lint ./*.ts
  ```

## Go

- to compile your program
  ```console
  go run main.go
  ```
- to format your program
  ```console
  gofmt -d ./*.go
  ```
- to lint your program
  ```console
  golangci-lint run ./*.go
  ```
