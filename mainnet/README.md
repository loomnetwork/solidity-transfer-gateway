# Transfer Gateway Contracts

The core components of the repository are:
- `Gateway.sol`
- `ValidatorManagerContract.sol`
`GameToken.sol` implements an ERC20 token contract which additionally implements functions to make callbacks to the receiver during a transfer, if it's a contract. This allows for 1-step transfer-and-call patterns, which greatly improve UX. `Cards.sol` implements an ERC721 Non-Fungible-Token contract. Both contracts are using the OpenZeppelin library and are used for demonstration purposes. `ECVerify.sol` verifies the sender of a signature. The expected `signature` argument must be in the form of [ MODE | V | R | S ], where MODE = [ 0, 1, 2]. 

## Validator Manager Contract

The `ValidatorManagerContract.sol` contract facilitates a general contract that can be inherited for implementing permission control on functions, based on a predefined set of validators. These validators are set during deployment via the constructor, along with a threshold in the form of numerator and denominator (due to solidity not supporting float numbers). 

In order to add or remove a validator, a number of signed messages from validators that exceed the threshold must be submitted. 

A `isVerifiedByValidator` modifier is also implemented which requires that the caller of a function must also submit a message signed by the validator for that function call. 

A global nonce is used for the add/remove validator process, and a per user nonce is used for the function call process for replay protection.

## Transfer Gateway

The `Gateway.sol` contract implements a contract which locks any ETH or ERC20/ERC721 tokens deposited to it, and emits a `Received` event depending on the type of coin that was deposited. A DAppChain is listening for this kind of events and corresponds accordingly. Balances are kept for each user to ensure that no coins are possible to be stolen. In order for a user to withdraw their coins they also need to submit a message signed by the validator, permitting them to withdraw that specific coin, by using the `isVerifiedByValidator` mentioned earlier.

## Gas Costs

All numbers are expected to have some small variance due to difference in gas costs because of input size in a transaction. Running `npm run test:gas` can also be used for rough estimates.

### Gateway:
- Ether Deposit: 42824 
- Ether Withdrawal: 51795
- ER20 Deposit: 61531
- ERC20 Withdrawal: 61355
- ERC721 Deposit: 140191
- ERC721 Withdrawal: 111292

### VMC: 
Refer to `ValidatorCosts.md`
