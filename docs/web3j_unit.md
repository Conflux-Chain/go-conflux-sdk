Web3j Unit
===========

Web3j-unit is an extension to [JUnit 5](https://junit.org/junit5/docs/current/user-guide/) which enables you to test solidity contracts like any other Java code.
It allows you to test with both an embedded and dockerized Ethereum nodes, with out-of-the box support for Geth, Besu, and OpenEthereum nodes. A docker-compose network can also be configured easily for more complex setups.

## Usage

First, let's add the gradle dependency.

```groovy
repositories {
  mavenCentral()
  jcenter()
}
implementation "org.web3j:core:4.5.11"
testCompile "org.web3j:web3j-unit:4.5.11"
```


### Using EVMTest annotation

Create a new test with the `@EVMTest` annotation. An embedded EVM is used by default. To use Geth or Besu pass the node type into the annotation: `@EVMTest(NodeType.GETH)` or `@EVMTest(NodeType.BESU)`
```java
@EVMTest
public class GreeterTest {

}
```

Inject instance of `Web3j`, `TransactionManager` and `ContractGasProvider` in your test method.
```java
@EVMTest
public class GreeterTest {

    @Test
    public void greeterDeploys(Web3j web3j, TransactionManager transactionManager, ContractGasProvider gasProvider ) {
    }

}
```

Deploy your contract in the test.
```java
@EVMTest
public class GreeterTest {

    @Test
    public void greeterDeploys(Web3j web3j, TransactionManager transactionManager, ContractGasProvider gasProvider) {
        Greeter greeter = Greeter.deploy(web3j, transactionManager, gasProvider, "Hello EVM").send();
        String greeting = greeter.greet().send();
        assertEquals("Hello EVM", greeting);
    }

}
```

Run the test!

### Using EVMComposeTest annotation

Create a new test with the `@EVMComposeTest` annotation. 

By default, it uses `test.yml` file in the project home, and runs web3j on service named `node1` exposing the port `8545`. 

Can be customised to use specific docker-compose file, service name and port by `@EVMComposeTest("src/test/resources/geth.yml", "ethnode1", 8080)` Here, we connect to the service named `ethnode1` in the `src/test/resources/geth.yml` docker-compose file which exposes the port `8080` for web3j to connect to.

```java
@EVMComposeTest("src/test/resources/geth.yml", "ethnode1", 8080)
public class GreeterTest {

}
```

Inject instance of `Web3j`, `TransactionManager` and `ContractGasProvider` in your test method.

```java
@EVMComposeTest("src/test/resources/geth.yml", "ethnode1", 8080)
public class GreeterTest {

    @Test
    public void greeterDeploys(Web3j web3j, TransactionManager transactionManager, ContractGasProvider gasProvider) {
    }

}
```

Deploy your contract in the test.

```java
@EVMComposeTest("src/test/resources/geth.yml", "ethnode1", 8080)
public class GreeterTest {

    @Test
    public void greeterDeploys(Web3j web3j, TransactionManager transactionManager, ContractGasProvider gasProvider) {
        Greeter greeter = Greeter.deploy(web3j, transactionManager, gasProvider, "Hello EVM").send();
        String greeting = greeter.greet().send();
        assertEquals("Hello EVM", greeting);
    }

}
```

Run the test!

## Sample projects

1. Sample project using `@EVMTest` can be found [here](https://github.com/web3j/web3j-unitexample).

2. [This](https://github.com/web3j/web3j-unit-docker-compose-example) uses `@EVMComposeTest` to test the Greeter contract using VMWare Concord nodes using a docker-compose file.

