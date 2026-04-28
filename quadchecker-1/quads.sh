#!/bin/bash

set -e

SRC="main.go"

FUNCS=("quadA" "quadB" "quadC" "quadD" "quadE")

echo "Building quadchecker..."
go build -o quadchecker "$SRC"

for FUNC in "${FUNCS[@]}"; do
    OUT="${FUNC}.go"

    echo "Creating $OUT..."

    # Extract function
    awk "/func $FUNC\\(/,/^}/" "$SRC" > tmp_func.go

    # Create full Go file
    cat > "$OUT" <<EOF
package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

$(cat tmp_func.go)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: ./$FUNC <width> <height>")
        return
    }

    x, _ := strconv.Atoi(os.Args[1])
    y, _ := strconv.Atoi(os.Args[2])

    fmt.Print($FUNC(x, y))
}
EOF

    # Build executable
    go build -o "$FUNC" "$OUT"

    # Remove generated .go file
    #rm -f "$OUT"
done

# Cleanup temp file
rm -f tmp_func.go

echo "Done. Executables created:"
ls -1 quadchecker quadA quadB quadC quadD quadE
