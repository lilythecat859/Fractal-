-- SPDX-License-Identifier: AGPL-3.0-only
CREATE TABLE IF NOT EXISTS blocks
(
    slot          UInt64,
    blockhash     FixedString(44),      -- base-58 length
    parent_slot   UInt64,
    block_time    Int64,
    tx_count      UInt32,
    updated_on    DateTime DEFAULT now()
) ENGINE = MergeTree()
ORDER BY slot;
