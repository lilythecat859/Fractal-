-- SPDX-License-Identifier: AGPL-3.0-only
CREATE TABLE IF NOT EXISTS transactions
(
    slot        UInt64,
    tx_hash     FixedString(88),        -- base-58 sig length
    idx         UInt32,                -- index inside block
    meta        String,                -- bincode blob
    message     String,                -- bincode blob
    block_time  Int64,
    updated_on  DateTime DEFAULT now()
) ENGINE = MergeTree()
ORDER BY (slot, idx);
