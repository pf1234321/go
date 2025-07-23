package main

import (
	"context"      // 用于设置超时、取消等上下文管理
	"crypto/ecdsa" // 用于处理以太坊账户的私钥
	"fmt"          // 用于格式化输出
	"log"          // 用于日志输出
	"math/big"     // 用于大整数运算（以太坊金额等）
	"time"

	count "Geth/count" // 导入本地 count 合约 Go 绑定包

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // go-ethereum 的合约绑定工具
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"    // go-ethereum 的加密工具
	"github.com/ethereum/go-ethereum/ethclient" // go-ethereum 的以太坊客户端
)

func main() {
	// 连接到以太坊节点（这里用的是 Sepolia 测试网的 RPC 地址）
	client, err := ethclient.Dial("https://still-blissful-diamond.ethereum-sepolia.quiknode.pro/ad6d4f1440eba30e70a48500a0729e86b6f9620d/")
	if err != nil {
		log.Fatal(err) // 连接失败则退出
	}

	// 通过私钥字符串创建私钥对象（HexToECDSA 需要 64 位十六进制字符串）
	privateKey, err := crypto.HexToECDSA("5e367b13123a61b3a995f5d93eb1cf1fb6144bc3e5e4f1e4299f6637f92b2cf4")
	if err != nil {
		log.Fatal(err) // 私钥格式错误则退出
	}

	// 获取公钥对象
	publicKey := privateKey.Public()
	// 类型断言，确保公钥是 *ecdsa.PublicKey 类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 由公钥推导出以太坊地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress==", fromAddress)
	// 查询当前账户的 nonce（即交易计数器，防止重放攻击）
	// nonce 就是“账户已发出的交易数”，新交易必须用最新的 nonce，确保唯一和顺序。
	// 	查询 fromAddress 账户的下一个可用 nonce（即当前已发出交易数），用于新交易的唯一标识。
	// 如果你发送一笔新交易，必须用这个 nonce，否则交易会被拒绝或卡住。
	// 如果你连续发多笔交易，nonce 必须手动递增，否则后面的交易不会被网络接受。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("nonce==", nonce)
	if err != nil {
		log.Fatal(err)
	}

	// 查询当前网络建议的 gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 获取当前链的 chainId（防止重放攻击）  11155111
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainId===", chainId)

	// 创建交易签名器（带 chainId，EIP-155 防重放）
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce)) // 设置 nonce
	auth.Value = big.NewInt(0)            // 发送以太币数量（这里为 0，表示只部署合约不转账）
	auth.GasLimit = uint64(900000)        // 设置 gas 上限
	// 增加 30%
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(13))
	gasPrice = new(big.Int).Div(gasPrice, big.NewInt(10))
	auth.GasPrice = gasPrice // 设置 gas price

	// input := "1.0" // 部署合约时的构造函数参数（版本号）
	// 部署 count 合约，返回合约地址、交易对象、合约实例和错误
	address, tx, instance, err := count.DeployCount(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	waitForTxMined(client, tx.Hash()) // 等待部署合約交易被打包

	// 根據合約地址和以太坊節點連接，創建一個 Go 語言的合約操作對象，
	// 方便後續調用合約方法。如果失敗就直接報錯退出  变成中文简体
	countContract, err := count.NewCount(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("countContract====", countContract)

	tx, err = countContract.IncreaceCount(auth)
	if err != nil {
		log.Fatal(err)
	}
	waitForTxMined(client, tx.Hash()) // 等待 IncreaceCount 交易被打包

	fmt.Println("tx hash:", tx.Hash().Hex()) // 输出交易哈希
	fmt.Println(address.Hex())               // 输出新部署的合约地址
	fmt.Println(instance)                    // 输出部署合约的交易哈希
	_ = instance                             // 占位，防止 instance 未使用的编译警告
}

// 等待交易被打包
func waitForTxMined(client *ethclient.Client, txHash common.Hash) {
	fmt.Printf("等待交易 %s 被打包...\n", txHash.Hex())
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil && receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
			fmt.Println("交易已被打包！")
			break
		}
		time.Sleep(5 * time.Second) // 每5秒查一次
	}
}
