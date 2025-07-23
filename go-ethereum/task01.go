package main

// package main 告訴 Go 編譯器這是一個可以直接執行的主程式（而不是一個函式庫）。

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接到以太坊Sepolia测试网的QuickNode节点，返回一个以太坊客户端对象
	client, err := ethclient.Dial("https://still-blissful-diamond.ethereum-sepolia.quiknode.pro/ad6d4f1440eba30e70a48500a0729e86b6f9620d/")
	if err != nil {
		log.Fatal(err) // 如果连接失败，输出错误并终止程序
	}
	// 环境搭建
	// 安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
	// 注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
	// 查询区块
	// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
	// 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
	// 输出查询结果到控制台。
	readBlockInfo(client)

	// 发送交易
	// 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
	// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
	// 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
	// 对交易进行签名，并将签名后的交易发送到网络。
	// 输出交易的哈希值。
	sendTranction(client)
	// 	发送方地址: 0x21721DaCd81DfcdD978ea8445e1C43dD3F5Ea2BB
	// 📍 接收方地址: 0xF13F669Ec7e2961B008BFBef02Fd046716791e79
	// 📝 交易哈希: 0xe8baface89572454542e9da68bc055daa2273fa162bacea7623b36a8108f203b

}

func readBlockInfo(client *ethclient.Client) {

	// 指定要查询的区块号，这里是8651576
	// 可以在 https://sepolia.etherscan.io/txs?block=8651576 查看该区块信息
	blockNumber := big.NewInt(8651576)

	// 获取指定区块号的区块头信息
	// 在 Go 语言中，context 是一个标准库（context 包）提供的类型，
	// 主要用于在跨 API、跨 goroutine 的操作中传递取消信号、超时控制、截止时间、请求范围内的元数据等。
	// context 让你可以优雅地控制和管理一组相关操作的生命周期，尤其适合网络请求、数据库操作等需要超时和取消的场景。
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err) // 如果获取区块头失败，输出错误并终止程序
	}
	// 打印区块头的区块号 8651576
	fmt.Println(header.Number.Uint64()) // 8651576
	// 打印区块头的时间戳（自1970年1月1日以来的秒数） 1751164092
	fmt.Println(header.Time)
	// 打印区块头的难度值  0
	fmt.Println(header.Difficulty.Uint64())
	// 打印区块头的哈希值 0x79da6b217a65d776d12d83ea5017941d97ad23f3f21270c080f5a804f4b12e5e
	fmt.Println(header.Hash().Hex())

	// 获取指定区块号的完整区块信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err) // 如果获取区块失败，输出错误并终止程序
	}
	// 打印区块的区块号 8651576
	fmt.Println(block.Number().Uint64())
	// 打印区块的时间戳 1751164092
	fmt.Println(block.Time())
	// 打印区块的难度值 0
	fmt.Println(block.Difficulty().Uint64())
	// 打印区块的哈希值 0x79da6b217a65d776d12d83ea5017941d97ad23f3f21270c080f5a804f4b12e5e
	fmt.Println(block.Hash().Hex())
	// 打印区块中交易的数量 258
	fmt.Println(len(block.Transactions()))
	// 获取该区块哈希下的交易数量（与上面打印的交易数量应一致）
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err) // 如果获取交易数量失败，输出错误并终止程序
	}
	// 打印该区块的交易数量 258
	fmt.Println(count)
}

func sendTranction(client *ethclient.Client) {
	defer client.Close()

	// 私钥（从程序生成）
	privateKeyHex := "5e367b13123a61b3a995f5d93eb1cf1fb6144bc3e5e4f1e4299f6637f92b2cf4"

	// 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("私钥解析失败:", err)
	}

	// 从私钥获取地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥类型断言失败")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("🔑 发送方地址:", fromAddress.Hex())

	// 接收方地址（示例地址）   account2
	toAddress := common.HexToAddress("0xF13F669Ec7e2961B008BFBef02Fd046716791e79")
	fmt.Println("📍 接收方地址:", toAddress.Hex())

	// 获取账户余额
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal("获取余额失败:", err)
	}
	fmt.Println("💰 账户余额:", balance.String(), "wei")

	// 获取当前gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取gas价格失败:", err)
	}
	fmt.Println("⛽ Gas价格:", gasPrice.String(), "wei")

	// 获取nonce（交易序号）
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取nonce失败:", err)
	}
	fmt.Println("🔢 Nonce:", nonce)

	// 转账金额（0.001 ETH）
	value := big.NewInt(1000000000000000) // 0.001 ETH in wei
	fmt.Println("💸 转账金额:", value.String(), "wei (0.001 ETH)")

	// 创建交易
	tx := types.NewTransaction(
		nonce,
		toAddress,
		value,
		21000, // gas limit for simple ETH transfer
		gasPrice,
		nil, // data
	)

	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("获取链ID失败:", err)
	}

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("签名交易失败:", err)
	}

	fmt.Println("✅ 交易已签名")
	fmt.Println("📝 交易哈希:", signedTx.Hash().Hex())

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("发送交易失败:", err)
	}
	fmt.Println("🚀 交易已发送！")

	newBalance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal("获取新余额失败:", err)
	}
	fmt.Println("💰 新账户余额:", newBalance.String(), "wei")

	fmt.Println("\n=== 重要提醒 ===")
	fmt.Println("❌ 此代码仅用于演示，实际使用时需要：")
	fmt.Println("   1. 配置正确的Infura项目ID")
	fmt.Println("   2. 确保账户有足够的ETH支付gas费")
	fmt.Println("   3. 在测试网络上进行测试")
	fmt.Println("   4. 妥善保管私钥，不要泄露")
}
