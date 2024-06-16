// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./IConnect3DATaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from connect3DA.
 * @author Layr Labs, Inc.
 */
contract Connect3DAServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    IConnect3DATaskManager
        public immutable connect3DATaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyConnect3DATaskManager() {
        require(
            msg.sender == address(connect3DATaskManager),
            "onlyConnect3DATaskManager: not from connect3 DA task manager"
        );
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        IConnect3DATaskManager _connect3DATaskManager
    )
        ServiceManagerBase(
            _avsDirectory,
            IPaymentCoordinator(address(0)), // c3-da doesn't need to deal with payments
            _registryCoordinator,
            _stakeRegistry
        )
    {
        connect3DATaskManager = _connect3DATaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlyConnect3DATaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
