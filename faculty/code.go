// Copyright 2021 The BBChain Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
// Code generated - DO NOT EDIT.
// This file was autogenerated with 'npm run abigen' from relab/bbchain-contracts repository and any manual changes will be lost.
package faculty
// The required libraries references are listed below.
var FacultyLinkReferences = map[string]string{"Notary":"1d2b6a72c2c8582bd594e19b02f34b62b6","CredentialSum":"9176519fc5fb44025266bd62b274a40525"}
// FacultyDeployedCode is the bytecode Faculty will have after deployment.
// It may contains unresolved link references for libraries.
const FacultyDeployedCode = "0x608060405234801561001057600080fd5b50600436106102535760003560e01c8063762b77bb11610146578063b9488546116100c3578063eedb973c11610087578063eedb973c14610646578063f537ff9d14610659578063f79fe5381461067c578063f936cc151461068f578063fc91a897146106a2578063fde08174146106c557610253565b8063b9488546146105b8578063c8bf608a146105c9578063dac46a62146105dc578063de07fc45146105fc578063e2d02cea1461063357610253565b8063a355f4a51161010a578063a355f4a514610523578063a6f9dae114610536578063abbff98414610549578063ae02c0cc1461059d578063affe39c1146105b057610253565b8063762b77bb146104c25780637f8d61e0146104e257806386509dd6146104f55780638ebd96231461050857806399d1dbc01461051b57610253565b80633610d3ca116101d45780634892e8e8116101985780634892e8e81461046e57806348aefc32146104765780635f889e17146104895780636f28eca61461049c5780637532ed26146104af57610253565b80633610d3ca1461038c5780633ac82e191461039f5780633eea79d1146103b25780634294857f146104385780634302802d1461044b57610253565b80631eee993a1161021b5780631eee993a146102ed5780632adfe8b2146103025780632cf0695a146103245780632f54bf6e14610337578063342577fa1461036357610253565b806301ffc9a714610258578063038ec1b814610280578063079cf76e146102a05780630ae07028146102c15780631703a018146102d4575b600080fd5b61026b610266366004613adf565b6106d8565b60405190151581526020015b60405180910390f35b61029361028e366004613a1a565b610711565b6040516102779190613d0e565b6102b36102ae3660046138bd565b61087a565b604051908152602001610277565b61026b6102cf36600461397e565b610898565b60005460ff165b60405160ff9091168152602001610277565b6103006102fb3660046138bd565b6108ad565b005b61026b610310366004613a1a565b6000908152600f6020526040902054151590565b61026b6103323660046138bd565b610a89565b61026b6103453660046138bd565b6001600160a01b031660009081526002602052604090205460ff1690565b6102b36103713660046138bd565b6001600160a01b031660009081526008602052604090205490565b61030061039a366004613abe565b610ad2565b6103006103ad3660046139a9565b610b0f565b6104166103c03660046138bd565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03949094168452600a8252928290208251938401835280548452600181015491840191909152600201549082015290565b6040805182518152602080840151908201529181015190820152606001610277565b61026b610446366004613a1a565b610f3d565b6102b3610459366004613a1a565b60009081526006602052604090206008015490565b610293610fcb565b61026b610484366004613a1a565b61102d565b61026b61049736600461397e565b61106f565b6102936104aa366004613a1a565b61107d565b61026b6104bd3660046138bd565b6110ec565b6104d56104d03660046138bd565b61113d565b6040516102779190613d34565b61026b6104f0366004613a4a565b6111a8565b61026b6105033660046138d9565b611245565b6104d56105163660046138bd565b611295565b61026b611448565b610293610531366004613a1a565b611476565b6103006105443660046138bd565b611535565b61055c610557366004613a1a565b611843565b604051610277919081516001600160a01b03908116825260208084015190911690820152604080830151908201526060918201519181019190915260800190565b61026b6105ab366004613a1a565b6118ed565b61029361193a565b6102db600054610100900460ff1690565b6103006105d7366004613a79565b61199a565b6105ef6105ea366004613a1a565b611ad3565b6040516102779190613e5e565b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152602001610277565b6102b36106413660046138d9565b611bad565b610300610654366004613a1a565b611be6565b6102b3610667366004613a1a565b60009081526006602052604090206007015490565b61026b61068a366004613a1a565b611bf2565b61026b61069d3660046138bd565b611c34565b61026b6106b03660046138bd565b600e6020526000908152604090205460ff1681565b600c5460ff166040516102779190613d91565b60006001600160e01b031982166301ffc9a760e01b148061070957506001600160e01b03198216635041329d60e01b145b90505b919050565b60015460609060009067ffffffffffffffff81111561074057634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610769578160200160208202803683370190505b5090506000805b60015481101561087157600085815260096020526040812060018054919291849081106107ad57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff161561085f57600181815481106107fb57634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b031683838151811061083957634e487b7160e01b600052603260045260246000fd5b6001600160a01b03909216602092830291909101909101528161085b81613fca565b9250505b8061086981613fca565b915050610770565b50909392505050565b6001600160a01b0381166000908152600a6020526040812054610709565b60006108a48383611e6b565b90505b92915050565b3360009081526002602052604090205460ff166108e55760405162461bcd60e51b81526004016108dc90613df0565b60405180910390fd5b6001600c5460ff16600181111561090c57634e487b7160e01b600052602160045260246000fd5b146109595760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f6e6f6465206d75737420626520496e6e657200000000000000000060448201526064016108dc565b60006109e230600b8054806020026020016040519081016040528092919081815260200182805480156109d857602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161099a5790505b5050505050611f78565b9050806109ff57634e487b7160e01b600052600160045260246000fd5b60008290506000816001600160a01b031663fde081746040518163ffffffff1660e01b815260040160206040518083038186803b158015610a3f57600080fd5b505afa158015610a53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a779190613b07565b9050610a838482611fec565b50505050565b6001600160a01b0381166000908152600560205260408120548290610ac05760405162461bcd60e51b81526004016108dc90613db9565b610acb6003846121e3565b9392505050565b3360009081526002602052604090205460ff16610b015760405162461bcd60e51b81526004016108dc90613df0565b610b0b8282612256565b5050565b3360009081526002602052604090205460ff16610b3e5760405162461bcd60e51b81526004016108dc90613df0565b6000600c5460ff166001811115610b6557634e487b7160e01b600052602160045260246000fd5b1415610bdd57805115610bba5760405162461bcd60e51b815260206004820152601f60248201527f4e6f64652f4c6561662063616e6e6f742068617665207769746e65737365730060448201526064016108dc565b60408051600080825260208201909252610bd8918591859190612326565b610f38565b6001600c5460ff166001811115610c0457634e487b7160e01b600052602160045260246000fd5b14610c1f57634e487b7160e01b600052600160045260246000fd5b6000815111610c695760405162461bcd60e51b8152602060048201526016602482015275139bd9194bddda5d1b995cdcc81b9bdd08199bdd5b9960521b60448201526064016108dc565b6000815167ffffffffffffffff811115610c9357634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610cbc578160200160208202803683370190505b50905060005b8251811015610e9c576000838281518110610ced57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600e90925260409091205490915060ff16610d655760405162461bcd60e51b815260206004820152601b60248201527f4e6f64652f61646472657373206e6f7420617574686f72697a6564000000000060448201526064016108dc565b6000610d7882635041329d60e01b6123f3565b905080610d9557634e487b7160e01b600052600160045260246000fd5b6040516303ce7bb760e11b81526001600160a01b038881166004830152839160009183169063079cf76e9060240160206040518083038186803b158015610ddb57600080fd5b505afa158015610def573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e139190613a32565b905080610e585760405162461bcd60e51b8152602060048201526013602482015272139bd9194bdc9bdbdd081b9bdd08199bdd5b99606a1b60448201526064016108dc565b80868681518110610e7957634e487b7160e01b600052603260045260246000fd5b602002602001018181525050505050508080610e9490613fca565b915050610cc2565b50604051637d71689560e11b815260009073__$9176519fc5fb44025266bd62b274a40525$__9063fae2d12a90610ed7908590600401613d21565b60206040518083038186803b158015610eef57600080fd5b505af4158015610f03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f279190613a32565b9050610f3585858386612326565b50505b505050565b604051631999f00560e21b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__90636667c014906044015b60206040518083038186803b158015610f9357600080fd5b505af4158015610fa7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107099190613a00565b6060600d80548060200260200160405190810160405280929190818152602001828054801561102357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611005575b5050505050905090565b6040516312f4761160e31b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__906397a3b08890604401610f7b565b60006108a46003848461240f565b6000818152600660209081526040918290206007018054835181840281018401909452808452606093928301828280156110e057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116110c2575b50505050509050919050565b6001600160a01b0381166000908152600a60205260408082209051634650570f60e01b8152600481019190915273__$9176519fc5fb44025266bd62b274a40525$__90634650570f90602401610f7b565b6001600160a01b0381166000908152600560209081526040918290208054835181840281018401909452808452606093928301828280156110e057602002820191906000526020600020905b8154815260200190600101908083116111895750505050509050919050565b60405163a8b0bcfd60e01b815260036004820152602481018390526001600160a01b038216604482015260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063a8b0bcfd906064015b60206040518083038186803b15801561120d57600080fd5b505af4158015611221573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a49190613a00565b6001600160a01b0382166000908152600a6020526040808220905163070f6e8d60e01b815273__$9176519fc5fb44025266bd62b274a40525$__9163070f6e8d916111f591908690600401613d78565b6001600160a01b03811660009081526005602052604090205460609082906112cf5760405162461bcd60e51b81526004016108dc90613db9565b6001600160a01b03831660009081526008602052604081205467ffffffffffffffff81111561130e57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611337578160200160208202803683370190505b5090506000805b6001600160a01b03861660009081526005602052604090205481101561143e576001600160a01b038616600090815260056020526040902080546113ab91908390811061139b57634e487b7160e01b600052603260045260246000fd5b9060005260206000200154610f3d565b1561142c576001600160a01b03861660009081526005602052604090208054829081106113e857634e487b7160e01b600052603260045260246000fd5b906000526020600020015483838151811061141357634e487b7160e01b600052603260045260246000fd5b60209081029190910101528161142881613fca565b9250505b8061143681613fca565b91505061133e565b5090949350505050565b600080600c5460ff16600181111561147057634e487b7160e01b600052602160045260246000fd5b14905090565b6000818152600f60205260409020546060906114cd5760405162461bcd60e51b8152602060048201526016602482015275119858dd5b1d1e4bdb9bdd081c9959da5cdd195c995960521b60448201526064016108dc565b6000828152600f6020908152604091829020805483518184028101840190945280845290918301828280156110e0576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116110c25750505050509050919050565b3360009081526002602052604090205460ff166115645760405162461bcd60e51b81526004016108dc90613df0565b6001600160a01b03811660009081526002602052604090205460ff1615801561159557506001600160a01b03811615155b6115e15760405162461bcd60e51b815260206004820152601c60248201527f4f776e6572732f696e76616c6964206164647265737320676976656e0000000060448201526064016108dc565b600154158015906115f5575060015460ff10155b61160f57634e487b7160e01b600052600160045260246000fd5b60015460009067ffffffffffffffff81111561163b57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611664578160200160208202803683370190505b50905060005b60015460ff8216101561179e57336001600160a01b031660018260ff16815481106116a557634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b0316146117495760018160ff16815481106116e357634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828260ff168151811061172457634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b03168152505061178e565b82828260ff168151811061176d57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b61179781613fe5565b905061166a565b50600054815160ff909116146117c457634e487b7160e01b600052600160045260246000fd5b6040516001600160a01b0383169033907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c90600090a3805161180d906001906020840190613750565b50506001600160a01b0316600090815260026020526040808220805460ff19908116600117909155338352912080549091169055565b6040805160808101825260008082526020820181905281830181905260608201529051636dbae62960e11b8152600360048201526024810183905273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063db75cc529060440160806040518083038186803b1580156118b557600080fd5b505af41580156118c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107099190613c0b565b60008054604051633e01fd5d60e21b8152600360048201526024810184905260ff909116604482015273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063f807f57490606401610f7b565b60606001805480602002602001604051908101604052809291908181526020018280548015611023576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611005575050505050905090565b3360009081526002602052604090205460ff166119c95760405162461bcd60e51b81526004016108dc90613df0565b6000828152600f602052604090205415611a255760405162461bcd60e51b815260206004820152601a60248201527f466163756c74792f616c7265616479207265676973746572656400000000000060448201526064016108dc565b6000815111611a765760405162461bcd60e51b815260206004820152601a60248201527f466163756c74792f656d70747920636f7572736573206c69737400000000000060448201526064016108dc565b6000828152600f602090815260409091208251611a9592840190613750565b5060408051338152602081018490527fcb185a4a072cf4e8ef43d02535b01a5ca5a4eeb8dcbc242651888e159e6c8bf0910160405180910390a15050565b604080516101408101825260008082526020820181905281830181905260608083018290526080830182905260a0830182905260c0830182905260e0830182905261010083015261012082015290516301c30fd760e71b8152600360048201526024810183905273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063e187eb809060440160006040518083038186803b158015611b7157600080fd5b505af4158015611b85573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107099190810190613b26565b3360009081526002602052604081205460ff16611bdc5760405162461bcd60e51b81526004016108dc90613df0565b6108a48383612517565b611bef81612678565b50565b60405163143ceab160e21b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__906350f3aac490604401610f7b565b600080611c408361113d565b90506000815111611c935760405162461bcd60e51b815260206004820152601960248201527f4e6f64652f63726564656e7469616c206e6f7420666f756e640000000000000060448201526064016108dc565b611c9c836110ec565b15611cb957611cab8382611245565b611cb957600091505061070c565b60005b8151811015611e6157611cf5828281518110611ce857634e487b7160e01b600052603260045260246000fd5b6020026020010151611bf2565b611d0f57634e487b7160e01b600052600160045260246000fd5b611d4084838381518110611d3357634e487b7160e01b600052603260045260246000fd5b602002602001015161106f565b611d4f5760009250505061070c565b6001600c5460ff166001811115611d7657634e487b7160e01b600052602160045260246000fd5b148015611dc257506000611dc0838381518110611da357634e487b7160e01b600052603260045260246000fd5b602002602001015160009081526006602052604090206007015490565b115b15611e4f57611e4084611e0b848481518110611dee57634e487b7160e01b600052603260045260246000fd5b602002602001015160009081526006602052604090206008015490565b611e3b858581518110611e2e57634e487b7160e01b600052603260045260246000fd5b602002602001015161107d565b61274b565b611e4f5760009250505061070c565b80611e5981613fca565b915050611cbc565b5060019392505050565b6001600160a01b0382166000908152600560205260408120548390611ea25760405162461bcd60e51b81526004016108dc90613db9565b6001600160a01b0384166000908152600a60209081526040808320600590925291829020915163070f6e8d60e01b815273__$9176519fc5fb44025266bd62b274a40525$__9263070f6e8d92611efd92909190600401613f32565b60206040518083038186803b158015611f1557600080fd5b505af4158015611f29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f4d9190613a00565b8015611f7057506001600160a01b0384166000908152600a602052604090205483145b949350505050565b6000611f8383612bea565b611f8f575060006108a7565b60005b8251811015611e6157611fcc84848381518110611fbf57634e487b7160e01b600052603260045260246000fd5b6020026020010151612c1d565b611fda5760009150506108a7565b80611fe481613fca565b915050611f92565b306001600160a01b038316141561203e5760405162461bcd60e51b81526020600482015260166024820152752737b23297b1b0b73737ba1030b2321034ba39b2b63360511b60448201526064016108dc565b6001600160a01b0382166000908152600e602052604090205460ff16156120a75760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f6e6f646520616c726561647920616464656400000000000000000060448201526064016108dc565b60008160018111156120c957634e487b7160e01b600052602160045260246000fd5b14806120f4575060018160018111156120f257634e487b7160e01b600052602160045260246000fd5b145b6121405760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f696e76616c6964206368696c6420726f6c6500000000000000000060448201526064016108dc565b6001600160a01b0382166000818152600e6020526040808220805460ff19166001908117909155600d805491820181559092527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb590910180546001600160a01b031916831790555133907f202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f906121d7908590613d91565b60405180910390a35050565b6001600160a01b0381166000908152600283016020908152604080832080548251818502810185019093528083526108a493879387939092909183018282801561224c57602002820191906000526020600020905b815481526020019060010190808311612238575b5050505050612c40565b8161226081610f3d565b1561227d5760405162461bcd60e51b81526004016108dc90613e27565b600083815260066020819052604090912001546001600160a01b03166122bb336001600160a01b031660009081526002602052604090205460ff1690565b806122ce57506001600160a01b03811633145b61231a5760405162461bcd60e51b815260206004820152601c60248201527f4973737565722f73656e646572206e6f7420617574686f72697a65640000000060448201526064016108dc565b610a8360038585612caa565b3360009081526002602052604090205460ff166123555760405162461bcd60e51b81526004016108dc90613df0565b8261235f81610f3d565b1561237c5760405162461bcd60e51b81526004016108dc90613e27565b6001600160a01b03851660009081526002602052604090205460ff16156123e55760405162461bcd60e51b815260206004820152601a60248201527f4973737565722f666f7262696464656e2072656769737472617200000000000060448201526064016108dc565b610f35600386868686612e59565b60006123fe83612bea565b80156108a457506108a48383612c1d565b600081815260038401602052604081206001015461246f5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f63726564656e7469616c206e6f7420666f756e64000000000060448201526064016108dc565b60008281526003850160205260409020600601546001600160a01b038481169116146124dd5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f6e6f74206f776e6564206279207375626a656374000000000060448201526064016108dc565b600082815260038501602052604090206005015460ff168015611f7057505060009081526004929092016020525060409020600201541590565b3360009081526002602052604081205460ff166125465760405162461bcd60e51b81526004016108dc90613df0565b6001600160a01b038316600090815260056020526040902054839061257d5760405162461bcd60e51b81526004016108dc90613db9565b61258960038585612c40565b6125d55760405162461bcd60e51b815260206004820152601e60248201527f4973737565722f68617320696e76616c69642063726564656e7469616c73000060448201526064016108dc565b6001600160a01b0384166000908152600a6020526040908190209051631b33c8e760e31b815273__$9176519fc5fb44025266bd62b274a40525$__9163d99e473891612628919088908890600401613f08565b60206040518083038186803b15801561264057600080fd5b505af4158015612654573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f709190613a32565b8061268281610f3d565b1561269f5760405162461bcd60e51b81526004016108dc90613e27565b60005460ff166126ea5760405162461bcd60e51b8152602060048201526016602482015275125cdcdd595c8bdb9bc81c5d5bdc9d5b48199bdd5b9960521b60448201526064016108dc565b612706826126fa60005460ff1690565b6003919060ff166134ad565b610b0b5760405162461bcd60e51b8152602060048201526016602482015275125cdcdd595c8bd85c1c1c9bdd985b0819985a5b195960521b60448201526064016108dc565b60008261279a5760405162461bcd60e51b815260206004820152601860248201527f4e6f64652f726f6f742063616e6e6f74206265206e756c6c000000000000000060448201526064016108dc565b6000825167ffffffffffffffff8111156127c457634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156127ed578160200160208202803683370190505b50905060005b8351811015612b5757600084828151811061281e57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600e90925260409091205490915060ff166128965760405162461bcd60e51b815260206004820152601b60248201527f4e6f64652f61646472657373206e6f7420617574686f72697a6564000000000060448201526064016108dc565b60006128a982635041329d60e01b6123f3565b9050806128c657634e487b7160e01b600052600160045260246000fd5b6000829050806001600160a01b03166399d1dbc06040518163ffffffff1660e01b815260040160206040518083038186803b15801561290457600080fd5b505afa158015612918573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061293c9190613a00565b15612ab6576040516303ce7bb760e11b81526001600160a01b038a8116600483015282169063079cf76e9060240160206040518083038186803b15801561298257600080fd5b505afa158015612996573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129ba9190613a32565b8585815181106129da57634e487b7160e01b600052603260045260246000fd5b602002602001018181525050806001600160a01b0316630ae070288a878781518110612a1657634e487b7160e01b600052603260045260246000fd5b60200260200101516040518363ffffffff1660e01b8152600401612a4f9291906001600160a01b03929092168252602082015260400190565b60206040518083038186803b158015612a6757600080fd5b505afa158015612a7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a9f9190613a00565b612ab157600095505050505050610acb565b612b41565b60405163f936cc1560e01b81526001600160a01b038a8116600483015282169063f936cc159060240160206040518083038186803b158015612af757600080fd5b505afa158015612b0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b2f9190613a00565b612b4157600095505050505050610acb565b5050508080612b4f90613fca565b9150506127f3565b50604051630b6d99fb60e11b815273__$9176519fc5fb44025266bd62b274a40525$__906316db33f690612b919087908590600401613d78565b60206040518083038186803b158015612ba957600080fd5b505af4158015612bbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612be19190613a00565b95945050505050565b6000612bfd826301ffc9a760e01b612c1d565b80156107095750612c16826001600160e01b0319612c1d565b1592915050565b6000806000612c2c8585613665565b91509150818015612be15750949350505050565b6000805b8251811015612c9f57612c7f8585858481518110612c7257634e487b7160e01b600052603260045260246000fd5b602002602001015161240f565b612c8d576000915050610acb565b80612c9781613fca565b915050612c44565b506001949350505050565b6000828152600384016020526040902060010154612d0a5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f63726564656e7469616c206e6f7420666f756e64000000000060448201526064016108dc565b60008281526003840160205260409020600601546001600160a01b031680612d745760405162461bcd60e51b815260206004820152601d60248201527f4e6f746172792f7375626a6563742063616e6e6f74206265207a65726f00000060448201526064016108dc565b6001600160a01b038116600090815260058501602052604081208054909190612d9c90613fca565b9091555060408051608081018252338082526001600160a01b03848116602080850182815243868801818152606088018b815260008d815260048f0186528a9020985189549088166001600160a01b0319918216178a55935160018a0180549190981694169390931790955593516002870155516003909501949094558451918252928101869052909286917fb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f910160405180910390a450505050565b6000838152600686016020908152604080832033845290915290205460ff1615612ec55760405162461bcd60e51b815260206004820152601c60248201527f4e6f746172792f73656e64657220616c7265616479207369676e65640000000060448201526064016108dc565b6000838152600386016020526040902060010154613235576001600160a01b038416600090815260018601602052604090205415613090576001600160a01b0384166000908152600180870160209081526040808420548452600389019091529091200154612f4457634e487b7160e01b600052600160045260246000fd5b6001600160a01b0380851660009081526001808801602090815260408084205484526003808b01835281852082516101408101845281548152948101548585015260028101548584015290810154606085015260048101546080850152600581015460ff8116151560a086015261010090819004871660c0860152600682015490961660e085015260078101805483518186028101860190945280845295969495919486019390919083018282801561302657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613008575b5050505050815260200160088201548152505090504381602001511061308e5760405162461bcd60e51b815260206004820152601d60248201527f4e6f746172792f626c6f636b206e756d6265722076696f6c6174696f6e00000060448201526064016108dc565b505b60405180610140016040528060018152602001438152602001428152602001866000016000876001600160a01b03166001600160a01b03168152602001908152602001600020600081546130e390613fca565b91829055508152602080820186905260006040808401829052336060808601919091526001600160a01b038a811660808088019190915260a08088018a905260c09788018b90528b865260038e81018852958590208951815589880151600182015594890151600286015592880151948401949094559286015160048301558501516005820180549587015160ff1990961691151591909117610100600160a81b031916610100958416860217905560e08501516006820180546001600160a01b03191691909316179091559183015180516131c59260078501920190613750565b5061012091909101516008909101556001600160a01b0384166000818152600187016020526040908190208590555133919085907fc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be906132289043815260200190565b60405180910390a461344b565b60008381526003808701602090815260408084208151610140810183528154815260018201548185015260028201548184015293810154606085015260048101546080850152600581015460ff8116151560a08601526001600160a01b0361010091829004811660c087015260068301541660e086015260078201805484518187028101870190955280855292949186019392909183018282801561330357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116132e5575b505050505081526020016008820154815250509050846001600160a01b03168160e001516001600160a01b03161461337d5760405162461bcd60e51b815260206004820181905260248201527f4e6f746172792f64696765737420616c7265616479207265676973746572656460448201526064016108dc565b82816101200151146133d15760405162461bcd60e51b815260206004820152601f60248201527f4e6f746172792f6d69736d6174636865642065766964656e636520726f6f740060448201526064016108dc565b815181610100015151146134275760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f6d69736d617463686564207769746e6573736573000000000060448201526064016108dc565b60008481526003870160205260408120805490919061344590613fca565b90915550505b600083815260068601602090815260408083203380855290835292819020805460ff19166001179055514381528592917f3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12910160405180910390a35050505050565b60008281526003840160205260408120600601546001600160a01b03163381146135105760405162461bcd60e51b8152602060048201526014602482015273139bdd185c9e4bdddc9bdb99c81cdd589a9958dd60621b60448201526064016108dc565b600084815260038601602052604090206005015460ff16156135745760405162461bcd60e51b815260206004820181905260248201527f4e6f746172792f63726564656e7469616c20616c7265616479207369676e656460448201526064016108dc565b60008481526003860160205260409020548311156135d45760405162461bcd60e51b815260206004820152601e60248201527f4e6f746172792f6e6f2071756f72756d206f66207369676e617475726573000060448201526064016108dc565b60008481526003860160209081526040808320600501805460ff191660019081179091556001600160a01b0385168452600289018352818420805491820181558452919092200185905551849033907f3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12906136529043815260200190565b60405180910390a3506001949350505050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b17905290516000918291829081906001600160a01b03881690617530906136cd908690613cd5565b6000604051808303818686fa925050503d8060008114613709576040519150601f19603f3d011682016040523d82523d6000602084013e61370e565b606091505b509150915060208151101561372c5760008094509450505050613749565b81818060200190518101906137419190613a00565b945094505050505b9250929050565b8280548282559060005260206000209081019282156137a5579160200282015b828111156137a557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613770565b506137b19291506137b5565b5090565b5b808211156137b157600081556001016137b6565b805161070c81614031565b600082601f8301126137e5578081fd5b813560206137fa6137f583613fa6565b613f75565b8281528181019085830183850287018401881015613816578586fd5b855b8581101561383d57813561382b81614031565b84529284019290840190600101613818565b5090979650505050505050565b600082601f83011261385a578081fd5b8151602061386a6137f583613fa6565b8281528181019085830183850287018401881015613886578586fd5b855b8581101561383d57815161389b81614031565b84529284019290840190600101613888565b8051801515811461070c57600080fd5b6000602082840312156138ce578081fd5b8135610acb81614031565b600080604083850312156138eb578081fd5b82356138f681614031565b915060208381013567ffffffffffffffff811115613912578283fd5b8401601f81018613613922578283fd5b80356139306137f582613fa6565b81815283810190838501858402850186018a101561394c578687fd5b8694505b8385101561396e578035835260019490940193918501918501613950565b5080955050505050509250929050565b60008060408385031215613990578182fd5b823561399b81614031565b946020939093013593505050565b6000806000606084860312156139bd578081fd5b83356139c881614031565b925060208401359150604084013567ffffffffffffffff8111156139ea578182fd5b6139f6868287016137d5565b9150509250925092565b600060208284031215613a11578081fd5b6108a4826138ad565b600060208284031215613a2b578081fd5b5035919050565b600060208284031215613a43578081fd5b5051919050565b60008060408385031215613a5c578182fd5b823591506020830135613a6e81614031565b809150509250929050565b60008060408385031215613a8b578182fd5b82359150602083013567ffffffffffffffff811115613aa8578182fd5b613ab4858286016137d5565b9150509250929050565b60008060408385031215613ad0578182fd5b50508035926020909101359150565b600060208284031215613af0578081fd5b81356001600160e01b031981168114610acb578182fd5b600060208284031215613b18578081fd5b815160028110610acb578182fd5b600060208284031215613b37578081fd5b815167ffffffffffffffff80821115613b4e578283fd5b8184019150610140808387031215613b64578384fd5b613b6d81613f75565b90508251815260208301516020820152604083015160408201526060830151606082015260808301516080820152613ba760a084016138ad565b60a0820152613bb860c084016137ca565b60c0820152613bc960e084016137ca565b60e08201526101008084015183811115613be1578586fd5b613bed8882870161384a565b91830191909152506101209283015192810192909252509392505050565b600060808284031215613c1c578081fd5b613c266080613f75565b8251613c3181614031565b81526020830151613c4181614031565b6020820152604083810151908201526060928301519281019290925250919050565b6000815180845260208085019450808401835b83811015613c9b5781516001600160a01b031687529582019590820190600101613c76565b509495945050505050565b6000815180845260208085019450808401835b83811015613c9b57815187529582019590820190600101613cb9565b60008251815b81811015613cf55760208186018101518583015201613cdb565b81811115613d035782828501525b509190910192915050565b6000602082526108a46020830184613c63565b6000602082526108a46020830184613ca6565b6020808252825182820181905260009190848201906040850190845b81811015613d6c57835183529284019291840191600101613d50565b50909695505050505050565b600083825260406020830152611f706040830184613ca6565b6020810160028310613db357634e487b7160e01b600052602160045260246000fd5b91905290565b6020808252601f908201527f4973737565722f746865726520617265206e6f2063726564656e7469616c7300604082015260600190565b6020808252601d908201527f4f776e6572732f73656e646572206973206e6f7420616e206f776e6572000000604082015260600190565b60208082526019908201527f4973737565722f63726564656e7469616c207265766f6b656400000000000000604082015260600190565b60006020825282516020830152602083015160408301526040830151606083015260608301516080830152608083015160a083015260a0830151613ea660c084018215159052565b5060c08301516001600160a01b03811660e08401525060e0830151610100613ed8818501836001600160a01b03169052565b808501519150506101406101208181860152613ef8610160860184613c63565b9501519301929092525090919050565b8381526001600160a01b0383166020820152606060408201819052600090612be190830184613ca6565b60006040820184835260206040818501528185548084526060860191508685528285209350845b8181101561383d57845483526001948501949284019201613f59565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f9e57613f9e61401b565b604052919050565b600067ffffffffffffffff821115613fc057613fc061401b565b5060209081020190565b6000600019821415613fde57613fde614005565b5060010190565b600060ff821660ff811415613ffc57613ffc614005565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611bef57600080fdfea2646970667358221220480cd05b9369ed671e507eaff98bbed09902abc0e87600691b0570bb046e558864736f6c63430008020033"
