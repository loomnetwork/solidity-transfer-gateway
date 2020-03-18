var SampleERC721Token = artifacts.require("SampleERC721Token");

module.exports = function(deployer) {
  deployer.deploy(SampleERC721Token, "0xb9fA0896573A89cF4065c43563C069b3B3C15c37");
};