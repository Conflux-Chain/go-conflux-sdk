package types

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
)

func TestRLPMarshalBlockHeader(t *testing.T) {
	testJson1 := `{"hash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","parentHash":"0xa0c5975f77a557ab65eb1a137de52cd9d9a88f4b36add157ec3c0e2edce1351f","height":"0xf7cf1c","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0x085123da2df1ab4d0af41b99396280ea8f7778048f78bc141118ca1b163d0d75","deferredReceiptsRoot":"0x7976c478fc5ae2d2abe95cd7ef488b439fbac961abedf4a1b0cdee4be6bab27e","deferredLogsBloomHash":"0xd397b3b043d87fcd6fad1291ff0bfd16401c274896d8c63a923727f077b8e0b5","blame":"0x0","transactionsRoot":"0xbf9add52641cdeb9fec7fc8bbacfaf71592c37df34b1f36220003af8797dfb41","epochNumber":"0xf7cf1c","blockNumber":"0xf7cf1c","gasLimit":"0x1c9c380","gasUsed":"0x2ef98","timestamp":"0x60b853f5","difficulty":"0x1371539f68f","powQuality":"0x204fb171a2c","refereeHashes":["0xd28aeb7aea7012a58d776b89f03bbed85d7ebd75e445323e02dcf28af8753750","0xa20f6152fb3434c0c1c3ce8176044476e4baa2db18d0cea9185932e7072fd0ac"],"adaptive":false,"nonce":"0xf1c5c1596190023b","size":"0x24a","custom":[],"posReference":null}`
	testJson2 := `{"hash":"0x11b5c88b4e42fcf95cb1454d5de03d7f31fb59f80df1e49c0723f4f86516ef01","parentHash":"0x372e5820b5f6cd0ffe03c27525694daf28da0ab236cbf961e41aa2e24880bcf7","height":"0xf7cf20","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0x1b04a03817a85ed558e660fb1a0b5b9a640acced757903f8f947cbfbca4cee9a","deferredReceiptsRoot":"0x30ce3b69dadfb10545672f166c953825cccfcb2fb2b6c9c3e205ed2d6f9e8ac1","deferredLogsBloomHash":"0x730d2fa11bef8d14e1f35948a6bdbd09e6ec7fb148ebe2fedae76e5af8cb5d4b","blame":"0x0","transactionsRoot":"0x4bbeac6fa3502f7d2e78eed5caec47ffefad6ec5bb85e84d02f682a05b81de14","epochNumber":"0xf7cf20","blockNumber":"0xf7cf20","gasLimit":"0x1c9c380","gasUsed":"0x4c136","timestamp":"0x60b853f9","difficulty":"0x1371539f68f","powQuality":"0x1db83607fe3","refereeHashes":[],"adaptive":false,"nonce":"0x11f684c0d194b2a3","size":"0x421","custom":[],"posReference":null}`
	testJson3 := `{"hash":"0x11b5c88b4e42fcf95cb1454d5de03d7f31fb59f80df1e49c0723f4f86516ef01","parentHash":"0x372e5820b5f6cd0ffe03c27525694daf28da0ab236cbf961e41aa2e24880bcf7","height":"0xf7cf20","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0x1b04a03817a85ed558e660fb1a0b5b9a640acced757903f8f947cbfbca4cee9a","deferredReceiptsRoot":"0x30ce3b69dadfb10545672f166c953825cccfcb2fb2b6c9c3e205ed2d6f9e8ac1","deferredLogsBloomHash":"0x730d2fa11bef8d14e1f35948a6bdbd09e6ec7fb148ebe2fedae76e5af8cb5d4b","blame":"0x0","transactionsRoot":"0x4bbeac6fa3502f7d2e78eed5caec47ffefad6ec5bb85e84d02f682a05b81de14","epochNumber":"0xf7cf20","blockNumber":"0xf7cf20","gasLimit":"0x1c9c380","gasUsed":"0x4c136","timestamp":"0x60b853f9","difficulty":"0x1371539f68f","powQuality":"0x1db83607fe3","refereeHashes":[],"adaptive":false,"nonce":"0x11f684c0d194b2a3","size":"0x421","custom":["0x0102"],"posReference":null}`
	testJson4 := `{"hash":"0x11b5c88b4e42fcf95cb1454d5de03d7f31fb59f80df1e49c0723f4f86516ef01","parentHash":"0x372e5820b5f6cd0ffe03c27525694daf28da0ab236cbf961e41aa2e24880bcf7","height":"0xf7cf20","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0x1b04a03817a85ed558e660fb1a0b5b9a640acced757903f8f947cbfbca4cee9a","deferredReceiptsRoot":"0x30ce3b69dadfb10545672f166c953825cccfcb2fb2b6c9c3e205ed2d6f9e8ac1","deferredLogsBloomHash":"0x730d2fa11bef8d14e1f35948a6bdbd09e6ec7fb148ebe2fedae76e5af8cb5d4b","blame":"0x0","transactionsRoot":"0x4bbeac6fa3502f7d2e78eed5caec47ffefad6ec5bb85e84d02f682a05b81de14","epochNumber":"0xf7cf20","blockNumber":null,"gasLimit":"0x1c9c380","gasUsed":null,"timestamp":"0x60b853f9","difficulty":"0x1371539f68f","powQuality":"0x1db83607fe3","refereeHashes":[],"adaptive":false,"nonce":"0x11f684c0d194b2a3","size":"0x421","custom":["0x01","0x02"],"posReference":null}`

	for _, bhJson := range []string{testJson4, testJson3, testJson2, testJson1} {
		var bh BlockHeader
		err := json.Unmarshal([]byte(bhJson), &bh)
		fatalIfErr(t, err)
		// RLP marshal block header to bytes
		dBytes, err := rlp.EncodeToBytes(bh)
		fatalIfErr(t, err)
		// RLP unmarshal bytes back to block header
		var bh2 BlockHeader
		err = rlp.DecodeBytes(dBytes, &bh2)
		fatalIfErr(t, err)

		// Json marshal
		jBytes1, err := json.Marshal(bh)
		fatalIfErr(t, err)
		jsonStr := string(jBytes1)
		// Json marshal
		jBytes2, err := json.Marshal(bh2)
		fatalIfErr(t, err)
		jsonStr2 := string(jBytes2)

		if jsonStr != jsonStr2 {
			t.Fatalf("expect %v, actual %v", jsonStr, jsonStr2)
		}

		// if !reflect.DeepEqual(bh, bh2) {
		// 	t.Fatalf("expect %#v, \nactual %#v", bh, bh2)
		// }
	}
}

