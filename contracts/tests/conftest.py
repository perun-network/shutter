from typing import Any
from typing import List
from typing import Sequence

import pytest
from brownie.network.account import Account
from brownie.network.contract import ContractContainer
from brownie.network.web3 import Web3
from eth_utils import decode_hex

ERC_1820_DEPLOYMENT_TX = decode_hex(
    "0xf90a388085174876e800830c35008080b909e5608060405234801561001057600080fd5b506109c5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c010000000000000000000000000000000000000000000000000000000090048063a41e7d5111610078578063a41e7d51146101d4578063aabbb8ca1461020a578063b705676514610236578063f712f3e814610280576100a5565b806329965a1d146100aa5780633d584063146100e25780635df8122f1461012457806365ba36c114610152575b600080fd5b6100e0600480360360608110156100c057600080fd5b50600160a060020a038135811691602081013591604090910135166102b6565b005b610108600480360360208110156100f857600080fd5b5035600160a060020a0316610570565b60408051600160a060020a039092168252519081900360200190f35b6100e06004803603604081101561013a57600080fd5b50600160a060020a03813581169160200135166105bc565b6101c26004803603602081101561016857600080fd5b81019060208101813564010000000081111561018357600080fd5b82018360208201111561019557600080fd5b803590602001918460018302840111640100000000831117156101b757600080fd5b5090925090506106b3565b60408051918252519081900360200190f35b6100e0600480360360408110156101ea57600080fd5b508035600160a060020a03169060200135600160e060020a0319166106ee565b6101086004803603604081101561022057600080fd5b50600160a060020a038135169060200135610778565b61026c6004803603604081101561024c57600080fd5b508035600160a060020a03169060200135600160e060020a0319166107ef565b604080519115158252519081900360200190f35b61026c6004803603604081101561029657600080fd5b508035600160a060020a03169060200135600160e060020a0319166108aa565b6000600160a060020a038416156102cd57836102cf565b335b9050336102db82610570565b600160a060020a031614610339576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b6103428361092a565b15610397576040805160e560020a62461bcd02815260206004820152601a60248201527f4d757374206e6f7420626520616e204552433136352068617368000000000000604482015290519081900360640190fd5b600160a060020a038216158015906103b85750600160a060020a0382163314155b156104ff5760405160200180807f455243313832305f4143434550545f4d4147494300000000000000000000000081525060140190506040516020818303038152906040528051906020012082600160a060020a031663249cb3fa85846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182600160a060020a0316600160a060020a031681526020019250505060206040518083038186803b15801561047e57600080fd5b505afa158015610492573d6000803e3d6000fd5b505050506040513d60208110156104a857600080fd5b5051146104ff576040805160e560020a62461bcd02815260206004820181905260248201527f446f6573206e6f7420696d706c656d656e742074686520696e74657266616365604482015290519081900360640190fd5b600160a060020a03818116600081815260208181526040808320888452909152808220805473ffffffffffffffffffffffffffffffffffffffff19169487169485179055518692917f93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db15391a450505050565b600160a060020a03818116600090815260016020526040812054909116151561059a5750806105b7565b50600160a060020a03808216600090815260016020526040902054165b919050565b336105c683610570565b600160a060020a031614610624576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b81600160a060020a031681600160a060020a0316146106435780610646565b60005b600160a060020a03838116600081815260016020526040808220805473ffffffffffffffffffffffffffffffffffffffff19169585169590951790945592519184169290917f605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a43509190a35050565b600082826040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012090505b92915050565b6106f882826107ef565b610703576000610705565b815b600160a060020a03928316600081815260208181526040808320600160e060020a031996909616808452958252808320805473ffffffffffffffffffffffffffffffffffffffff19169590971694909417909555908152600284528181209281529190925220805460ff19166001179055565b600080600160a060020a038416156107905783610792565b335b905061079d8361092a565b156107c357826107ad82826108aa565b6107b85760006107ba565b815b925050506106e8565b600160a060020a0390811660009081526020818152604080832086845290915290205416905092915050565b6000808061081d857f01ffc9a70000000000000000000000000000000000000000000000000000000061094c565b909250905081158061082d575080155b1561083d576000925050506106e8565b61084f85600160e060020a031961094c565b909250905081158061086057508015155b15610870576000925050506106e8565b61087a858561094c565b909250905060018214801561088f5750806001145b1561089f576001925050506106e8565b506000949350505050565b600160a060020a0382166000908152600260209081526040808320600160e060020a03198516845290915281205460ff1615156108f2576108eb83836107ef565b90506106e8565b50600160a060020a03808316600081815260208181526040808320600160e060020a0319871684529091529020549091161492915050565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161590565b6040517f01ffc9a7000000000000000000000000000000000000000000000000000000008082526004820183905260009182919060208160248189617530fa90519096909550935050505056fea165627a7a72305820377f4a2d4301ede9949f163f319021a6e9c687c292a5e2b2c4734c126b524e6c00291ba01820182018201820182018201820182018201820182018201820182018201820a01820182018201820182018201820182018201820182018201820182018201820"  # noqa: E501
)
ERC_1820_DEPLOYMENT_ADDRESS = "0xa990077c3205cbDf861e17Fa532eeB069cE9fF96"


