app_build_frontend() {
  go generate
}

app_build_backend() {
  go build -o goreact  || return 1
  return 0
}

app_build() {
  app_build_frontend || return 1
  app_build_backend || return 1
  return 0
}

app_build_and_run() {
  echo "rebuilding backend ..."
  app_build_backend || return 1
  echo "starting app ..."
  ./goreact --dev || return 1
  return 0
}

app_find_backend_files() {
  find -name "vendor" -prune \
    -o -name "*.go" -print
}

app_watch_backend() {
  app_find_backend_files | entr -r ./build_and_run.sh || return 1
  return 0
}

app_watch_frontend() {
  (cd frontend && npm run watch) || return 1
  return 0
}