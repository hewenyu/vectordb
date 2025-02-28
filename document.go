package vectordb

import (
	"time"
)

// DocumentManager 定义了文档管理器必须实现的接口
type DocumentManager interface {
	// StoreDocument 存储文档
	StoreDocument(doc *Document) (string, error)

	// UpdateDocument 更新文档
	UpdateDocument(documentID string, updates map[string]interface{}) error

	// DeleteDocument 删除文档
	DeleteDocument(documentID string) error

	// GetDocument 获取文档
	GetDocument(documentID string) (*Document, error)

	// ListDocuments 列出符合条件的文档
	ListDocuments(workspaceID string, filter DocumentFilter) ([]*Document, error)

	// ProcessDocument 处理文档，包括分块和向量化
	ProcessDocument(documentID string, options ProcessOptions) error

	// GetDocumentChunks 获取文档的分块
	GetDocumentChunks(documentID string) ([]*DocumentChunk, error)

	// GetChunk 获取特定分块
	GetChunk(chunkID string) (*DocumentChunk, error)

	// DocumentExists 检查文档是否存在
	DocumentExists(documentID string) (bool, error)

	// PinDocument 将文档固定到工作区
	PinDocument(documentID string, workspaceID string) error

	// UnpinDocument 从工作区取消固定文档
	UnpinDocument(documentID string, workspaceID string) error

	// GetPinnedDocuments 获取工作区中固定的文档
	GetPinnedDocuments(workspaceID string) ([]*Document, error)
}

// Document 表示存储的文档
type Document struct {
	ID          string                 // 文档唯一标识符
	Name        string                 // 文档名称
	Type        string                 // 文档类型
	ContentType string                 // 内容MIME类型
	Size        int64                  // 文档大小
	Path        string                 // 存储路径
	WorkspaceID string                 // 所属工作区ID
	UserID      string                 // 上传用户ID
	CreatedAt   time.Time              // 创建时间
	UpdatedAt   time.Time              // 更新时间
	Metadata    map[string]interface{} // 文档元数据
	IsProcessed bool                   // 是否已处理
	ChunkCount  int                    // 分块数量
}

// DocumentChunk 表示文档分块
type DocumentChunk struct {
	ID         string                 // 分块唯一标识符
	DocumentID string                 // 所属文档ID
	Content    string                 // 分块内容
	TokenCount int                    // Token数量
	Sequence   int                    // 顺序号
	Metadata   map[string]interface{} // 分块元数据
	Vector     []float64              // 向量表示
}

// DocumentFilter 表示文档过滤条件
type DocumentFilter struct {
	Types       []string  // 文档类型列表
	StartDate   time.Time // 开始日期
	EndDate     time.Time // 结束日期
	UserID      string    // 用户ID
	IsProcessed *bool     // 是否已处理
	NamePrefix  string    // 名称前缀
	Limit       int       // 结果数量限制
	Offset      int       // 结果偏移量
}

// ProcessOptions 表示文档处理选项
type ProcessOptions struct {
	ChunkSize               int     // 分块大小
	ChunkOverlap            int     // 分块重叠
	IncludeMetadata         bool    // 是否在分块中包含元数据
	ExtractImages           bool    // 是否提取图片
	PreserveFormatting      bool    // 是否保留格式
	SkipEmbedding           bool    // 是否跳过嵌入
	EmbeddingModel          string  // 嵌入模型
	MaxTokenCount           int     // 最大Token数量
	SimilarityThreshold     float64 // 相似度阈值
	RemoveDuplicates        bool    // 是否移除重复内容
	AllowFallbackProcessors bool    // 是否允许备用处理器
}