func TestRLPMarshalBlock(t *testing.T) {
	testJsonStr1 := `{"hash":"0xa6528367a9287ed3a66fc64457db15e2aaa93104a3fd06d4f0a2beb6cc1f26c8","parentHash":"0x5aef321e4e49f430ad6322af8a0133eae83e635f7893c996eb127dcf24a00b14","height":"0x792776","miner":"cfx:aatxetsp0kdarpdb5stdyex11dr3x6sb0jw2gykec0","deferredStateRoot":"0xa979a8c492c44a512aa9529911a7862e1b61ce2aa441645e865def9219d2c68b","deferredReceiptsRoot":"0xd5f7e7960e9b56753868260c280746c01353dcd1b91a20cee2c919d0dc7bf78b","deferredLogsBloomHash":"0xd397b3b043d87fcd6fad1291ff0bfd16401c274896d8c63a923727f077b8e0b5","blame":"0x0","transactionsRoot":"0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470","epochNumber":"0x792778","blockNumber":"0xf7cf1c","gasLimit":"0x1c9c380","gasUsed":"0x0","timestamp":"0x6026478e","difficulty":"0xa8b175a4dc","powQuality":"0x2223e36adc5","refereeHashes":["0x4e4fca2593068b1dc83ecae3c1eaf0e4d41623985fd03d7f15fc1d63f653e7d2"],"adaptive":false,"nonce":"0x209fc5fbe719dace","size":"0x0","custom":[],"posReference":null,"transactions":[]}`
	testJsonStr2 := `{"hash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","parentHash":"0xa0c5975f77a557ab65eb1a137de52cd9d9a88f4b36add157ec3c0e2edce1351f","height":"0xf7cf1c","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0x085123da2df1ab4d0af41b99396280ea8f7778048f78bc141118ca1b163d0d75","deferredReceiptsRoot":"0x7976c478fc5ae2d2abe95cd7ef488b439fbac961abedf4a1b0cdee4be6bab27e","deferredLogsBloomHash":"0xd397b3b043d87fcd6fad1291ff0bfd16401c274896d8c63a923727f077b8e0b5","blame":"0x0","transactionsRoot":"0xbf9add52641cdeb9fec7fc8bbacfaf71592c37df34b1f36220003af8797dfb41","epochNumber":"0xf7cf1c","blockNumber":"0xf7cf20","gasLimit":"0x1c9c380","gasUsed":"0x2ef98","timestamp":"0x60b853f5","difficulty":"0x1371539f68f","powQuality":"0x204fb171a2c","refereeHashes":["0xd28aeb7aea7012a58d776b89f03bbed85d7ebd75e445323e02dcf28af8753750","0xa20f6152fb3434c0c1c3ce8176044476e4baa2db18d0cea9185932e7072fd0ac"],"adaptive":false,"nonce":"0xf1c5c1596190023b","size":"0x24a","custom":[],"posReference":null,"transactions":[{"hash":"0x740b71de5591fe87bf661d5c4a39cf2a1fbf8cf21b68928033a7382154c78d19","nonce":"0xe033d","blockHash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","transactionIndex":"0x0","from":"cfx:aapkcjr28dg976fzr43c5hf1rwn5xv8t1uy4r2yyeu","to":"cfx:aan4cta4wa1av51anm00844cg01hf19zw2vufsvk4n","value":"0x3bd913e6c1df4000","gasPrice":"0xa","gas":"0x5208","contractCreated":null,"data":"0x","storageLimit":"0x0","epochHeight":"0xf7cf19","chainId":"0x405","status":"0x0","v":"0x0","r":"0x57199fc1aced9c518fb8f2a31978c434fd280173602f38002b7252c02573cb00","s":"0x3d0ef3327d226d46ed0a53aa4f5486666aab2a9423a8261b9456975ba1fde346"},{"hash":"0x536cb069dc9024625c3ae27fef0a32df6733bec0a159f1b4871741b73b0419cb","nonce":"0xe033e","blockHash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","transactionIndex":"0x1","from":"cfx:aapkcjr28dg976fzr43c5hf1rwn5xv8t1uy4r2yyeu","to":"cfx:aam9ee2tdeajz4akcx3nyhgnmu7zw5hd2u0nnce7j5","value":"0x3ab4b07cc8db8000","gasPrice":"0xa","gas":"0x5208","contractCreated":null,"data":"0x","storageLimit":"0x0","epochHeight":"0xf7cf19","chainId":"0x405","status":"0x0","v":"0x0","r":"0x7b5ff9bb9a93bf9d974102a22a6323c86558c1cc776ad93af5f2f1df0f1c9918","s":"0x6aaddd6e3c988bd6b02b37cd8f211a26f6ecd759fb8744f5233a3dfcd2d10183"},{"hash":"0x18e0f546df2f56e8149bb9424d70201feb41f8e2ea7c1a850a03b8a2f508a8f7","nonce":"0x1063","blockHash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","transactionIndex":"0x2","from":"cfx:aakycdtv194ws11y73tsam7va08ex0ztmjrjaxuds3","to":"cfx:acam64yj323zd4t1fhybxh3jsg7hu4012yz9kakxs9","value":"0x3635c9adc5dea00000","gasPrice":"0x1","gas":"0x28e04","contractCreated":null,"data":"0x5c350838000000000000000000000000000000000000000000000017931cda8bd511fb00000000000000000000000000000000000000000000000000000000000000008000000000000000000000000013410df1bff5275ef4ee5ee02bb105bc49daaf520000000000000000000000000000000000000000000000000000000060b9a56e00000000000000000000000000000000000000000000000000000000000000020000000000000000000000008d7df9316faa0586e175b5e6d03c6bda76e3d9500000000000000000000000008b8689c7f3014a4d86e4d1d0daaf74a47f5e0f27","storageLimit":"0x1cc","epochHeight":"0xf7cf19","chainId":"0x405","status":"0x0","v":"0x0","r":"0x2e2237b2baf72e80725d0cbd5aee2dd755a606c905e77e45883b2ed8998523ab","s":"0x3f3be41bbfab7d4aca6d8c45887882f61b2ed0cc8f00a4ab2ab197eaad2a8c70"}]}`
	testJsonStr3 := `{"hash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","parentHash":"0xa0c5975f77a557ab65eb1a137de52cd9d9a88f4b36add157ec3c0e2edce1351f","height":"0xf7cf1c","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0x085123da2df1ab4d0af41b99396280ea8f7778048f78bc141118ca1b163d0d75","deferredReceiptsRoot":"0x7976c478fc5ae2d2abe95cd7ef488b439fbac961abedf4a1b0cdee4be6bab27e","deferredLogsBloomHash":"0xd397b3b043d87fcd6fad1291ff0bfd16401c274896d8c63a923727f077b8e0b5","blame":"0x0","transactionsRoot":"0xbf9add52641cdeb9fec7fc8bbacfaf71592c37df34b1f36220003af8797dfb41","epochNumber":"0xf7cf1c","blockNumber":"0xf7cf20","gasLimit":"0x1c9c380","gasUsed":"0x2ef98","timestamp":"0x60b853f5","difficulty":"0x1371539f68f","powQuality":"0x204fb171a2c","refereeHashes":["0xd28aeb7aea7012a58d776b89f03bbed85d7ebd75e445323e02dcf28af8753750","0xa20f6152fb3434c0c1c3ce8176044476e4baa2db18d0cea9185932e7072fd0ac"],"adaptive":false,"nonce":"0xf1c5c1596190023b","size":"0x24a","custom":["0x0102","0x0203"],"posReference":null,"transactions":[{"hash":"0x740b71de5591fe87bf661d5c4a39cf2a1fbf8cf21b68928033a7382154c78d19","nonce":"0xe033d","blockHash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","transactionIndex":"0x0","from":"cfx:aapkcjr28dg976fzr43c5hf1rwn5xv8t1uy4r2yyeu","to":"cfx:aan4cta4wa1av51anm00844cg01hf19zw2vufsvk4n","value":"0x3bd913e6c1df4000","gasPrice":"0xa","gas":"0x5208","contractCreated":null,"data":"0x","storageLimit":"0x0","epochHeight":"0xf7cf19","chainId":"0x405","status":"0x0","v":"0x0","r":"0x57199fc1aced9c518fb8f2a31978c434fd280173602f38002b7252c02573cb00","s":"0x3d0ef3327d226d46ed0a53aa4f5486666aab2a9423a8261b9456975ba1fde346"},{"hash":"0x536cb069dc9024625c3ae27fef0a32df6733bec0a159f1b4871741b73b0419cb","nonce":"0xe033e","blockHash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","transactionIndex":"0x1","from":"cfx:aapkcjr28dg976fzr43c5hf1rwn5xv8t1uy4r2yyeu","to":"cfx:aam9ee2tdeajz4akcx3nyhgnmu7zw5hd2u0nnce7j5","value":"0x3ab4b07cc8db8000","gasPrice":"0xa","gas":"0x5208","contractCreated":null,"data":"0x","storageLimit":"0x0","epochHeight":"0xf7cf19","chainId":"0x405","status":"0x0","v":"0x0","r":"0x7b5ff9bb9a93bf9d974102a22a6323c86558c1cc776ad93af5f2f1df0f1c9918","s":"0x6aaddd6e3c988bd6b02b37cd8f211a26f6ecd759fb8744f5233a3dfcd2d10183"},{"hash":"0x18e0f546df2f56e8149bb9424d70201feb41f8e2ea7c1a850a03b8a2f508a8f7","nonce":"0x1063","blockHash":"0x26f15dc6f353485cdfb1b370becc4abfdacbd36e39c3f9f42be724fe4073cfeb","transactionIndex":"0x2","from":"cfx:aakycdtv194ws11y73tsam7va08ex0ztmjrjaxuds3","to":"cfx:acam64yj323zd4t1fhybxh3jsg7hu4012yz9kakxs9","value":"0x3635c9adc5dea00000","gasPrice":"0x1","gas":"0x28e04","contractCreated":null,"data":"0x5c350838000000000000000000000000000000000000000000000017931cda8bd511fb00000000000000000000000000000000000000000000000000000000000000008000000000000000000000000013410df1bff5275ef4ee5ee02bb105bc49daaf520000000000000000000000000000000000000000000000000000000060b9a56e00000000000000000000000000000000000000000000000000000000000000020000000000000000000000008d7df9316faa0586e175b5e6d03c6bda76e3d9500000000000000000000000008b8689c7f3014a4d86e4d1d0daaf74a47f5e0f27","storageLimit":"0x1cc","epochHeight":"0xf7cf19","chainId":"0x405","status":"0x0","v":"0x0","r":"0x2e2237b2baf72e80725d0cbd5aee2dd755a606c905e77e45883b2ed8998523ab","s":"0x3f3be41bbfab7d4aca6d8c45887882f61b2ed0cc8f00a4ab2ab197eaad2a8c70"}]}`

	for _, blockJson := range []string{testJsonStr3, testJsonStr2, testJsonStr1} {
		var block Block
		err := json.Unmarshal([]byte(blockJson), &block)
		fatalIfErr(t, err)
		// RLP marshal block to bytes
		dBytes, err := rlp.EncodeToBytes(block)
		fatalIfErr(t, err)
		// RLP unmarshal bytes back to block
		var block2 Block
		err = rlp.DecodeBytes(dBytes, &block2)
		fatalIfErr(t, err)
		// // Json marshal block
		jBytes1, err := json.Marshal(block)
		fatalIfErr(t, err)
		blockJsonStr := string(jBytes1)
		// Json marshal block2
		jBytes2, err := json.Marshal(block2)
		fatalIfErr(t, err)
		blockJsonStr2 := string(jBytes2)

		if blockJsonStr != blockJsonStr2 {
			t.Fatalf("expect %v, actual %v", blockJsonStr, blockJsonStr2)
		}
	}
}

