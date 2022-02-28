cd "$(dirname "$0")"

mv ./interface.go ./interface.go.tmd

godocdown -o client.md
cd ./utils && godocdown -o ../utils.md && cd ..
cd ./contract_meta/internal_contract && godocdown -o ../../internal_contract.md && cd ../..
cd ./cfxclient/bulk && godocdown -o ../../bulk.md && cd ../..
cd ./types/cfxaddress && godocdown -o ../../cfxaddress.md && cd ../..
cd ./types/unit && godocdown -o ../../unit.md && cd ../..

cat client.md utils.md internal_contract.md bulk.md cfxaddress.md unit.md> tmp.md

mv ./interface.go.tmd ./interface.go

sed -i "" 's/## Usage//g' tmp.md
sed -i "" 's/#### type/### type/g' tmp.md

rm -f client.md
rm -f utils.md
rm -f internal_contract.md
rm -f bulk.md
rm -f cfxaddress.md
rm -f unit.md

mv ./tmp.md ./api.md

# read -r -p "The api.md will be overwritten, are you sure ? [y/n] " input

# case $input in
#     [yY][eE][sS]|[yY])
# 		echo "Yes"
#         mv ./tmp.md ./api.md
# 		;;

#     [nN][oO]|[nN])
# 		echo "No"
#        	;;

#     *)
# 		echo "Invalid input..."
#         rm ./tmp.md
# 		exit 1
# 		;;
# esac