// TODO: create a web page tool

const goTypeMap = {
    "u64": "hexutil.Uint64",
    "U64": "hexutil.Uint64",
    "U256": "*hexutil.Big",
    "H256": "types.Hash",
    "RpcAddress": "types.Address",
    "RpcTransaction": "types.Transaction",
}

function convertStruct(struct) {
    // console.log("struct", struct)
    const any = "[\\s\\S]"
    const regex = new RegExp(`(pub struct .*)\\{(${any}*)\\}`, "i")

    if (!regex.test(struct)) return ""

    let [, head, body] = struct.match(regex)

    const goHead = head.replace(/pub struct(.*)/, "type $1 struct")
    const goBody = body.replaceAll(new RegExp(`${any}*?pub (.*)?: (.*)?\\,`, "img"),
        (match, name, type) => `${toAllCamel(name)} ${convertType(type)} \`json:"${toCamel(name)}"\`\n`)
    return `${goHead}{\n${goBody}}`
}

function convertStructs(structs) {
    return structs
        .split("pub struct")
        .map(line => convertStruct("pub struct" + line))
        .join("\n")
}

function convertFunction(nameSpace, func) {
    const any = "[\\s\\S]"
    const regex = new RegExp(`(${any}*)?#\\[rpc.*?"(.*)?"\\)\\]${any}*?fn(${any}*)?\\(\\s*\\&self,*(${any}*)?\\)\\s*->\\s*.*?<(.*)?>`, "img")

    let goFunc = func.replaceAll(regex, (m, comment, rpcMethod, funName, args, returnType) => {
        // args = convertArgs(args)
        return `${comment || ''}func(c *Rpc${nameSpace}Client) ${toAllCamel(funName).replace(nameSpace, "")}(${convertArgs(args)})(val ${convertType(returnType)}, err error) {
        err = c.core.CallRPC(&val, "${rpcMethod}", ${convertArgs(args,false)})
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
    if (/Option<(.*)>/.test(returnType)) {
        return returnType
            .replace(/Option<(.*)>/, (m, c) => `*${convertType(c)}`)
            .replace("**", "*")
    }

    if (/Vec<(.*)>/.test(returnType)) {
        return returnType
            .replace(/Vec<(.*)>/, (m, c) => `[]${convertType(c)}`)
    }

    return goTypeMap[returnType] || returnType
}

function convertArgs(args, isNeedType = true) {
    if (!args) return ""

    const argReplaceFunc = (m, name, type) => `${toCamel(name)} ${isNeedType ? convertType(type) : ""}`
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
    convertStructs,
    convertFunction,
    convertTrait2Funcs
}

/** Demo
function run() {
    const rust = `    // return account ready + deferred transactions
    #[rpc(name = "txpool_accountTransactions")]
    fn txpool_get_account_transactions(
        &self, address: RpcAddress,
    ) -> JsonRpcResult<Vec<RpcTransaction>>;
    `
    console.log(convertTrait2Funcs(rust, "Debug"))

    const rustStructs = `#[derive(Debug, Serialize, Clone, Deserialize)]
    #[serde(rename_all = "camelCase")]
    pub struct PoSEconomics {
        // This is the total number of CFX used for pos staking.
        pub total_pos_staking_tokens: U256,
        // This is the total distributable interest.
        pub distributable_pos_interest: U256,
        // This is the block number of last .
        pub last_distribute_block: u64,
    }`
    console.log(convertStructs(rustStructs))
}

run()
*/