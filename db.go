package vectordb

// VectorDatabase 定义了向量数据库必须实现的接口
type VectorDatabase interface {
	// Connect 连接到向量数据库
	Connect() (interface{}, error)

	// Heartbeat 检查数据库连接状态
	Heartbeat() (map[string]interface{}, error)

	// CreateNamespace 创建新的命名空间/集合
	CreateNamespace(namespace string) error

	// DeleteNamespace 删除指定的命名空间/集合
	DeleteNamespace(namespace string) error

	// HasNamespace 检查命名空间是否存在
	HasNamespace(namespace string) bool

	// NamespaceCount 返回命名空间中的向量数量
	NamespaceCount(namespace string) (int, error)

	// TotalVectors 返回所有命名空间中的向量总数
	TotalVectors() (int, error)

	// StoreVectors 在指定命名空间中存储向量
	StoreVectors(namespace string, documents []Document) error

	// DeleteVectorsByMetadata 根据元数据删除向量
	DeleteVectorsByMetadata(namespace string, filter map[string]interface{}) error

	// PerformSimilaritySearch 执行相似度搜索
	PerformSimilaritySearch(params SimilaritySearchParams) (SimilaritySearchResult, error)

	// ListNamespaces 列出所有命名空间
	ListNamespaces() ([]string, error)

	// DistanceToSimilarity 将距离值转换为相似度分数
	DistanceToSimilarity(distance float64) float64
}

// SimilaritySearchParams 定义了相似度搜索的参数
type SimilaritySearchParams struct {
	Namespace           string      // 要搜索的命名空间
	Input               string      // 搜索查询文本
	LLMConnector        interface{} // LLM连接器实例
	SimilarityThreshold float64     // 相似度阈值
	TopN                int         // 返回结果数量
	FilterIdentifiers   []string    // 要过滤的文档标识符
	Rerank              bool        // 是否重排序结果
}

// SimilaritySearchResult 定义了相似度搜索的结果
type SimilaritySearchResult struct {
	ContextTexts []string                 // 相关文本内容
	Sources      []map[string]interface{} // 源文档信息
	Message      string                   // 处理消息
	Scores       []float64                // 相似度分数
}