func TestRLPMarshalBlockSummary(t *testing.T) {
	testJson1 := `{"hash":"0xa6528367a9287ed3a66fc64457db15e2aaa93104a3fd06d4f0a2beb6cc1f26c8","parentHash":"0x5aef321e4e49f430ad6322af8a0133eae83e635f7893c996eb127dcf24a00b14","height":"0x792776","miner":"cfx:aatxetsp0kdarpdb5stdyex11dr3x6sb0jw2gykec0","deferredStateRoot":"0xa979a8c492c44a512aa9529911a7862e1b61ce2aa441645e865def9219d2c68b","deferredReceiptsRoot":"0xd5f7e7960e9b56753868260c280746c01353dcd1b91a20cee2c919d0dc7bf78b","deferredLogsBloomHash":"0xd397b3b043d87fcd6fad1291ff0bfd16401c274896d8c63a923727f077b8e0b5","blame":"0x0","transactionsRoot":"0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470","epochNumber":"0x792778","blockNumber":"0xf7cf1c","gasLimit":"0x1c9c380","gasUsed":"0x0","timestamp":"0x6026478e","difficulty":"0xa8b175a4dc","powQuality":"0x2223e36adc5","refereeHashes":["0x4e4fca2593068b1dc83ecae3c1eaf0e4d41623985fd03d7f15fc1d63f653e7d2"],"adaptive":false,"nonce":"0x209fc5fbe719dace","size":"0x0","custom":[],"posReference":null,"transactions":[]}`
	testJson2 := `{"hash":"0xf2855662d53e36d32bece4e8a3ac7bd8368dd721c8b3f1749fcabc0d1c25d698","parentHash":"0xf31040dd3e47210a6efc006f75e6517dc602597b20328b4ef2310a9b0efeb005","height":"0x79279d","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0xb60444c84c14210b127e9c50a6df7de0c64f318d39950d6b90535f12576909d8","deferredReceiptsRoot":"0xb5ebf9aa3f401b3d1c189bffd746ca2819bafb623fe853c1d3b7eb219de01929","deferredLogsBloomHash":"0x57c769a4c976741eb0702f31f8be891c3d1a7415dd18999d6f93d38d0768d34a","blame":"0x0","transactionsRoot":"0x45ed3eda85877aa9f7fdf18808db522b6303c2a5e35ce5c565583e6d52790ec0","epochNumber":"0x79279d","blockNumber":"0xf7cf20","gasLimit":"0x1c9c380","gasUsed":"0x42c22","timestamp":"0x602647d5","difficulty":"0xa8b175a4dc","powQuality":"0x117dfd6ba32","refereeHashes":[],"adaptive":false,"nonce":"0x8995446ca72b2d56","size":"0x47d","custom":[],"posReference":null,"transactions":["0xbc75e11b03eec00d1129134af9e568e4687b53d1d80aaa745bbeb77bbb5c6ab0","0x2f1b8dde4bffb052c70ee082404171cf2399aeee8e46f09cc58d44d4b47ef84c","0x4af03a966b291eb344044fa513f1060bc1ca7f5ec170819044973c95bc7b584f","0x689ae502a41ebeff360549cea025b556a6fe160dd7dd49d9646895cf0697e33e"]}`
	testJson3 := `{"hash":"0xf2855662d53e36d32bece4e8a3ac7bd8368dd721c8b3f1749fcabc0d1c25d698","parentHash":"0xf31040dd3e47210a6efc006f75e6517dc602597b20328b4ef2310a9b0efeb005","height":"0x79279d","miner":"cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5","deferredStateRoot":"0xb60444c84c14210b127e9c50a6df7de0c64f318d39950d6b90535f12576909d8","deferredReceiptsRoot":"0xb5ebf9aa3f401b3d1c189bffd746ca2819bafb623fe853c1d3b7eb219de01929","deferredLogsBloomHash":"0x57c769a4c976741eb0702f31f8be891c3d1a7415dd18999d6f93d38d0768d34a","blame":"0x0","transactionsRoot":"0x45ed3eda85877aa9f7fdf18808db522b6303c2a5e35ce5c565583e6d52790ec0","epochNumber":"0x79279d","blockNumber":"0xf7cf20","gasLimit":"0x1c9c380","gasUsed":"0x42c22","timestamp":"0x602647d5","difficulty":"0xa8b175a4dc","powQuality":"0x117dfd6ba32","refereeHashes":[],"adaptive":false,"nonce":"0x8995446ca72b2d56","size":"0x47d","custom":["0x01","0x0203"],"posReference":null,"transactions":["0xbc75e11b03eec00d1129134af9e568e4687b53d1d80aaa745bbeb77bbb5c6ab0","0x2f1b8dde4bffb052c70ee082404171cf2399aeee8e46f09cc58d44d4b47ef84c","0x4af03a966b291eb344044fa513f1060bc1ca7f5ec170819044973c95bc7b584f","0x689ae502a41ebeff360549cea025b556a6fe160dd7dd49d9646895cf0697e33e"]}`

	for _, bsJson := range []string{testJson3, testJson2, testJson1} {
		var bs BlockSummary
		err := json.Unmarshal([]byte(bsJson), &bs)
		fatalIfErr(t, err)
		// RLP marshal block summary to bytes
		dBytes, err := rlp.EncodeToBytes(bs)
		fatalIfErr(t, err)
		// RLP unmarshal bytes back to block summary
		var bs2 BlockSummary
		err = rlp.DecodeBytes(dBytes, &bs2)
		fatalIfErr(t, err)
		// Json marshal bs
		jBytes1, err := json.Marshal(bs)
		fatalIfErr(t, err)
		bsJsonStr := string(jBytes1)
		// Json marshal bs2
		jBytes2, err := json.Marshal(bs2)
		fatalIfErr(t, err)
		bsJsonStr2 := string(jBytes2)

		if bsJsonStr != bsJsonStr2 {
			t.Fatalf("expect %v, actual %v", bsJsonStr, bsJsonStr2)
		}
	}
}

