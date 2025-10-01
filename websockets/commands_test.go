package websockets

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/anchorageoss/ripple-client/data"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MessagesSuite struct{}

var _ = Suite(&MessagesSuite{})

func readResponseFile(c *C, msg interface{}, path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		c.Error(err)
	}

	if err = json.Unmarshal(b, msg); err != nil {
		c.Error(err)
	}
}

func (s *MessagesSuite) TestLedgerResponse(c *C) {
	msg := &LedgerCommand{}
	readResponseFile(c, msg, "testdata/ledger.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields
	c.Assert(msg.Result.Ledger.LedgerSequence, Equals, uint32(6917762))
	c.Assert(msg.Result.Ledger.CloseTime.String(), Equals, "2014-May-30 13:11:50 UTC")
	c.Assert(msg.Result.Ledger.Closed, Equals, true)
	c.Assert(msg.Result.Ledger.LedgerHash.String(), Equals, "0C5C5B39EA40D40ACA6EB47E50B2B85FD516D1A2BA67BA3E050349D3EF3632A4")
	c.Assert(msg.Result.Ledger.PreviousLedger.String(), Equals, "F8F0363803C30E659AA24D6A62A6512BA24BEA5AC52A29731ABA1E2D80796E8B")
	c.Assert(msg.Result.Ledger.TotalXRP, Equals, uint64(99999990098968782))
	c.Assert(msg.Result.Ledger.StateHash.String(), Equals, "46D3E36FE845B9A18293F4C0F134D7DAFB06D4D9A1C7E4CB03F8B293CCA45FA0")
	c.Assert(msg.Result.Ledger.TransactionHash.String(), Equals, "757CCB586D44F3C58E366EC7618988C0596277D3D5D0B412E49563B5EEDF04FF")

	c.Assert(msg.Result.Ledger.Transactions, HasLen, 7)
	tx0 := msg.Result.Ledger.Transactions[0]
	c.Assert(tx0.GetHash().String(), Equals, "2D0CE11154B655A2BFE7F3F857AAC344622EC7DAB11B1EBD920DCDB00E8646FF")
	c.Assert(tx0.MetaData.AffectedNodes, HasLen, 4)
}

func (s *MessagesSuite) TestLedgerResponseWithDIDSetTransaction(c *C) {
	msg := &LedgerCommand{}
	readResponseFile(c, msg, "testdata/ledger_DIDSet_tx.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields
	c.Assert(msg.Result.Ledger.LedgerSequence, Equals, uint32(91781709))
	c.Assert(msg.Result.Ledger.CloseTime.String(), Equals, "2024-Oct-30 23:22:21 UTC")
	c.Assert(msg.Result.Ledger.Closed, Equals, true)
	c.Assert(msg.Result.Ledger.LedgerHash.String(), Equals, "D40AA9E74A6345D737841FF9CB013DCBBB9824D70A8F7F0A1B182B4435723D08")
	c.Assert(msg.Result.Ledger.PreviousLedger.String(), Equals, "66C936B5C953E324C0118733A625EBB187D1E242BC9A3FB169CDB370959BC0AD")
	c.Assert(msg.Result.Ledger.TotalXRP, Equals, uint64(99987028538007593))
	c.Assert(msg.Result.Ledger.StateHash.String(), Equals, "3A8341781AAAF05C84E90BAB4BF0C09899790B831A00A6E236A1627314FE7578")
	c.Assert(msg.Result.Ledger.TransactionHash.String(), Equals, "ED48F333B8E4233A55975C62C1C88DF88957F7BAE61B534CF5CAE2B21D6AE310")

	c.Assert(msg.Result.Ledger.Transactions, HasLen, 36)
	tx0 := msg.Result.Ledger.Transactions[0]
	c.Assert(tx0.GetHash().String(), Equals, "0AF92DA7B9768692416FD447A7569AAE263D729E5418FFB9268E6A43B5480E59")
	c.Assert(tx0.MetaData.AffectedNodes, HasLen, 4)
}

func (s *MessagesSuite) TestLedgerResponseWithOracleSetTransaction(c *C) {
	msg := &LedgerCommand{}
	readResponseFile(c, msg, "testdata/ledger_OracleSet_tx.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields
	c.Assert(msg.Result.Ledger.LedgerSequence, Equals, uint32(91833494))
	c.Assert(msg.Result.Ledger.CloseTime.String(), Equals, "2024-Nov-02 07:20:00 UTC")
	c.Assert(msg.Result.Ledger.Closed, Equals, true)
	c.Assert(msg.Result.Ledger.LedgerHash.String(), Equals, "24977C8ACA83D1FB667DB39BDB97C2616B1BC711C980DDAEB708FE9467032C05")
	c.Assert(msg.Result.Ledger.PreviousLedger.String(), Equals, "F827D07203DB30BF8E852F55F3D94A24408CBDFD64CCA0EBB0B078F686182BBD")
	c.Assert(msg.Result.Ledger.TotalXRP, Equals, uint64(99987019809399404))
	c.Assert(msg.Result.Ledger.StateHash.String(), Equals, "DCC7EB8349810C1A0876D1B3D36CBB7F67276E900C88E105E1713633B997672D")
	c.Assert(msg.Result.Ledger.TransactionHash.String(), Equals, "EA18D87E62D1D82FCA75875D340D7792B9E3073133876D4B9371102BCB81D2F4")

	c.Assert(msg.Result.Ledger.Transactions, HasLen, 40)
	tx0 := msg.Result.Ledger.Transactions[0]
	c.Assert(tx0.GetHash().String(), Equals, "014B5D65F810BCCD6FB53177AF4035FBE5CFA200F8AEFCE886BE09CE7813C9C6")
	c.Assert(tx0.MetaData.AffectedNodes, HasLen, 4)
}

func (s *MessagesSuite) TestLedgerResponseWithNFTokenModifyTransaction(c *C) {
	msg := &LedgerCommand{}
	readResponseFile(c, msg, "testdata/ledger_NFTokenModify_tx.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields
	c.Assert(msg.Result.Ledger.LedgerSequence, Equals, uint32(96729412))
	c.Assert(msg.Result.Ledger.CloseTime.String(), Equals, "2025-Jun-11 05:16:10 UTC")
	c.Assert(msg.Result.Ledger.Closed, Equals, true)
	c.Assert(msg.Result.Ledger.LedgerHash.String(), Equals, "887DA6F5E0A59ABFBD0130A781A7665318D29D4BF4EEEC703A575E986D57078F")

	c.Assert(msg.Result.Ledger.Transactions, HasLen, 1)
	tx0 := msg.Result.Ledger.Transactions[0]
	c.Assert(tx0.GetHash().String(), Equals, "09D2CFA2D2D239408204E359A85A19324B4B4C915FE9ABE0244B583E98BBE281")
	c.Assert(tx0.MetaData.AffectedNodes, HasLen, 2)

	// Check that this is specifically an NFTokenModify transaction
	nftModify := tx0.Transaction.(*data.NFTokenModify)
	c.Assert(nftModify.GetTransactionType(), Equals, data.NFTOKEN_MODIFY)
	c.Assert(nftModify.NFTokenID.String(), Equals, "0018000025B89F24B381CABA5921FF0B634DE9111D915B2A4E84C37C000008E1")
	c.Assert(nftModify.URI.String(), Equals, "7B226E616D65223A225365636F6E6420506C616365204D6564616C222C226465736372697074696F6E223A22552B3146393438222C22696D616765223A22646174613A696D6167652F7376672B786D6C2C25336373766720786D6C6E733D27687474703A2F2F7777772E77332E6F72672F323030302F737667272076696577426F783D272D3430202D3430203830203830272533652533637465787420646F6D696E616E742D626173656C696E653D2763656E7472616C2720746578742D616E63686F723D276D6964646C652720666F6E742D73697A653D27363527253365F09FA5882533632F746578742533652533632F737667253365227D")
	c.Assert(nftModify.Account.String(), Equals, "rhSTwqSK13zdRmzHMZZP8i7DnuG27pwX76")
	c.Assert(nftModify.Sequence, Equals, uint32(2781))
	c.Assert(nftModify.Fee.String(), Equals, "0.000012")
}

func (s *MessagesSuite) TestLedgerResponseWithCredentialCreateTransaction(c *C) {
	msg := &LedgerCommand{}
	readResponseFile(c, msg, "testdata/ledger_credentialCreate_tx.json")

	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields
	c.Assert(msg.Result.Ledger.LedgerSequence, Equals, uint32(98615996))
	c.Assert(msg.Result.Ledger.CloseTime.String(), Equals, "2025-Sep-04 04:40:40 UTC")
	c.Assert(msg.Result.Ledger.Closed, Equals, true)
	c.Assert(msg.Result.Ledger.LedgerHash.String(), Equals, "311457C92C09A6AFDB2BEDA4967E2BCFE2D8B9E5F472A148B8F91DFA38FF7BDC")

	c.Assert(msg.Result.Ledger.Transactions, HasLen, 88)
	var tx *data.TransactionWithMetaData
	for _, txn := range msg.Result.Ledger.Transactions {
		if txn.MetaData.TransactionIndex == 49 { // CreateCredential has the index 49 but they are not ordered
			tx = txn
		}
	}

	c.Assert(tx.GetHash().String(), Equals, "F919D0BEA7328A55FE9B218CCB7F9AB9A2F48C789E0BDBE9AFA542E2EA4D1E2F")
	c.Assert(tx.MetaData.AffectedNodes, HasLen, 4)

	// Check that this is specifically a CreateCredential transaction
	CredentialCreate := tx.Transaction.(*data.CredentialCreate)
	c.Assert(CredentialCreate.GetTransactionType(), Equals, data.CREDENTIAL_CREATE)
}

func (s *MessagesSuite) TestLedgerResponseWithMPTokenLedgerEntryTypeTransaction(c *C) {
	msg := &LedgerCommand{}
	readResponseFile(c, msg, "testdata/ledger_MPTokenLedgerEntryType_tx.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields - basic checks first
	c.Assert(msg.Result.Ledger.LedgerSequence, Equals, uint32(99228110))
	c.Assert(msg.Result.Ledger.Closed, Equals, true)
	c.Assert(msg.Result.Ledger.LedgerHash.String(), Equals, "6F0E280A73195BF0EE3E4AB9063D783D79633E2C1F3E04EB6F9650D3D0B0C6C4")

	c.Assert(msg.Result.Ledger.Transactions, HasLen, 38)

	// Find the transaction that affects MPToken ledger entries
	var tx *data.TransactionWithMetaData
	for _, txn := range msg.Result.Ledger.Transactions {
		if txn.MetaData.TransactionIndex == 33 { // MPToken transaction has index 33
			tx = txn
		}
	}

	c.Assert(tx, NotNil)
	c.Assert(tx.GetHash().String(), Equals, "F5BB2974F02BD6670731F8CD53295795208B20FCB909E062E227DDD3124AB76F")
	c.Assert(tx.MetaData.AffectedNodes, HasLen, 3)

	// Verify that we can parse MPToken and MPTokenIssuance ledger entries
	mpTokenFound := false
	mpTokenIssuanceFound := false
	for _, node := range tx.MetaData.AffectedNodes {
		if node.ModifiedNode != nil {
			if node.ModifiedNode.LedgerEntryType.String() == "MPToken" {
				mpTokenFound = true
				// Verify it's actually an MPToken entry
				mpToken, ok := node.ModifiedNode.FinalFields.(*data.MPToken)
				c.Assert(ok, Equals, true)
				c.Assert(mpToken != nil, Equals, true)
				// Check some fields exist
				c.Assert(mpToken.MPTAmount, NotNil)
				c.Assert(mpToken.MPTokenIssuanceID, NotNil)
			}
			if node.ModifiedNode.LedgerEntryType.String() == "MPTokenIssuance" {
				mpTokenIssuanceFound = true
				// Verify it's actually an MPTokenIssuance entry
				mpTokenIssuance, ok := node.ModifiedNode.FinalFields.(*data.MPTokenIssuance)
				c.Assert(ok, Equals, true)
				c.Assert(mpTokenIssuance != nil, Equals, true)
				// Check some fields exist
				c.Assert(mpTokenIssuance.OutstandingAmount, NotNil)
				c.Assert(mpTokenIssuance.Issuer, NotNil)
			}
		}
	}

	c.Assert(mpTokenFound, Equals, true)
	c.Assert(mpTokenIssuanceFound, Equals, true)
}

func (s *MessagesSuite) TestLedgerHeaderResponse(c *C) {
	msg := &LedgerHeaderCommand{}
	readResponseFile(c, msg, "testdata/ledger_header.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields
	c.Assert(len(msg.Result.LedgerData), Equals, 118)
	c.Assert(msg.Result.LedgerSequence, Equals, uint32(32570))
	c.Assert(msg.Result.Ledger.LedgerSequence, Equals, uint32(32570))
	c.Assert(msg.Result.Ledger.Accepted, Equals, true)
	c.Assert(msg.Result.Ledger.CloseTime.String(), Equals, "2013-Jan-01 03:21:10 UTC")
	c.Assert(msg.Result.Ledger.Closed, Equals, true)
	c.Assert(msg.Result.Ledger.LedgerHash.String(), Equals, "4109C6F2045FC7EFF4CDE8F9905D19C28820D86304080FF886B299F0206E42B5")
	c.Assert(msg.Result.Ledger.PreviousLedger.String(), Equals, "60A01EBF11537D8394EA1235253293508BDA7131D5F8710EFE9413AA129653A2")
	c.Assert(msg.Result.Ledger.TotalXRP, Equals, uint64(99999999999996320))
	c.Assert(msg.Result.Ledger.StateHash.String(), Equals, "3806AF8F22037DE598D30D38C8861FADF391171D26F7DE34ACFA038996EA6BEB")
	c.Assert(msg.Result.Ledger.TransactionHash.String(), Equals, "0000000000000000000000000000000000000000000000000000000000000000")
	c.Assert(msg.Result.Ledger.Transactions, HasLen, 0)
}

func (s *MessagesSuite) TestTxResponse(c *C) {
	msg := &TxCommand{}
	readResponseFile(c, msg, "testdata/tx.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	// Result fields
	c.Assert(msg.Result.Date.String(), Equals, "2014-May-30 13:11:50 UTC")
	c.Assert(msg.Result.Validated, Equals, true)
	c.Assert(msg.Result.MetaData.AffectedNodes, HasLen, 4)
	c.Assert(msg.Result.MetaData.TransactionResult.String(), Equals, "tesSUCCESS")

	offer := msg.Result.Transaction.(*data.OfferCreate)
	c.Assert(msg.Result.GetHash().String(), Equals, "2D0CE11154B655A2BFE7F3F857AAC344622EC7DAB11B1EBD920DCDB00E8646FF")
	c.Assert(offer.GetType(), Equals, "OfferCreate")
	c.Assert(offer.Account.String(), Equals, "rwpxNWdpKu2QVgrh5LQXEygYLshhgnRL1Y")
	c.Assert(offer.Fee.String(), Equals, "0.00001")
	c.Assert(offer.SigningPubKey.String(), Equals, "02BD6F0CFD0182F2F408512286A0D935C58FF41169DAC7E721D159D711695DFF85")
	c.Assert(offer.TxnSignature.String(), Equals, "30440220216D42DF672C1CC7EF0CA9C7840838A2AF5FEDD4DEFCBA770C763D7509703C8702203C8D831BFF8A8BC2CC993BECB4E6C7BE1EA9D394AB7CE7C6F7542B6CDA781467")
	c.Assert(offer.Sequence, Equals, uint32(1681497))
}

func (s *MessagesSuite) TestAccountTxResponse(c *C) {
	msg := &AccountTxCommand{}
	readResponseFile(c, msg, "testdata/account_tx.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	c.Assert(len(msg.Result.Transactions), Equals, 2)
	c.Assert(msg.Result.Transactions[1].Date.String(), Equals, "2014-Jun-19 14:14:40 UTC")
	offer := msg.Result.Transactions[1].Transaction.(*data.OfferCreate)
	c.Assert(offer.TakerPays.String(), Equals, "0.034800328/BTC/rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B")
}

func (s *MessagesSuite) TestLedgerDataResponse(c *C) {
	msg := &LedgerDataCommand{}
	readResponseFile(c, msg, "testdata/ledger_data.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	c.Assert(msg.Result.LedgerSequence, Equals, uint32(6281820))
	c.Assert(msg.Result.Hash.String(), Equals, "83CC350B1CDD9792D47F60D3DBB7673518FD6E71821070673E6EAE65DE69086B")
	c.Assert(msg.Result.Marker.String(), Equals, "02DE1A2AD4332A1AF01C59F16E45218FA70E5792BD963B6D7ACF188D6D150607")
	c.Assert(len(msg.Result.State), Equals, 2048)
	c.Assert(msg.Result.State[0].GetType(), Equals, "AccountRoot")
	c.Assert(msg.Result.State[0].GetLedgerIndex().String(), Equals, "00001A2969BE1FC85F1D7A55282FA2E6D95C71D2E4B9C0FDD3D9994F3C00FF8F")
}

func (s *MessagesSuite) TestRipplePathFindResponse(c *C) {
	msg := &RipplePathFindCommand{}
	readResponseFile(c, msg, "testdata/ripple_path_find.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	c.Assert(msg.Result.DestAccount.String(), Equals, "r9Dr5xwkeLegBeXq6ujinjSBLQzQ1zQGjH")
	c.Assert(msg.Result.DestCurrencies, HasLen, 6)
	c.Assert(msg.Result.Alternatives, HasLen, 1)
	c.Assert(msg.Result.Alternatives[0].SrcAmount.String(), Equals, "0.9940475268/USD/rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B")
	c.Assert(msg.Result.Alternatives[0].PathsComputed, HasLen, 4)
	c.Assert(msg.Result.Alternatives[0].PathsCanonical, HasLen, 0)
	c.Assert(msg.Result.Alternatives[0].PathsComputed[0], HasLen, 2)
	c.Assert(msg.Result.Alternatives[0].PathsComputed[0].String(), Equals, "XRP => SGD/r9Dr5xwkeLegBeXq6ujinjSBLQzQ1zQGjH")
}

func (s *MessagesSuite) TestAccountInfoResponse(c *C) {
	msg := &AccountInfoCommand{}
	readResponseFile(c, msg, "testdata/account_info.json")

	// Response fields
	c.Assert(msg.Status, Equals, "success")
	c.Assert(msg.Type, Equals, "response")

	c.Assert(msg.Result.LedgerSequence, Equals, uint32(7636529))
	c.Assert(*msg.Result.AccountData.TransferRate, Equals, uint32(1002000000))
	c.Assert(msg.Result.AccountData.LedgerEntryType, Equals, data.ACCOUNT_ROOT)
	c.Assert(*msg.Result.AccountData.Sequence, Equals, uint32(546))
	c.Assert(msg.Result.AccountData.Balance.String(), Equals, "10321199.422233")
}
