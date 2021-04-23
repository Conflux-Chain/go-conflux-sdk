cd "$(dirname "$0")"

mv ./interface.go ./interface.go.tmd
mv ./interface_mock.go ./interface_mock.go.tmd

godocdown -o client.md
cd ./utils && godocdown -o ../utils.md && cd ..
cd ./contract_meta/internal_contract && godocdown -o ../../internal_contract.md && cd ../..

cat client.md utils.md internal_contract.md > tmp.md

mv ./interface.go.tmd ./interface.go
mv ./interface_mock.go.tmd ./interface_mock.go 

sed -i "" 's/## Usage//g' tmp.md
sed -i "" 's/#### type/### type/g' tmp.md

rm -f client.md
rm -f utils.md
rm -f internal_contract.md

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