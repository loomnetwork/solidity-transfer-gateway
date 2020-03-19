pragma solidity ^0.4.23;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./ECVerify.sol";


contract ValidatorManagerContract is Ownable {
  using ECVerify for bytes32;

  mapping (address => bool) public allowedTokens;
  mapping (address => uint256) public nonces;
  mapping (address => bool) validators;

  uint8 threshold_num;
  uint8 threshold_denom;
  uint256 public numValidators;
  uint256 public nonce; // used for replay protection when adding/removing validators

  event AddedValidator(address validator);
  event RemovedValidator(address validator);

  modifier onlyValidator() { require(checkValidator(msg.sender)); _; }

  modifier onlyWithThreshold(bytes32 message, uint8[] _v, bytes32[] _r, bytes32[] _s) {
    // Check that we have enough signatures

    checkThreshold(message, _v, _r, _s);
    _;
    nonce++;
  }

  constructor (address _validators, uint8 _threshold_num, uint8 _threshold_denom) public {
    // uint256 length = _validators.length;
    // require(length > 0);

    threshold_num = _threshold_num;
    threshold_denom = _threshold_denom;
    // for (uint256 i = 0; i < length ; i++) {
      require(_validators != address(0), "Validator shouldnt be 0x0");
      validators[_validators] = true;
      emit AddedValidator(_validators);
    // }
    numValidators = 1;
  }

  modifier signedByValidator(bytes32 _message, bytes32 r, bytes32 s, uint8 v) {
    // prevent replay attacks by adding the nonce in the sig
    // if a validator signs an invalid nonce,
    // it won't pass the signature verification
    // since the nonce in the hash is stored in the contract
    address signer = ecrecover(_message, v, r, s);
    require(validators[signer], "Message not signed by a validator");
    _;
    nonces[msg.sender]++; // increment nonce after execution
  }

  function checkValidator(address _address) public view returns (bool) {
    // owner is a permanent validator
    if (_address == owner) {
      return true;
    }
    return validators[_address];
  }

  function addValidator(address _validator, uint8[] _v, bytes32[] _r, bytes32[] _s)
    external
    onlyWithThreshold(createMessage(keccak256(abi.encodePacked("add", _validator))), _v, _r, _s)
  {
    require(!validators[_validator], "Already a validator");

    // Add validator and increment nonce
    validators[_validator] = true;
    numValidators++;
    emit AddedValidator(_validator);
  }

  function removeValidator(address _validator, uint8[] _v, bytes32[] _r, bytes32[] _s)
    external
    onlyWithThreshold(createMessage(keccak256(abi.encodePacked("remove", _validator))), _v, _r, _s)
  {
    require(validators[_validator], "Not a validator");
    require(numValidators > 1);
    delete validators[_validator];
    numValidators--;

    emit RemovedValidator(_validator);
  }

  // Can't pass bytes[] to use the whole sig due to ABI enc
  //, so we need to send v,r,s params
  function checkThreshold(bytes32 _message, uint8[] _v, bytes32[] _r, bytes32[] _s) private view {
    require(_v.length > 0 && _v.length == _r.length && _r.length == _s.length,
      "Incorrect number of params");
    require(_v.length >= (threshold_num * numValidators / threshold_denom ),
      "Not enough votes");

    bytes32 hash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", _message));
    uint256 sig_length = _v.length;

    // Check that all addresess were from validators
    // Prevent duplicates by requiring that the sender sigs
    // get submitted in increasing order
    // influenced by
    // https://github.com/christianlundkvist/simple-multisig
    address lastAdd = address(0);
    for (uint256 i = 0; i < sig_length; i++) {
      address signer = ecrecover(hash, _v[i], _r[i], _s[i]);
      require(signer > lastAdd && validators[signer], "Not signed by enough validators");
      lastAdd = signer;
    }
  }

  /// @dev Toggles the token validate
  function toggleToken(address _token) public onlyValidator {
    allowedTokens[_token] = !allowedTokens[_token];
  }

  function createMessage(bytes32 hash)
      private
      view returns (bytes32)
  {
      return keccak256(
          abi.encodePacked(
              address(this),
              nonce,
              hash
          )
      );
  }
}
