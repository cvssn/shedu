ISSUES=$(gh issue list --state=all --limit=1000 --json "number" -t '{{range .}}{{printf "%.0f\n" .number}}{{end}}')
PRS=$(gh pr list --state=all --limit=1000 --json "number" -t '{{range .}}{{printf "%.0f\n" .number}}{{end}}')

for issue in $ISSUES; do
    echo "enviando labeler.yml para $issue"

    gh workflow run labeler.yml -f issue-number="$issue"
done

for pr in $PRS; do
    echo "enviando labeler.yml para $pr"

    gh workflow run labeler.yml -f issue-number="$pr"
done