while true
do
    watchman-wait -p "**/*.go" -- .
    eval "clear"
    eval "go test -p 1 -count=1 -timeout 30s -cover ./..."
    echo "-- finished --"
done
