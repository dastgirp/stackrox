// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	metrics "github.com/stackrox/rox/central/metrics"
	mappings "github.com/stackrox/rox/central/namespace/index/mappings"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	batcher "github.com/stackrox/rox/pkg/batcher"
	blevehelper "github.com/stackrox/rox/pkg/blevehelper"
	ops "github.com/stackrox/rox/pkg/metrics"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
	"time"
)

const batchSize = 5000

const resourceName = "NamespaceMetadata"

type indexerImpl struct {
	index *blevehelper.BleveWrapper
}

type namespaceMetadataWrapper struct {
	*storage.NamespaceMetadata `json:"namespace_metadata"`
	Type                       string `json:"type"`
}

func (b *indexerImpl) AddNamespaceMetadata(namespacemetadata *storage.NamespaceMetadata) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Add, "NamespaceMetadata")
	if err := b.index.Index.Index(namespacemetadata.GetId(), &namespaceMetadataWrapper{
		NamespaceMetadata: namespacemetadata,
		Type:              v1.SearchCategory_NAMESPACES.String(),
	}); err != nil {
		return err
	}
	return b.index.IncTxnCount()
}

func (b *indexerImpl) AddNamespaceMetadatas(namespacemetadatas []*storage.NamespaceMetadata) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.AddMany, "NamespaceMetadata")
	batchManager := batcher.New(len(namespacemetadatas), batchSize)
	for {
		start, end, ok := batchManager.Next()
		if !ok {
			break
		}
		if err := b.processBatch(namespacemetadatas[start:end]); err != nil {
			return err
		}
	}
	return b.index.IncTxnCount()
}

func (b *indexerImpl) processBatch(namespacemetadatas []*storage.NamespaceMetadata) error {
	batch := b.index.NewBatch()
	for _, namespacemetadata := range namespacemetadatas {
		if err := batch.Index(namespacemetadata.GetId(), &namespaceMetadataWrapper{
			NamespaceMetadata: namespacemetadata,
			Type:              v1.SearchCategory_NAMESPACES.String(),
		}); err != nil {
			return err
		}
	}
	return b.index.Batch(batch)
}

func (b *indexerImpl) DeleteNamespaceMetadata(id string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Remove, "NamespaceMetadata")
	if err := b.index.Delete(id); err != nil {
		return err
	}
	return b.index.IncTxnCount()
}

func (b *indexerImpl) DeleteNamespaceMetadatas(ids []string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.RemoveMany, "NamespaceMetadata")
	batch := b.index.NewBatch()
	for _, id := range ids {
		batch.Delete(id)
	}
	if err := b.index.Batch(batch); err != nil {
		return err
	}
	return b.index.IncTxnCount()
}

func (b *indexerImpl) GetTxnCount() uint64 {
	return b.index.GetTxnCount()
}

func (b *indexerImpl) ResetIndex() error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Reset, "NamespaceMetadata")
	return blevesearch.ResetIndex(v1.SearchCategory_NAMESPACES, b.index.Index)
}

func (b *indexerImpl) Search(q *v1.Query) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "NamespaceMetadata")
	return blevesearch.RunSearchRequest(v1.SearchCategory_NAMESPACES, q, b.index.Index, mappings.OptionsMap)
}

func (b *indexerImpl) SetTxnCount(seq uint64) error {
	return b.index.SetTxnCount(seq)
}
