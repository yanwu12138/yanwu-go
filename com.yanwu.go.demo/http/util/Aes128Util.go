package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 10:49
 *
 * Description: Aes128加解密工具
 **/

const SECRET = "yanwu0527@163com" // 使用AES加密和解密的默认密钥

// Encrypt 使用默认密钥对字符串进行加密，输出Base64转码后的字符串
func Encrypt(content string) string {
	return EncryptByKey(content, SECRET)
}

// EncryptByKey 使用自定义密钥对字符串进行加密，输出Base64转码后的字符串
func EncryptByKey(content, key string) string {
	// ===== 校验数据
	if len(content) == 0 || len(key) == 0 {
		return content
	}
	// ----- 处理参数
	keyBytes := []byte(key)
	contentBytes := []byte(content)
	// ----- 对字符串进行加密
	block, _ := aes.NewCipher(keyBytes)
	blockSize := block.BlockSize()                                   // 获取秘钥块的长度
	contentBytes = pkcs5Padding(contentBytes, blockSize)             // 补全码
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize]) // 加密模式
	encrypted := make([]byte, len(contentBytes))                     // 创建数组
	blockMode.CryptBlocks(encrypted, contentBytes)                   // 加密
	return base64.StdEncoding.EncodeToString(encrypted)
}

// Decrypt 使用默认密钥对Base64转码后的字符串进行解密
func Decrypt(content string) string {
	return DecryptByKey(content, SECRET)
}

// DecryptByKey 使用自定义密钥对Base64转码后的字符串进行解密
func DecryptByKey(content, key string) string {
	// ===== 校验数据
	if len(content) == 0 || len(key) == 0 {
		return content
	}
	// ----- 处理参数
	keyBytes := []byte(key)
	contentBytes, _ := base64.StdEncoding.DecodeString(content)
	// ----- 对字符串进行解密
	block, _ := aes.NewCipher(keyBytes)                              // 分组秘钥
	blockSize := block.BlockSize()                                   // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize]) // 加密模式
	decrypted := make([]byte, len(contentBytes))                     // 创建数组
	blockMode.CryptBlocks(decrypted, contentBytes)                   // 解密
	return string(pkcs5UnPadding(decrypted))
}

// 补全码
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	return append(ciphertext, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

// 去除补全码
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	return origData[:(length - int(origData[length-1]))]
}

// test 测试函数
func test() {
	source := "dfasdfasdtwertrty5677jmh.l.o;p['90-sdafsdegerty56tr	qwr42345e  fghf lio';['opl/,k.?;k;lkilohiup;uio"
	fmt.Println("length:", len(source), " source:", source)
	encrypt := Encrypt(source)
	fmt.Println("length:", len(encrypt), "encrypt:", encrypt)
	decrypt := Decrypt(encrypt)
	fmt.Println("length:", len(decrypt), "decrypt:", decrypt)
	fmt.Println("result:", source == decrypt)
}