@pytest.fixture
def config_change_heads_up_blocks() -> int:
    return 30


@pytest.fixture
def owner(accounts: Sequence[Account]) -> Account:
    return accounts[0]


@pytest.fixture
def non_owner(accounts: Sequence[Account], owner: Account) -> Account:
    non_owner = accounts[1]
    assert non_owner != owner
    return non_owner


@pytest.fixture
def keypers(accounts: Sequence[Account]) -> List[Account]:
    # as opposed to existing accounts, added ones have private keys known private keys which we
    # need to sign things
    return [accounts.add() for _ in range(3)]  # type: ignore


@pytest.fixture
def keyper_private_keys(keypers: Sequence[Account]) -> List[bytes]:
    return [decode_hex(keyper.private_key) for keyper in keypers]


@pytest.fixture
def appeal_blocks() -> int:
    return 10


@pytest.fixture
def config_contract(
    ConfigContract: ContractContainer, owner: Account, config_change_heads_up_blocks: int
) -> Any:
    config_contract = owner.deploy(ConfigContract, config_change_heads_up_blocks)
    return config_contract


@pytest.fixture
def batcher_contract(
    BatcherContract: ContractContainer,
    config_contract: Any,
    fee_bank_contract: Any,
    owner: Account,
) -> Any:
    config_contract = owner.deploy(BatcherContract, config_contract, fee_bank_contract)
    return config_contract


@pytest.fixture
def executor_contract(
    ExecutorContract: ContractContainer,
    config_contract: Any,
    mock_batcher_contract: Any,
    deposit_contract: Any,
    owner: Account,
) -> Any:
    executor_contract = owner.deploy(
        ExecutorContract, config_contract, mock_batcher_contract, deposit_contract
    )
    return executor_contract


@pytest.fixture
def keyper_slasher(
    KeyperSlasher: ContractContainer,
    config_contract: Any,
    executor_contract: Any,
    deposit_contract: Any,
    owner: Account,
    appeal_blocks: int,
) -> Any:
    keyper_slasher = owner.deploy(
        KeyperSlasher, appeal_blocks, config_contract, executor_contract, deposit_contract
    )
    return keyper_slasher


@pytest.fixture
def erc_1820_registry(owner: Account, web3: Web3) -> None:
    owner.transfer(ERC_1820_DEPLOYMENT_ADDRESS, "0.08 ether")
    web3.eth.send_raw_transaction(ERC_1820_DEPLOYMENT_TX)


@pytest.fixture
def deposit_token_contract(
    TestDepositTokenContract: ContractContainer,
    owner: Account,
    erc_1820_registry: None,
) -> Any:
    return owner.deploy(TestDepositTokenContract)


@pytest.fixture
def deposit_contract(
    DepositContract: ContractContainer,
    deposit_token_contract: Any,
    owner: Account,
) -> Any:
    deposit_contract = owner.deploy(DepositContract, deposit_token_contract)
    return deposit_contract


@pytest.fixture
def mock_target_contract(MockTargetContract: ContractContainer, owner: Account) -> Any:
    mock_target_contract = owner.deploy(MockTargetContract)
    return mock_target_contract


@pytest.fixture
def mock_batcher_contract(MockBatcherContract: ContractContainer, owner: Account) -> Any:
    mock_batcher_contract = owner.deploy(MockBatcherContract)
    return mock_batcher_contract


@pytest.fixture
def mock_target_function_selector(MockTargetContract: ContractContainer) -> bytes:
    function_name = "call"
    for selector, name in MockTargetContract.selectors.items():
        if name == function_name:
            return bytes(decode_hex(selector))
    raise AssertionError


@pytest.fixture
def fee_bank_contract(FeeBankContract: ContractContainer, accounts: Sequence[Account]) -> Any:
    fee_bank_contract = accounts[0].deploy(FeeBankContract)
    return fee_bank_contract


@pytest.fixture
def key_broadcast_contract(
    KeyBroadcastContract: ContractContainer, config_contract: Any, accounts: Sequence[Account]
) -> Any:
    key_broadcast_contract = accounts[0].deploy(KeyBroadcastContract, config_contract.address)
    return key_broadcast_contract


@pytest.fixture
def target_proxy_contract(
    TargetProxyContract: ContractContainer,
    owner: Account,
) -> Any:
    # make owner executor so that we can call it directly to simplify tests
    target_proxy_contract = owner.deploy(TargetProxyContract, owner)
    return target_proxy_contract


@pytest.fixture
def test_proxy_receiver(
    TestProxyReceiver: ContractContainer,
    accounts: Sequence[Account],
) -> Any:
    test_proxy_receiver = accounts[0].deploy(TestProxyReceiver)
    return test_proxy_receiver


@pytest.fixture(autouse=True)
def isolation(fn_isolation: Any) -> None:
    pass
