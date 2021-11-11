// TODO: create a web page tool

const typeMap = {
    "u64": "*hexutil.Uint64",
    "U64": "*hexutil.Uint64",
    "U256": "*hexutil.Big",
    "H256": "types.Hash",
    "RpcAddress": "types.Address",
}

function convertStruct(struct) {
    const any = "[\\s\\S]"
    const regex = new RegExp(`(pub struct .*)\\{(${any}*)\\}`, "i")
    let [, head, body] = struct.match(regex)

    const goHead = head.replace(/pub struct(.*)/, "type $1 struct")
    const goBody = body.replaceAll(new RegExp(`${any}*?pub (.*)?: (.*)?\\,`, "img"),
        (match, name, type) => `${toAllCamel(name)} ${typeMap[type]} \`json:"${toCamel(name)}"\`\n`)
    return `${goHead}{\n${goBody}}`
}

function convertFunction(nameSpace, func) {
    const any = "[\\s\\S]"
    const regex = new RegExp(`(${any}*)?#\\[rpc.*?"(.*)?"\\)\\]${any}*?fn(${any}*)?\\(\\s*\\&self,*(${any}*)?\\)\\s*->\\s*.*?<(.*)?>`, "img")

    let goFunc = func.replaceAll(regex, (m, comment, rpc, funName, args, returnType) => {
        args = convertArgs(args)
        return `${comment || ''}func(c *Rpc${nameSpace}Client) ${toAllCamel(funName).replace(nameSpace, "")}(${args})(val ${convertType(returnType)}, err error) {
        err = c.core.CallRPC(&val, "${rpc}")
        return  
    }`})
    return goFunc
}

function convertTrait2Funcs(rsustTrait, nameSpace) {
    return rsustTrait
        .replace(/^.*?pub trait.*?\{/igs, "")
        .split(";")
        .map(line => convertFunction(nameSpace, line))
        .join("\n")
}

function convertType(returnType) {
    if (!/Option.*/.test(returnType)) {
        return typeMap[returnType] || returnType
    }

    return returnType
        .replace(/Option<(.*)>/, (m, c) => `*${typeMap[c] || c}`)
        .replace("**", "*")
}

function convertArgs(args) {
    if (!args) return ""

    const argReplaceFunc = (m, name, type) => `${toCamel(name)} ${convertType(type)}`
    return args.replace("\n", "")
        .replace(/\s/g, '')
        .replace(/,$/g, "")
        .split(",")
        .map(arg => arg.replace(/(.*)?:(.*)/g, argReplaceFunc)).join(", ")
}


function toAllCamel(str) {
    let camel = toCamel(str)
    return camel.charAt(0).toUpperCase() + camel.substr(1)
}

function toCamel(str) {
    return str.trim().replace(/[-_](\w)/g, (m, c) => c.toUpperCase())
}

module.exports = {
    convertStruct,
    convertFunction,
    convertTrait2Funcs
}

/** Demo
function run() {
    const rust = `// Copyright 2020 Conflux Foundation. All rights reserved.
    // Conflux is free software and distributed under GNU General Public License.
    // See http://www.gnu.org/licenses/
    
    use crate::rpc::types::{
        AccountPendingInfo, AccountPendingTransactions, RpcAddress,
        Transaction as RpcTransaction, TxPoolPendingNonceRange, TxPoolStatus,
        TxWithPoolInfo,
    };
    use cfx_types::{H256, U256, U64};
    use jsonrpc_core::{BoxFuture, Result as JsonRpcResult};
    use jsonrpc_derive::rpc;
    
    /// Transaction pool RPCs
    #[rpc(server)]
    pub trait TransactionPool {
        #[rpc(name = "txpool_status")]
        fn txpool_status(&self) -> JsonRpcResult<TxPoolStatus>;
    
        #[rpc(name = "txpool_nextNonce")]
        fn txpool_next_nonce(&self, address: RpcAddress) -> JsonRpcResult<U256>;
    
        #[rpc(name = "txpool_transactionByAddressAndNonce")]
        fn txpool_transaction_by_address_and_nonce(
            &self, address: RpcAddress, nonce: U256,
        ) -> JsonRpcResult<Option<RpcTransaction>>;
    
        #[rpc(name = "txpool_pendingNonceRange")]
        fn txpool_pending_nonce_range(
            &self, address: RpcAddress,
        ) -> JsonRpcResult<TxPoolPendingNonceRange>;
    
        #[rpc(name = "txpool_txWithPoolInfo")]
        fn txpool_tx_with_pool_info(
            &self, hash: H256,
        ) -> JsonRpcResult<TxWithPoolInfo>;
    
        /// Get transaction pending info by account address
        #[rpc(name = "txpool_accountPendingInfo")]
        fn account_pending_info(
            &self, address: RpcAddress,
        ) -> BoxFuture<Option<AccountPendingInfo>>;
    
        /// Get transaction pending info by account address
        #[rpc(name = "txpool_accountPendingTransactions")]
        fn account_pending_transactions(
            &self, address: RpcAddress, maybe_start_nonce: Option<U256>,
            maybe_limit: Option<U64>,
        ) -> BoxFuture<AccountPendingTransactions>;
    }
    `
    console.log(convertTrait2Funcs(rust, "Txpool"))

    const rustStruct = `pub struct AccountPendingInfo {
        pub local_nonce: U256,
        pub pending_count: U256,
        pub pending_nonce: U256,
        pub next_pending_tx: H256,
    }`
    console.log(convertStruct(rustStruct))
}

run()
*/