package wallet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tickettypes "gitlab.33.cn/chain33/chain33/plugin/dapp/ticket/types"
	ticketwallet "gitlab.33.cn/chain33/chain33/plugin/dapp/ticket/wallet"
	"gitlab.33.cn/chain33/chain33/types"
	"gitlab.33.cn/chain33/chain33/util/testnode"

	_ "gitlab.33.cn/chain33/chain33/plugin"
	_ "gitlab.33.cn/chain33/chain33/system"
)

var mock33 = testnode.New("testdata/chain33.test.toml", nil, false)

func getTicketList(addr string, t *testing.T) *tickettypes.ReplyTicketList {
	msg, err := mock33.GetAPI().Query(types.TicketX, "TicketList", &tickettypes.TicketList{addr, 1})
	require.Equal(t, err, nil)
	return msg.(*tickettypes.ReplyTicketList)
}

func Test_WalletTicket(t *testing.T) {
	t.Log("Begin wallet ticket test")
	defer mock33.Stop()

	err := mock33.WaitHeight(0)
	assert.Nil(t, err)
	ticketList := getTicketList("12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv", t)
	if ticketList == nil {
		t.Error("getTicketList return nil")
		return
	}
	ticketwallet.FlushTicket(mock33.GetAPI())
	err = mock33.WaitHeight(2)
	assert.Nil(t, err)
	header, err := mock33.GetAPI().GetLastHeader()
	require.Equal(t, err, nil)
	require.Equal(t, header.Height >= 2, true)
	t.Log("End wallet ticket test")
}
