# The basics

1. Install npm packages `npm install`
2. Run targets against Go projects, for example - `npx nx run go-lib-1:lint`, `npx nx run go-app-1:serve`, `npx nx run go-lib-2:test`
   - Run all the tests - `npx nx run-many --target=test`
3. View the project dependency graph `npx nx graph`

## Create a new Go App

`npx nx g @nx-go/nx-go:app go-app-3`

## Create a new Go Lib

`npx nx g @nx-go/nx-go:lib go-lib-3`
