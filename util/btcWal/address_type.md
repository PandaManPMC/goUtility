

地址类型	| 前缀	| 编码格式	| 兼容性 |	手续费	| 支持 |
----------|-----------|------------|--------------|--------|----------|
| Legacy（P2PKH）   | `` / `D`  | BaseBase58Check        | 高  最早的地址格式，兼容性强，但交易体积较大，手续费较高。SH） | `3` / `A` / `9` | Base58Check| 中      `3` 或 `M`  | 广泛支持 |
| NaBase58Check2WPKH） | `ltc将 SegWit 嵌套在 P2SH 中，兼容旧钱包，交易手续费较低。 ([Send and receive BTC/LTC - difference between SegWit and ...](https://help.crypto.com/en/articles/4056348-send-and-receive-btc-ltc-difference-between-segwit-and-legacy-address?utm_source=chatgpt.com))   | 低     | 逐渐普及 |
| Taproot（P2TR）    | `ltc1`ltc1`ge1p` | Bech32mBech32         | 最低原生 SegWit 地址，交易效率高，手续费低，但部分旧钱包可能不支持。cite:23]{index=23}


| 地址类型           | 输入字节大小（大约） | 输出字节大小（大约） | 说明 |
|--------------------|-----------------------|------------------------|------|
| Legacy (P2PKH)     | ~148 bytes            | ~34 bytes              | 最早、体积最大 |
| SegWit nested (P2SH-P2WPKH) | ~91 bytes (压缩) | ~32 bytes              | SegWit 包裹在 P2SH 中，体积较小 |
| Native SegWit (Bech32) | ~67 bytes            | ~31 bytes              | 纯 SegWit，最节省空间 |
| Taproot (P2TR)     | ~58 bytes             | ~43 bytes              | 新格式，节省空间，支持更多功能（主要用于 BTC） |