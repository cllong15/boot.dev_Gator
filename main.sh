go run . reset
echo "======"
go run . register kahya
echo "======"
go run . addfeed "Hacker News RSS" "https://hnrss.org/newest"
echo "======"
go run . register holgith
echo "======"
go run . addfeed "Lanes Blog" "https://www.wagslane.dev/index.xml"
echo "======"
# go run . addfeed "Boot.dev Blog" "https://www.boot.dev/blog/index.xml"
# echo "======"
# go run . addfeed "TechCrunch" "https://techcrunch.com/feed/"
# echo "======"
# go run . addfeed "Hacker News" "https://news.ycombinator.com/rss"
# echo "======"
go run . follow "https://hnrss.org/newest"
echo "======"
go run . following
echo "======"
go run . login kahya
echo "======"
# go run . follow "https://news.ycombinator.com/rss"
echo "======"
go run . following

# TechCrunch: https://techcrunch.com/feed/
# Hacker News: https://news.ycombinator.com/rss