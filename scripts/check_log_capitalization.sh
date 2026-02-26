#!/bin/bash

if grep -rE 'slog\.(Error|Info|Warn|Debug|Fatal|Print|Println|Printf)\(["\"][a-z]' --include="*.go" . 2>/dev/null; then
    echo "❌ as mensagens de registro devem começar com letra maiúscula. foram encontradas mensagens de registro em minúsculas acima"

    exit 1
fi