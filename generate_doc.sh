mv ./interface.go ./interface.go.tmd
mv ./interface_mock.go ./interface_mock.go.tmd

godocdown -o tmp1.md
cd ./utils && godocdown -o ../tmp2.md && cd ..
cat tmp1.md tmp2.md > tmp.md

mv ./interface.go.tmd ./interface.go
mv ./interface_mock.go.tmd ./interface_mock.go 

sed -i "" 's/## Usage//g' tmp.md 
sed -i "" 's/#### type/### type/g' tmp.md

rm -f tmp1.md
rm -f tmp2.md

read -r -p "The api.md will be overwritten, are you sure ? [y/n] " input

case $input in
    [yY][eE][sS]|[yY])
		echo "Yes"
        mv ./tmp.md ./api.md
		;;

    [nN][oO]|[nN])
		echo "No"
       	;;

    *)
		echo "Invalid input..."
        rm ./tmp.md
		exit 1
		;;
esac