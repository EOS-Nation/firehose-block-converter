package main

//import (
//	pbcodec "github.com/dfuse-io/dfuse-eosio/pb/dfuse/eosio/codec/v1"
//	"github.com/streamingfast/playground-firehose-eosio-go/pbantelope"
//)
//
//func convertToAntelopeBlock(block *pbcodec.Block) *pbantelope.Block {
//
//	return &pbantelope.Block{
//		Id:                                      block.Id,
//		Number:                                  block.Number,
//		Version:                                 3,
//		Header:                                  block.Header,
//		ProducerSignature:                       block.ProducerSignature,
//		BlockExtensions:                         block.BlockExtensions,
//		DposProposedIrreversibleBlocknum:        block.DposProposedIrreversibleBlocknum,
//		DposIrreversibleBlocknum:                block.DposIrreversibleBlocknum,
//		BlockrootMerkle:                         block.BlockrootMerkle,
//		ProducerToLastProduced:                  block.ProducerToLastProduced,
//		ProducerToLastImpliedIrb:                block.ProducerToLastImpliedIrb,
//		ConfirmCount:                            block.ConfirmCount,
//		PendingSchedule:                         block.PendingSchedule,
//		ActivatedProtocolFeatures:               block.ActivatedProtocolFeatures,
//		RlimitOps:                               block.RlimitOps,
//		UnfilteredTransactions:                  block.UnfilteredTransactions,
//		FilteredTransactions:                    nil,
//		UnfilteredTransactionCount:              0,
//		FilteredTransactionCount:                0,
//		UnfilteredImplicitTransactionOps:        nil,
//		FilteredImplicitTransactionOps:          nil,
//		UnfilteredTransactionTraces:             nil,
//		FilteredTransactionTraces:               nil,
//		UnfilteredTransactionTraceCount:         0,
//		FilteredTransactionTraceCount:           0,
//		UnfilteredExecutedInputActionCount:      0,
//		FilteredExecutedInputActionCount:        0,
//		UnfilteredExecutedTotalActionCount:      0,
//		FilteredExecutedTotalActionCount:        0,
//		BlockSigningKey:                         "",
//		ActiveScheduleV1:                        nil,
//		ValidBlockSigningAuthorityV2:            nil,
//		ActiveScheduleV2:                        nil,
//		FilteringApplied:                        false,
//		FilteringIncludeFilterExpr:              "",
//		FilteringExcludeFilterExpr:              "",
//		FilteringSystemActionsIncludeFilterExpr: "",
//	}
//
//}
//
//func convertBlockHeader(blockHeader *pbcodec.BlockHeader) *pbantelope.BlockHeader {
//	return &pbantelope.BlockHeader{
//		Timestamp:        blockHeader.Timestamp,
//		Producer:         blockHeader.Producer,
//		Confirmed:        blockHeader.Confirmed,
//		Previous:         blockHeader.Previous,
//		TransactionMroot: blockHeader.TransactionMroot,
//		ActionMroot:      blockHeader.ActionMroot,
//		ScheduleVersion:  blockHeader.ScheduleVersion,
//		HeaderExtensions: convertExtensionSlice(blockHeader.HeaderExtensions),
//		NewProducersV1: &pbantelope.ProducerSchedule{
//			Version:   blockHeader.NewProducersV1.Version,
//			Producers: blockHeader.NewProducersV1.Producers,
//		},
//	}
//}
//
//func convertExtensionSlice(extensions []*pbcodec.Extension) []*pbantelope.Extension {
//
//	res := make([]*pbantelope.Extension, len(extensions))
//
//	for i, e := range extensions {
//		res[i] = &pbantelope.Extension{
//			Type: e.Type,
//			Data: e.Data,
//		}
//	}
//
//	return res
//}
