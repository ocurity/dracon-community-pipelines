orgname="kubernetes"
touch results.txt
gh repo list $orgname --limit 1000 | while read -r repo _; do
 echo $repo:$(GO111MODULE=on go run main.go -apiKey $DT_API_KEY  -name "$repo" -url $DT_URL)>>results.txt
done
