-- SPDX-License-Identifier: AGPL-3.0-only
CREATE TABLE IF NOT EXISTS sigs_for_address
(
    address     String,                -- base-58
    signature FixedString(88),
    slot        UInt64,
    memo        String,                -- optional memo bytes
    error       Nullable(String),      -- text if failed
    block_time  Int64,
    updated_on  DateTime DEFAULT now(),
    INDEX addr_idx address TYPE bloom_filter GRANULARITY 1
) ENGINE = MergeTree()
ORDER BY (address, slot DESC);
