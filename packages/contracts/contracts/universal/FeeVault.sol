// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { L2StandardBridge } from "../L2/L2StandardBridge.sol";

/**
 * @title FeeVault
 * @notice The FeeVault contract contains the basic logic for the various different vault contracts
 *         used to hold fee revenue generated by the L2 system.
 */
abstract contract FeeVault {
    /**
     * @notice Emits each time that a withdrawal occurs.
     *
     * @param value Amount that was withdrawn (in wei).
     * @param to    Address that the funds were sent to.
     * @param from  Address that triggered the withdrawal.
     */
    event Withdrawal(uint256 value, address to, address from);

    /**
     * @notice Minimum balance before a withdrawal can be triggered.
     */
    uint256 public immutable MIN_WITHDRAWAL_AMOUNT;

    /**
     * @notice Wallet that will receive the fees on L1.
     */
    address public immutable RECIPIENT;

    /**
     * @notice The minimum gas limit for the FeeVault withdrawal transaction.
     */
    uint32 internal constant WITHDRAWAL_MIN_GAS = 35_000;

    /**
     * @notice Total amount of wei processed by the contract.
     */
    uint256 public totalProcessed;

    /**
     * @param _recipient           Wallet that will receive the fees on L1.
     * @param _minWithdrawalAmount Minimum balance before a withdrawal can be triggered.
     */
    constructor(address _recipient, uint256 _minWithdrawalAmount) {
        MIN_WITHDRAWAL_AMOUNT = _minWithdrawalAmount;
        RECIPIENT = _recipient;
    }

    /**
     * @notice Allow the contract to receive ETH.
     */
    receive() external payable {}

    /**
     * @notice Triggers a withdrawal of funds to the L1 fee wallet.
     */
    function withdraw() external {
        require(
            address(this).balance >= MIN_WITHDRAWAL_AMOUNT,
            "FeeVault: withdrawal amount must be greater than minimum withdrawal amount"
        );

        uint256 value = address(this).balance;
        totalProcessed += value;

        emit Withdrawal(value, RECIPIENT, msg.sender);

        L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE)).bridgeETHTo{ value: value }(
            RECIPIENT,
            WITHDRAWAL_MIN_GAS,
            bytes("")
        );
    }
}