func TestJsonMarhsalBlock(t *testing.T) {
	jsons := []string{
		`{"hash":"0x6720dc2e79931b4727289612cb3a8ead65f65a3cd1d079348601c52116713b73","parentHash":"0x9154c74219f6e556372e252cf2dd9e675ebcf1a7c257e31cf5a8c3cd79c336f5","height":"0x22","miner":"net8888:aajaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaph895n0mm","deferredStateRoot":"0x223ca6a2e599d487f739a0586bf26c88b8032f10f7607cb4e623d100675ef3b0","deferredReceiptsRoot":"0x09f8709ea9f344a810811a373b30861568f5686e649d6177fd92ea2db7477508","deferredLogsBloomHash":"0xd397b3b043d87fcd6fad1291ff0bfd16401c274896d8c63a923727f077b8e0b5","blame":"0x0","transactionsRoot":"0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470","epochNumber":"0x22","blockNumber":"0x22","gasLimit":"0x1c9c380","gasUsed":"0x0","timestamp":"0x619daead","difficulty":"0x1f4","powQuality":"0x1f6","refereeHashes":[],"adaptive":false,"nonce":"0x865a421661e90c0e","size":"0x0","custom":[],"posReference":null,"transactions":[]}`,
		`{"hash":"0x6720dc2e79931b4727289612cb3a8ead65f65a3cd1d079348601c52116713b73","parentHash":"0x9154c74219f6e556372e252cf2dd9e675ebcf1a7c257e31cf5a8c3cd79c336f5","height":"0x22","miner":"net8888:aajaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaph895n0mm","deferredStateRoot":"0x223ca6a2e599d487f739a0586bf26c88b8032f10f7607cb4e623d100675ef3b0","deferredReceiptsRoot":"0x09f8709ea9f344a810811a373b30861568f5686e649d6177fd92ea2db7477508","deferredLogsBloomHash":"0xd397b3b043d87fcd6fad1291ff0bfd16401c274896d8c63a923727f077b8e0b5","blame":"0x0","transactionsRoot":"0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470","epochNumber":"0x22","blockNumber":"0x22","gasLimit":"0x1c9c380","gasUsed":"0x0","timestamp":"0x619daead","difficulty":"0x1f4","powQuality":"0x1f6","refereeHashes":[],"adaptive":false,"nonce":"0x865a421661e90c0e","size":"0x0","custom":["0x0102"],"posReference":null,"transactions":[]}`,
	}
	for _, j := range jsons {
		var b Block
		e := json.Unmarshal([]byte(j), &b)
		if e != nil {
			t.Fatal(e)
		}
		// fmt.Printf("%+v\n", b)
		d, e := json.Marshal(b)
		if e != nil {
			t.Fatal(e)
		}
		if string(d) != j {
			t.Fatalf("expect %v, actual %v", j, string(d))
		}
	}
}
