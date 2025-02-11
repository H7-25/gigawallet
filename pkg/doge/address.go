package doge

func Hash160(bytes []byte) []byte {
	return RIPEMD160(Sha256(bytes))
}

func PubKeyToP2PKH(key ECPubKeyCompressed, chain *ChainParams) string {
	if len(key) != ECPubKeyCompressedLen {
		panic("PubKeyToP2PKH: wrong pubkey length")
	}
	payload := Hash160(key[:])
	ver_hash := [1 + 20 + 4]byte{}
	ver_hash[0] = chain.p2pkh_address_prefix
	if copy(ver_hash[1:], payload) != 20 {
		panic("PubKeyToP2PKH: wrong RIPEMD-160 length")
	}
	return Base58EncodeCheck(ver_hash[0:21])
}

func ScriptToP2SH(redeemScript []byte, chain *ChainParams) string {
	if len(redeemScript) < 1 {
		panic("ScriptToP2SH: bad script length")
	}
	payload := Hash160(redeemScript)
	ver_hash := [1 + 20 + 4]byte{}
	ver_hash[0] = chain.p2sh_address_prefix
	if copy(ver_hash[1:], payload) != 20 {
		panic("ScriptToP2SH: wrong RIPEMD-160 length")
	}
	return Base58EncodeCheck(ver_hash[0:21])
}

func ValidateP2PKH(address string, chain *ChainParams) bool {
	key, err := Base58DecodeCheck(address)
	if err != nil {
		return false
	}
	return key[0] == chain.p2pkh_address_prefix
}

func ValidateP2SH(address string, chain *ChainParams) bool {
	key, err := Base58DecodeCheck(address)
	if err != nil {
		return false
	}
	return key[0] == chain.p2sh_address_prefix
}
