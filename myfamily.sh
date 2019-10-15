
curl -O -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json 
cat all.json | jq ".[] | select(.id == $HERO_ID ) |  .connections  | .relatives" |sed 's/"//g'  
 


